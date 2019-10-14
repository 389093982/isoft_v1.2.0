package iworkbuild

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

// 构建动态输入值
func BuildDynamicInput(step models.WorkStep, o orm.Ormer) {
	workCache, _ := iworkcache.LoadWorkCache(step.WorkId)
	parser := node.ParamSchemaParser{
		WorkStep:          &step,
		ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step, O: o, WorkCache: workCache},
	}
	// 获取默认数据
	defaultParamInputSchema := parser.GetDefaultParamInputSchema()
	// 获取动态数据
	runtimeParamInputSchema := parser.GetRuntimeParamInputSchema()
	// 合并默认数据和动态数据作为新数据
	newInputSchemaItems := append(defaultParamInputSchema.ParamInputSchemaItems, runtimeParamInputSchema.ParamInputSchemaItems...)
	// 获取历史数据
	historyParamInputSchema := parser.GetCacheParamInputSchema()
	for index, newInputSchemaItem := range newInputSchemaItems {
		// 存在则不添加且沿用旧值
		if exist, item := CheckAndGetItemByParamName(historyParamInputSchema.ParamInputSchemaItems, newInputSchemaItem.ParamName); exist {
			newInputSchemaItems[index].ParamValue = item.ParamValue
			newInputSchemaItems[index].PureText = item.PureText
		}
	}
	paramInputSchema := &iworkmodels.ParamInputSchema{ParamInputSchemaItems: newInputSchemaItems}
	step.WorkStepInput = paramInputSchema.RenderToJson()
	if _, err := models.InsertOrUpdateWorkStep(&step, o); err != nil {
		panic(err)
	}
}

func CheckAndGetItemByParamName(items []iworkmodels.ParamInputSchemaItem, paramName string) (bool, *iworkmodels.ParamInputSchemaItem) {
	for _, _item := range items {
		if _item.ParamName == paramName {
			return true, &_item
		}
	}
	return false, nil
}

// 构建动态输出值
func BuildDynamicOutput(step models.WorkStep, o orm.Ormer) { // o 传递的目的是控制在同一个事物中
	parser := node.ParamSchemaParser{WorkStep: &step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step, O: o}}
	runtimeParamOutputSchema := parser.GetRuntimeParamOutputSchema()
	defaultParamOutputSchema := parser.GetDefaultParamOutputSchema()
	defaultParamOutputSchema.ParamOutputSchemaItems = append(defaultParamOutputSchema.ParamOutputSchemaItems, runtimeParamOutputSchema.ParamOutputSchemaItems...)
	// 构建输出参数,使用全新值
	step.WorkStepOutput = defaultParamOutputSchema.RenderToJson()
	if _, err := models.InsertOrUpdateWorkStep(&step, o); err != nil {
		panic(err)
	}
}
