package iworkcache

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/models"
	"strings"
	"sync"
)

type IParamSchemaCacheParser interface {
	GetCacheParamInputSchema(replaceStep ...*models.WorkStep) *iworkmodels.ParamInputSchema
}

func getBlockStepExecuteOrder(blockSteps []*block.BlockStep) []*block.BlockStep {
	order := make([]*block.BlockStep, 0)
	deferOrder := make([]*block.BlockStep, 0)
	var end *block.BlockStep
	for _, blockStep := range blockSteps {
		if blockStep.Step.IsDefer == "true" {
			deferOrder = append(deferOrder, blockStep)
		} else if blockStep.Step.WorkStepType == "work_end" {
			end = blockStep
		} else {
			order = append(order, blockStep)
		}
	}
	// is_defer 和 work_end 都是需要延迟执行
	order = append(order, datatypeutil.ReverseSlice(deferOrder).([]*block.BlockStep)...)
	if end != nil {
		order = append(order, end)
	}
	return order
}

var workCacheMap = new(sync.Map)

var mutex sync.Mutex

func UpdateWorkCache(work_id int64, paramSchemaCacheParser IParamSchemaCacheParser) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	defer func() {
		if err1 := recover(); err1 != nil {
			err = err1.(error)
		}
	}()
	cache := &WorkCache{WorkId: work_id}
	workCacheMap.Store(work_id, &cache)
	cache.FlushCache(paramSchemaCacheParser)
	return
}

var mutex2 sync.Mutex

func GetWorkCache(work_id int64, paramSchemaCacheParser IParamSchemaCacheParser) (*WorkCache, error) {
	if cache, ok := workCacheMap.Load(work_id); ok {
		return cache.(*WorkCache), nil
	}
	UpdateWorkCache(work_id, paramSchemaCacheParser)
	if cache, ok := workCacheMap.Load(work_id); ok {
		return cache.(*WorkCache), nil
	}
	return nil, errors.New("cache was not exist")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type WorkCache struct {
	WorkId              int64
	Work                models.Work
	Steps               []models.WorkStep
	BlockStepOrdersMap  map[int64][]*block.BlockStep            // key 为父节点 StepId
	ParamInputSchemaMap map[int64]*iworkmodels.ParamInputSchema // key 为 WorkStepId
	SubWorkNameMap      map[int64]string                        // key 为 WorkStepId
	err                 error
}

func (this *WorkCache) FlushCache(paramSchemaCacheParser IParamSchemaCacheParser) {

	o := orm.NewOrm()
	// 缓存 work
	this.Work, this.err = models.QueryWorkById(this.WorkId, o)
	checkError(this.err)
	// 缓存 workSteps
	this.Steps, this.err = models.QueryAllWorkStepInfo(this.WorkId, o)
	checkError(this.err)
	blockSteps := getBlockStepExecuteOrder(block.ParseToBlockStep(this.Steps))
	// 缓存 blockStepOrder
	this.BlockStepOrdersMap = map[int64][]*block.BlockStep{
		-1: blockSteps,
	}
	// 缓存子节点 blockStepOrder
	for _, blockStep := range blockSteps {
		this.cacheChildrenBlockStepOrders(blockStep)
	}
	// 缓存 paramInputSchema
	this.ParamInputSchemaMap = make(map[int64]*iworkmodels.ParamInputSchema, 0)
	for _, workStep := range this.Steps {
		paramInputSchema := paramSchemaCacheParser.GetCacheParamInputSchema(&workStep)
		this.ParamInputSchemaMap[workStep.WorkStepId] = paramInputSchema
	}
	// 缓存 subWorkName
	this.SubWorkNameMap = make(map[int64]string, 0)
	for _, workStep := range this.Steps {
		if workStep.WorkStepType == iworkconst.NODE_TYPE_WORK_SUB {
			this.SubWorkNameMap[workStep.WorkStepId], _ = this.getWorkSubName(&workStep)
		}
	}
}

func (this *WorkCache) cacheChildrenBlockStepOrders(blockStep *block.BlockStep) {
	if blockStep.ChildBlockSteps != nil && len(blockStep.ChildBlockSteps) > 0 {
		// 获取并记录 order
		childrenBlockSteps := getBlockStepExecuteOrder(blockStep.ChildBlockSteps)
		this.BlockStepOrdersMap[blockStep.Step.WorkStepId] = childrenBlockSteps
		// 循环递归
		for _, childrenBlockStep := range childrenBlockSteps {
			this.cacheChildrenBlockStepOrders(childrenBlockStep)
		}
	}
}

func (this *WorkCache) getWorkSubName(workStep *models.WorkStep) (string, error) {
	paramInputSchema, err := iworkmodels.ParseToParamInputSchema(workStep.WorkStepInput)
	if err == nil {
		workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
		if strings.TrimSpace(workSubName) == "" {
			return "", errors.New("empty workSubName was found!")
		}
		return workSubName, nil
	}
	return "", err
}
