package node

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/models/iwork"
)

type DefineVarNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *DefineVarNode) Execute(trackingId string) {
	dataMap := make(map[string]interface{}, 0)
	for paramName, paramType := range this.TmpDataMap {
		var paramValue interface{}
		switch paramType {
		case `string`:
			paramValue = ""
		case `interface{}`:
			paramValue = new(interface{})
		case `[]interface{}`:
			paramValue = make([]interface{}, 0)
		case `map[string]interface{}`:
			paramValue = make(map[string]interface{}, 0)
		default:
			panic(errors.New(fmt.Sprintf("unsupport paramType for %s", paramType)))
		}
		dataMap[paramName] = paramValue
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, dataMap)
}

func (this *DefineVarNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []iworkmodels.ParamMapping
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{
			ParamName: paramMapping.ParamMappingName,
			ParamChoices: []string{
				"`string`",
				"`interface{}`",
				"`[]interface{}`",
				"`map[string]interface{}`",
			},
		})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *DefineVarNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	parser := schema.WorkStepFactorySchemaParser{WorkStep: this.WorkStep, ParamSchemaParser: &WorkStepFactory{WorkStep: this.WorkStep}}
	inputSchema := parser.GetCacheParamInputSchema()
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
