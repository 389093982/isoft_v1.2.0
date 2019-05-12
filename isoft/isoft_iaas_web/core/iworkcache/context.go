package iworkcache

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkplugin/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"sync"
)

var cacheContextMap = make(map[int64]*CacheContext, 0)

var mutex sync.Mutex

func GetCacheContext(work_id int64) *CacheContext {
	mutex.Lock()
	if _, ok := cacheContextMap[work_id]; !ok {
		context := &CacheContext{WorkId: work_id}
		context.LoadCache()
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
	WorkId          int64
	Work            iwork.Work
	Steps           []iwork.WorkStep
	BlockStepOrders []*block.BlockStep
	err             error
}

func (this *CacheContext) LoadCache() {
	o := orm.NewOrm()
	this.Work, this.err = iwork.QueryWorkById(this.WorkId, o)
	checkError(this.err)
	this.Steps, this.err = iwork.QueryAllWorkStepInfo(this.WorkId, o)
	checkError(this.err)
	this.BlockStepOrders = iworknode.GetBlockStepExecuteOrder(block.ParseToBlockStep(this.Steps))
}
