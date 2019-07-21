package framework

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type TemplateNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *TemplateNode) Execute(trackingId string) {

}

func (this *TemplateNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "template_text", "模板文字"},
		2: {iworkconst.STRING_PREFIX + "template_varName", "模板变量名称"},
		3: {iworkconst.STRING_PREFIX + "template_varValue", "模板变量值"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *TemplateNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "template_text"})
}
