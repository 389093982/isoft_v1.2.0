package framework

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type DefineVarNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DefineVarNode) Execute(trackingId string) {

}

func (this *DefineVarNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "define_vars", "待定义的变量"},
	}
	return this.BPIS1(paramMap)
}

func (this *DefineVarNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{}
	define_vars := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"define_vars", this.WorkStep).(string)
	if define_vars != "" {
		items := make([]iworkmodels.ParamOutputSchemaItem, 0)
		for _, define_var := range strings.Split(define_vars, ",") {
			items = append(items, iworkmodels.ParamOutputSchemaItem{
				ParamName: strings.TrimSpace(define_var),
			})
		}
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, items...)
	}
	return pos
}
