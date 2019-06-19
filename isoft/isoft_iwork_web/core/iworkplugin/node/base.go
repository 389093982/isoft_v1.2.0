package node

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/interfaces"
	"isoft/isoft_iwork_web/core/iworkplugin/params"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/models"
	"sort"
	"strings"
)

// 所有 node 的基类
type BaseNode struct {
	interfaces.IWorkStep
	DataStore          *datastore.DataStore
	O                  orm.Ormer
	LogWriter          *iworklog.CacheLoggerWriter
	WorkCache          *iworkcache.WorkCache
	TmpDataMap         map[string]interface{}
	PureTextTmpDataMap map[string]interface{}
	Dispatcher         *entry.Dispatcher
}

func (this *BaseNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	fmt.Println("execute default GetDefaultParamInputSchema method...")
	return &iworkmodels.ParamInputSchema{}
}

func (this *BaseNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	fmt.Println("execute default GetRuntimeParamInputSchema method...")
	return &iworkmodels.ParamInputSchema{}
}

func (this *BaseNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	fmt.Println("execute default GetDefaultParamOutputSchema method...")
	return &iworkmodels.ParamOutputSchema{}
}

func (this *BaseNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	fmt.Println("execute default GetRuntimeParamOutputSchema method...")
	return &iworkmodels.ParamOutputSchema{}
}

func (this *BaseNode) ValidateCustom() (checkResult []string) {
	fmt.Println("execute default ValidateCustom method...")
	return
}

func (this *BaseNode) GetReceiver() *entry.Receiver {
	return nil
}

// 存储 pureText 值
func (this *BaseNode) FillPureTextParamInputSchemaDataToTmp(workStep *models.WorkStep) {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	parser := schema.WorkStepFactoryParamSchemaParser{WorkStep: workStep, ParamSchemaParser: &WorkStepFactory{WorkStep: workStep}}
	paramInputSchema := parser.GetCacheParamInputSchema()
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// tmpDataMap 存储引用值 pureText
		tmpDataMap[item.ParamName] = item.ParamValue
	}
	this.PureTextTmpDataMap = tmpDataMap
}

// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
func (this *BaseNode) FillParamInputSchemaDataToTmp(workStep *models.WorkStep) {

	// dispatcher 非空时替换成父流程参数
	if this.Dispatcher != nil && len(this.Dispatcher.TmpDataMap) > 0 {
		tmpDataMap := make(map[string]interface{})
		// 从父流程中获取值,即从 Dispatcher 中获取值
		for key, value := range this.Dispatcher.TmpDataMap {
			if value != "__default__" { // __default__ 则表示不用替换,还是使用子流程默认值参数
				tmpDataMap[key] = value
			}
		}
		this.TmpDataMap = tmpDataMap
	} else {
		pis := this.WorkCache.ParamInputSchemaMap[workStep.WorkStepId]
		this.TmpDataMap = params.FillParamInputSchemaDataToTmp(pis, this.DataStore)
	}
}

// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
func (this *BaseNode) SubmitParamOutputSchemaDataToDataStore(workStep *models.WorkStep, dataStore *datastore.DataStore, tmpDataMap map[string]interface{}) {
	parser := schema.WorkStepFactoryParamSchemaParser{WorkStep: workStep, ParamSchemaParser: &WorkStepFactory{WorkStep: workStep}}
	paramOutputSchema := parser.GetCacheParamOutputSchema()
	paramMap := make(map[string]interface{})
	for _, item := range paramOutputSchema.ParamOutputSchemaItems {
		paramMap[item.ParamName] = tmpDataMap[item.ParamName]
	}
	// 将数据数据存储到数据中心
	dataStore.CacheDatas(workStep.WorkStepName, paramMap)
}

// 根据传入的 paramMap 构建 ParamInputSchema 对象
func (this *BaseNode) BuildParamInputSchemaWithDefaultMap(paramMap map[int][]string) *iworkmodels.ParamInputSchema {
	keys := datatypeutil.GetMapKeySlice(paramMap, []int{}).([]int)
	sort.Ints(keys)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, key := range keys {
		_paramMap := paramMap[key]
		// 前两位分别是名称和描述
		paramName := _paramMap[0]
		paramDesc := _paramMap[1]
		item := iworkmodels.ParamInputSchemaItem{ParamName: paramName, ParamDesc: paramDesc}
		// 后面字段为 extra 字段
		for _, paramExtra := range _paramMap[1:] {
			if strings.HasPrefix(paramExtra, "repeatable__") {
				item.Repeatable = true
				item.ForeachRefer = strings.Replace(paramExtra, "repeatable__", "", 1)
			}
		}
		items = append(items, item)
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

// 根据传入的 paramNames 构建 ParamOutputSchema 对象
func (this *BaseNode) BuildParamOutputSchemaWithSlice(paramNames []string) *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, paramName := range paramNames {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: paramName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
