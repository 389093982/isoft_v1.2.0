package framework

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type AssignVarNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *AssignVarNode) Execute(trackingId string) {
	assignNodeName := this.getAssignNodeName()
	paramMap := make(map[string]interface{})
	for paramName, paramValue := range this.TmpDataMap {
		if paramName != iworkconst.STRING_PREFIX+"assign_node" {
			paramMap[strings.Replace(paramName, "?", "", -1)] = paramValue
		}
	}
	// 重新将值绑定到对应的 assign 节点
	this.DataStore.CacheDatas(assignNodeName, paramMap)
}

func (this *AssignVarNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "assign_node", "待赋值的对象"},
	}
	return this.BPIS1(paramMap)
}

func (this *AssignVarNode) getAssignNodeName() string {
	assign_node := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"assign_node", this.WorkStep).(string)
	assign_node_name := assign_node[strings.LastIndex(assign_node, "$")+1 : strings.LastIndex(assign_node, ";")]
	return assign_node_name
}

func (this *AssignVarNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	assignNodeName := this.getAssignNodeName()
	for _, step := range this.BaseNode.WorkCache.Steps {
		if step.WorkStepName == assignNodeName {
			if paramOutputSchema, err := iworkmodels.ParseToParamOutputSchema(step.WorkStepOutput); err == nil {
				for _, item := range paramOutputSchema.ParamOutputSchemaItems {
					items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: item.ParamName + "?"})
				}
			}
			break
		}
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}
