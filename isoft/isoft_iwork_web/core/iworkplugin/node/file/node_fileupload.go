package file

import (
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type DoReceiveFileNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoReceiveFileNode) Execute(trackingId string) {
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{
		"filename": this.Dispatcher.TmpDataMap["__filename"],
		"fileExt":  this.Dispatcher.TmpDataMap["__fileExt"],
		"filepath": this.Dispatcher.TmpDataMap["__filepath"],
	})
}

func (this *DoReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: make([]iworkmodels.ParamInputSchemaItem, 0)}
}

func (this *DoReceiveFileNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{"filename", "fileExt", "filepath"})
}

type DoResponseReceiveFileNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoResponseReceiveFileNode) Execute(trackingId string) {
	paramMap := make(map[string]interface{}, 0)
	paramMap["filename"] = this.TmpDataMap["filename"]
	paramMap["fileServerPath"] = this.TmpDataMap["fileServerPath"]
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *DoResponseReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {"filename", "上传文件名称"},
		2: {"fileServerPath", "服务器地址"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}
