package node

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil/reflectutil"
	"isoft/isoft_iwork_web/models"
	"reflect"
	"strings"
	"sync"
)

var RegistTypeMap map[string]reflect.Type

func Regist(m map[string]reflect.Type) {
	once := &sync.Once{}
	once.Do(func() {
		RegistTypeMap = m
	})
}

type WorkStepFactory struct {
	Work             *models.Work
	WorkStep         *models.WorkStep                                                                                // 普通步骤执行时使用的参数
	BlockStep        *block.BlockStep                                                                                // 块步骤执行时使用的参数
	WorkSubRunFunc   func(work_id int64, dispatcher *entry.Dispatcher) (trackingId string, receiver *entry.Receiver) // 执行步骤时遇到子流程时的回调函数
	BlockStepRunFunc func(args *interfaces.RunOneStepArgs) (receiver *entry.Receiver)                                // 执行步骤时使用 BlockStep 时的回调函数
	Dispatcher       *entry.Dispatcher
	Receiver         *entry.Receiver // 代理了 Receiver,值从 work_end 节点获取
	DataStore        *datastore.DataStore
	O                orm.Ormer
	LogWriter        *iworklog.CacheLoggerWriter
	WorkCache        *iworkcache.WorkCache
}

func parseToError(err interface{}) error {
	if _err, ok := err.(error); ok {
		return _err
	} else if _err, ok := err.(string); ok {
		return errors.New(_err)
	}
	return errors.New("error...")
}

func (this *WorkStepFactory) Execute(trackingId string) {
	defer func() {
		if err := recover(); err != nil {
			insensitiveErrorFlag := false
			if _err, ok := err.(*interfaces.InsensitiveError); ok {
				insensitiveErrorFlag = true
				err = _err.Error
			}
			wsError := interfaces.WorkStepError{
				Err:          parseToError(err),
				WorkStepName: this.WorkStep.WorkStepName,
			}
			insensitiveErrorMsg := wsError.Error()
			if !insensitiveErrorFlag {
				insensitiveErrorMsg = "InternalError"
			}
			// 将错误写入 Error 中去
			this.DataStore.CacheDatas("Error", map[string]interface{}{
				"isError":             true,
				"isNoError":           false,
				"errorMsg":            wsError.Error(),
				"insensitiveErrorMsg": insensitiveErrorMsg,
			})
			panic(wsError) // 对已经进行包装,异常不吞掉,继续抛出
		}
	}()
	proxy := this.getProxy()
	// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
	proxy.FillParamInputSchemaDataToTmp(this.WorkStep)
	// 存储 pureText 值
	proxy.FillPureTextParamInputSchemaDataToTmp(this.WorkStep)
	// 执行任务
	proxy.Execute(trackingId)
	if receiver := proxy.GetReceiver(); receiver != nil {
		this.Receiver = receiver
	}
}

func GetIWorkStep(workStepType string) interfaces.IWorkStep {
	// 调整 workStepType
	_workStepType := strings.ToUpper(strings.Replace(workStepType, "_", "", -1) + "NODE")
	if t, ok := RegistTypeMap[_workStepType]; ok {
		return reflect.New(t).Interface().(interfaces.IWorkStep)
	}
	panic(fmt.Sprintf("invalid workStepType for %s", workStepType))
}

func (this *WorkStepFactory) getProxy() interfaces.IWorkStep {
	fieldMap := map[string]interface{}{
		"WorkStep": this.WorkStep,
		"BaseNode": BaseNode{
			DataStore:              this.DataStore,
			O:                      this.O,
			LogWriter:              this.LogWriter,
			WorkCache:              this.WorkCache,
			Dispatcher:             this.Dispatcher,
			ParamSchemaCacheParser: &ParamSchemaParser{WorkStep: this.WorkStep, ParamSchemaParser: &WorkStepFactory{WorkStep: this.WorkStep}},
		},
		"Receiver":         this.Receiver,
		"WorkSubRunFunc":   this.WorkSubRunFunc,
		"BlockStep":        this.BlockStep,
		"BlockStepRunFunc": this.BlockStepRunFunc,
	}
	stepNode := GetIWorkStep(this.WorkStep.WorkStepType)
	if stepNode == nil {
		panic(errors.New(fmt.Sprintf("[%v-%v]unsupport workStepType:%s",
			this.WorkStep.WorkId, this.WorkStep.WorkStepName, this.WorkStep.WorkStepType)))
	}
	// 从 map 中找出属性值赋值给对象
	reflectutil.FillFieldValueToStruct(stepNode, fieldMap)
	return stepNode
}

func (this *WorkStepFactory) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	var inputSchema *iworkmodels.ParamInputSchema
	if _schema := this.getProxy().GetDefaultParamInputSchema(); _schema != nil {
		inputSchema = _schema
	} else {
		inputSchema = &iworkmodels.ParamInputSchema{}
	}
	return inputSchema
}

func (this *WorkStepFactory) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	if schema := this.getProxy().GetRuntimeParamInputSchema(); schema != nil {
		return schema
	}
	return &iworkmodels.ParamInputSchema{}
}

func (this *WorkStepFactory) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	if schema := this.getProxy().GetDefaultParamOutputSchema(); schema != nil {
		return schema
	}
	return &iworkmodels.ParamOutputSchema{}
}

func (this *WorkStepFactory) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	if schema := this.getProxy().GetRuntimeParamOutputSchema(); schema != nil {
		return schema
	}
	return &iworkmodels.ParamOutputSchema{}
}

func (this *WorkStepFactory) ValidateCustom() (checkResult []string) {
	return this.getProxy().ValidateCustom()
}

// 获取入参 schema
func GetCacheParamInputSchema(step *models.WorkStep) *iworkmodels.ParamInputSchema {
	parser := ParamSchemaParser{WorkStep: step, ParamSchemaParser: &WorkStepFactory{WorkStep: step}}
	return parser.GetCacheParamInputSchema()
}

func GetCacheParamOutputSchema(step *models.WorkStep) *iworkmodels.ParamOutputSchema {
	parser := ParamSchemaParser{WorkStep: step, ParamSchemaParser: &WorkStepFactory{WorkStep: step}}
	return parser.GetCacheParamOutputSchema()
}
