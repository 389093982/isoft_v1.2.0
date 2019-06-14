package interfaces

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkmodels"
)

// 使用接口来解决循环依赖问题
// 使用接口的好处是传入实现类,而不是引用实现类并创建实例（这样就解决了引用导致的循环依赖问题）
type IParamSchemaParser interface {
	GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema
	GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema
	GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema
	GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema
}

type IWorkStep interface {
	// 节点执行的方法
	Execute(trackingId string)
	IParamSchemaParser
	// 节点定制化校验函数,校验不通过会触发 panic
	ValidateCustom() (checkResult []string)
}

type OrmerProvider interface {
	GetOrmer() orm.Ormer
}

type RunOneStepArgs struct {
	TrackingId string
	Logwriter  *iworklog.CacheLoggerWriter
	Datastore  *datastore.DataStore
	BlockStep  *block.BlockStep
	Dispatcher *entry.Dispatcher
	WorkCache  *iworkcache.WorkCache
}

type RunOneStep func(args *RunOneStepArgs) (receiver *entry.Receiver)
