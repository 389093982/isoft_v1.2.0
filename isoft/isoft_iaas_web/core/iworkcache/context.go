package iworkcache

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"sync"
)

var cacheContextMap = make(map[int64]*CacheContext, 0)

var mutex sync.Mutex

type GetBlockStepExecuteOrder func(blockSteps []*block.BlockStep) []*block.BlockStep
type GetCacheParamInputSchemaFunc func(step *iwork.WorkStep) *iworkmodels.ParamInputSchema

func UpdateCacheContext(work_id int64, orderFunc GetBlockStepExecuteOrder, paramInputSchemaFunc GetCacheParamInputSchemaFunc) {
	mutex.Lock()
	defer mutex.Unlock()
	context := &CacheContext{WorkId: work_id}
	context.LoadCache(orderFunc, paramInputSchemaFunc)
	cacheContextMap[work_id] = context
}

var mutex2 sync.Mutex

func GetCacheContext(work_id int64, orderFunc GetBlockStepExecuteOrder, paramInputSchemaFunc GetCacheParamInputSchemaFunc) *CacheContext {
	mutex2.Lock()
	defer mutex2.Unlock()
	if _, ok := cacheContextMap[work_id]; !ok {
		UpdateCacheContext(work_id, orderFunc, paramInputSchemaFunc)
	}
	return cacheContextMap[work_id]
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type CacheContext struct {
	WorkId              int64
	Work                iwork.Work
	Steps               []iwork.WorkStep
	BlockStepOrdersMap  map[int64][]*block.BlockStep // key 为父节点 StepId
	ParamInputSchemaMap map[int64]*iworkmodels.ParamInputSchema
	err                 error
}

func (this *CacheContext) LoadCache(orderFunc GetBlockStepExecuteOrder, paramInputSchemaFunc GetCacheParamInputSchemaFunc) {
	o := orm.NewOrm()
	// 缓存 work
	this.Work, this.err = iwork.QueryWorkById(this.WorkId, o)
	checkError(this.err)
	// 缓存 workSteps
	this.Steps, this.err = iwork.QueryAllWorkStepInfo(this.WorkId, o)
	checkError(this.err)
	blockSteps := orderFunc(block.ParseToBlockStep(this.Steps))
	// 缓存 blockStepOrder
	this.BlockStepOrdersMap = map[int64][]*block.BlockStep{
		-1: blockSteps,
	}
	// 缓存子节点 blockStepOrder
	for _, blockStep := range blockSteps {
		this.loadChildrenBlockStepOrders(blockStep, orderFunc)
	}
	// 缓存 paramInputSchema
	this.ParamInputSchemaMap = make(map[int64]*iworkmodels.ParamInputSchema, 0)
	for _, workStep := range this.Steps {
		paramInputSchema := paramInputSchemaFunc(&workStep)
		this.ParamInputSchemaMap[workStep.WorkStepId] = paramInputSchema
	}
}

func (this *CacheContext) loadChildrenBlockStepOrders(blockStep *block.BlockStep, orderFunc GetBlockStepExecuteOrder) {
	if blockStep.ChildBlockSteps != nil && len(blockStep.ChildBlockSteps) > 0 {
		childrenBlockSteps := orderFunc(blockStep.ChildBlockSteps)
		this.BlockStepOrdersMap[blockStep.Step.WorkStepId] = childrenBlockSteps
		for _, childrenBlockStep := range childrenBlockSteps {
			this.loadChildrenBlockStepOrders(childrenBlockStep, orderFunc)
		}
	}
}
