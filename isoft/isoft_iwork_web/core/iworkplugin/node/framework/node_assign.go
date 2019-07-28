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
	//for paramName, paramValue := range this.TmpDataMap {
	//	if strings.HasSuffix(paramName, "_operate") {
	//		_paramName := paramName[:strings.LastIndex(paramName, "_operate")]
	//		assignVar := this.TmpDataMap[_paramName]
	//		assignOperate := paramValue.(string)
	//		assignVal := this.TmpDataMap[_paramName+"_value"]
	//
	//		// pureText
	//		assignVar_pureText := this.PureTextTmpDataMap[_paramName].(string)
	//		assignNodeName, assignDataName := parseAssignRefer(_paramName, assignVar_pureText)
	//
	//		assign := NodeAssign{
	//			AssignVar:     assignVar,
	//			AssignOperate: assignOperate,
	//			AssignData:    assignVal,
	//		}
	//		assignVar, err := assign.Calculate()
	//		if err != nil {
	//			panic(err)
	//		}
	//		// 重新将值绑定到对应的 assign 节点
	//		this.DataStore.CacheDatas(assignNodeName, map[string]interface{}{
	//			assignDataName: assignVar,
	//		})
	//	}
	//}
}

func (this *AssignVarNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "assign_obj", "待赋值的对象"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *AssignVarNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	assign_obj := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"assign_obj", this.WorkStep).(string)
	assign_obj_name := assign_obj[strings.LastIndex(assign_obj, "$")+1 : strings.LastIndex(assign_obj, ";")]
	for _, step := range this.BaseNode.WorkCache.Steps {
		if step.WorkStepName == assign_obj_name {
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
