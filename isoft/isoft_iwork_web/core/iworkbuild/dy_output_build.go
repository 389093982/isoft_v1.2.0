package iworkbuild

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

// 构建动态输出值
func BuildDynamicOutput(step models.WorkStep, o orm.Ormer) {
	parser := node.ParamSchemaParser{WorkStep: &step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step}}
	runtimeParamOutputSchema := parser.GetRuntimeParamOutputSchema()
	defaultParamOutputSchema := parser.GetDefaultParamOutputSchema()
	defaultParamOutputSchema.ParamOutputSchemaItems = append(defaultParamOutputSchema.ParamOutputSchemaItems, runtimeParamOutputSchema.ParamOutputSchemaItems...)
	// 构建输出参数,使用全新值
	step.WorkStepOutput = defaultParamOutputSchema.RenderToJson()
	if _, err := models.InsertOrUpdateWorkStep(&step, o); err != nil {
		panic(err)
	}
}
