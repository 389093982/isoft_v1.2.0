package framework

import (
	"encoding/json"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type MapperNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *MapperNode) Execute(trackingId string) {
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END、Mapper 等节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, this.TmpDataMap)
}

func (this *MapperNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []iworkmodels.ParamMapping
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping.ParamMappingName})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *MapperNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	inputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema()
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
