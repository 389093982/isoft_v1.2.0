package iworkcache

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup/dipool/pool"
	"reflect"
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

var workCacheMap sync.Map

func DeleteWorkCache(work_id int64) {
	workCacheMap.Delete(work_id)
}

func UpdateWorkCache(work_id int64) (err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println(string(errorutil.PanicTrace(4)))
			err = errorutil.ToError(err1)
		}
	}()
	cache := &WorkCache{WorkId: work_id}
	cache.FlushCache()
	workCacheMap.Store(work_id, cache)
	workCacheMap.Store(cache.Work.WorkName, cache)
	return
}

func LoadWorkCache(work_id int64) (*WorkCache, error) {
	if cache, ok := workCacheMap.Load(work_id); ok {
		return cache.(*WorkCache), nil
	} else {
		return nil, errors.New(fmt.Sprintf(`%d is not exist!`, work_id))
	}
}

func GetWorkCacheWithName(work_name string) (*WorkCache, error) {
	if cache, ok := workCacheMap.Load(work_name); ok {
		return cache.(*WorkCache), nil
	}
	work, err := models.QueryWorkByName(work_name, orm.NewOrm())
	if err != nil {
		return nil, err
	}
	return GetWorkCache(work.Id)
}

func GetWorkCache(work_id int64) (*WorkCache, error) {
	if cache, ok := workCacheMap.Load(work_id); ok {
		return cache.(*WorkCache), nil
	}
	if err := UpdateWorkCache(work_id); err != nil {
		return nil, err
	} else {
		return GetWorkCache(work_id)
	}
}

func GetAllWorkCache() []*WorkCache {
	works, _ := models.QueryAllFilterAndWorks()
	var m sync.WaitGroup
	for _, work := range works {
		m.Add(1)
		go func(work models.Work) {
			if _, ok := workCacheMap.Load(work.Id); !ok {
				UpdateWorkCache(work.Id)
			}
			defer m.Done()
		}(work)
	}
	m.Wait()

	var workCaches = make([]*WorkCache, 0)
	workCacheMap.Range(func(k, v interface{}) bool {
		// k 为 workName
		if reflect.TypeOf(k).Kind() == reflect.String {
			workCaches = append(workCaches, v.(*WorkCache))
		}
		return true
	})
	return workCaches
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Usage struct {
	UsageMap map[int64][]string // key 为当前步骤 id, value 为当前步骤引用值
	UsedMap  map[int64][]int64  // 被使用统计 map, key 为当前步骤 id, value 为引用当前节点名的步骤 id 数组
}

type WorkCache struct {
	XMLName             xml.Name                                `xml:"workdl"`
	WorkId              int64                                   `xml:"work_id,attr"`
	Work                models.Work                             `xml:"work"`
	Steps               []models.WorkStep                       `xml:"steps"`
	StepsMap            map[int64]*models.WorkStep              `xml:"-"`
	BlockStepOrdersMap  map[int64][]*block.BlockStep            `xml:"-"` // key 为父节点 StepId
	ParamInputSchemaMap map[int64]*iworkmodels.ParamInputSchema `xml:"-"` // key 为 WorkStepId
	SubWorkNameMap      map[int64]string                        `xml:"-"` // key 为 WorkStepId
	Usage               *Usage                                  `xml:"-"` // 引值计算,节点引用值统计
	err                 error                                   `xml:"-"`
	ParamMappings       []iworkmodels.ParamMapping              `xml:"-"`
}

func (this *WorkCache) RenderToString() (s string) {
	bytes, err := xml.MarshalIndent(this, "", "\t")
	if err == nil {
		s = string(bytes)
	}
	return
}

func (this *WorkCache) FlushCache() {
	o := orm.NewOrm()
	// 缓存 work
	this.Work, this.err = models.QueryWorkById(this.WorkId, o)
	checkError(this.err)
	// 缓存 workSteps
	this.Steps, this.err = models.QueryAllWorkStepInfo(this.WorkId, o)
	checkError(this.err)
	this.StepsMap = make(map[int64]*models.WorkStep)
	for _, step := range this.Steps {
		this.StepsMap[step.WorkStepId] = &step
	}

	blockSteps := getBlockStepExecuteOrder(block.GetTopLevelBlockSteps(block.ParseToBlockStep(this.Steps)))
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
		parser := pool.Container.GetSingleton("parser")
		paramInputSchema := parser.(IParamSchemaCacheParser).GetCacheParamInputSchema(&workStep)
		this.ParamInputSchemaMap[workStep.WorkStepId] = paramInputSchema
	}
	// 缓存 subWorkName
	this.SubWorkNameMap = make(map[int64]string, 0)
	for _, workStep := range this.Steps {
		if workStep.WorkStepType == iworkconst.NODE_TYPE_WORK_SUB {
			this.SubWorkNameMap[workStep.WorkStepId], _ = this.getWorkSubName(&workStep)
		}
	}

	this.Usage = &Usage{}
	// 缓存引值计数
	this.cacheReferUsage()

	this.evalParamMapping()
}

func (this *WorkCache) evalParamMapping() {
	for _, step := range this.Steps {
		if step.WorkStepType == iworkconst.NODE_TYPE_WORK_START {
			var ParamMappings []iworkmodels.ParamMapping
			json.Unmarshal([]byte(step.WorkStepParamMapping), &ParamMappings)
			this.ParamMappings = ParamMappings
		}
	}
}

func (this *WorkCache) evalUsageMap() {
	usageMap := make(map[int64][]string) // key 为当前步骤 id, value 为当前步骤引用值
	for workStepId, paramInputSchema := range this.ParamInputSchemaMap {
		relatives := make([]string, 0)
		for _, item := range paramInputSchema.ParamInputSchemaItems {
			// 根据正则找到关联的节点名和字段名
			relativeValues := iworkutil.GetRelativeValueWithReg(item.ParamValue)
			relatives = append(relatives, relativeValues...)
		}
		usageMap[workStepId] = relatives
	}
	this.Usage.UsageMap = usageMap
}

func (this *WorkCache) calUseds(workStepName string) (workStepIds []int64) {
	for workStepId, relatives := range this.Usage.UsageMap {
		for _, relative := range relatives {
			if strings.HasPrefix(relative, fmt.Sprintf(`$%s`, workStepName)) {
				workStepIds = append(workStepIds, workStepId)
				break
			}
		}
	}
	return workStepIds
}

func (this *WorkCache) evalUsedMap() {
	usedMap := make(map[int64][]int64)
	for _, step := range this.Steps {
		usedMap[step.WorkStepId] = this.calUseds(step.WorkStepName)
	}
	this.Usage.UsedMap = usedMap
}

func (this *WorkCache) cacheReferUsage() {
	this.evalUsageMap()
	this.evalUsedMap()
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
