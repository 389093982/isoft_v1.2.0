package framework

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type WorkStartNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *WorkStartNode) Execute(trackingId string) {
	fillInfo := make([]string, 0)
	for key, value := range this.TmpDataMap {
		fillInfo = append(fillInfo, fmt.Sprintf("fill param for %s:%s", key, value))
	}
	this.LogWriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, strings.Join(fillInfo, "<br/>"))
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END、Mapper 等节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, this.TmpDataMap)
}

func (this *WorkStartNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []iworkmodels.ParamMapping
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping.ParamMappingName})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkStartNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	inputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema()
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

type WorkEndNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
	Receiver *entry.Receiver
}

func (this *WorkEndNode) Execute(trackingId string) {
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, this.TmpDataMap)
	this.fillExtraTmpDataMap()
	// 同时需要将数据提交到 Receiver
	this.Receiver = &entry.Receiver{TmpDataMap: this.TmpDataMap}
}

func (this *WorkEndNode) fillExtraTmpDataMap() {
	if doErrorFilter := this.DataStore.GetData(iworkconst.DO_ERROR_FILTER, iworkconst.DO_ERROR_FILTER); doErrorFilter != nil {
		this.TmpDataMap[iworkconst.DO_ERROR_FILTER] = doErrorFilter
	}
	if doResoponseReceiveFile := this.DataStore.GetData(iworkconst.DO_RESPONSE_RECEIVE_FILE, iworkconst.DO_RESPONSE_RECEIVE_FILE); doResoponseReceiveFile != nil {
		this.TmpDataMap[iworkconst.DO_RESPONSE_RECEIVE_FILE] = doResoponseReceiveFile
	}
}

func (this *WorkEndNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []iworkmodels.ParamMapping
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping.ParamMappingName})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkEndNode) GetReceiver() *entry.Receiver {
	return this.Receiver
}

func (this *WorkEndNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	inputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema()
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
