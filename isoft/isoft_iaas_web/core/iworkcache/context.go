package iworkcache

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/models/iwork"
	"sync"
)

var cacheContextMap = make(map[int64]*CacheContext, 0)

var mutex sync.Mutex

type GetBlockStepExecuteOrder func(blockSteps []*block.BlockStep) []*block.BlockStep

func GetCacheContext(work_id int64, orderFunc GetBlockStepExecuteOrder) *CacheContext {
	mutex.Lock()
	if _, ok := cacheContextMap[work_id]; !ok {
		context := &CacheContext{WorkId: work_id}
		context.LoadCache(orderFunc)
		cacheContextMap[work_id] = context
	}
	mutex.Unlock()
	return cacheContextMap[work_id]
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type CacheContext struct {
	WorkId             int64
	Work               iwork.Work
	Steps              []iwork.WorkStep
	BlockStepOrdersMap map[int64][]*block.BlockStep
	err                error
}

func (this *CacheContext) LoadCache(orderFunc GetBlockStepExecuteOrder) {
	o := orm.NewOrm()
	this.Work, this.err = iwork.QueryWorkById(this.WorkId, o)
	checkError(this.err)
	this.Steps, this.err = iwork.QueryAllWorkStepInfo(this.WorkId, o)
	checkError(this.err)
	blockSteps := orderFunc(block.ParseToBlockStep(this.Steps))
	this.BlockStepOrdersMap = map[int64][]*block.BlockStep{
		-1: blockSteps,
	}
	for _, blockStep := range blockSteps {
		this.loadChildrenBlockStepOrders(blockStep, orderFunc)
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
