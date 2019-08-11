package file

import (
	"isoft/isoft_iwork_web/core/iworkconst"
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
		"fileName":           this.Dispatcher.TmpDataMap["__fileName"],
		"tempFileName":       this.Dispatcher.TmpDataMap["__tempFileName"],
		"fileExt":            this.Dispatcher.TmpDataMap["__fileExt"],
		"tempFilePath":       this.Dispatcher.TmpDataMap["__tempFilePath"],
		"tempFileServerPath": this.Dispatcher.TmpDataMap["__tempFileServerPath"],
	})
}

func (this *DoReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: make([]iworkmodels.ParamInputSchemaItem, 0)}
}

func (this *DoReceiveFileNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{"fileName", "tempFileName", "fileExt", "tempFilePath", "tempFileServerPath"})
}

type DoResponseReceiveFileNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoResponseReceiveFileNode) Execute(trackingId string) {
	this.TmpDataMap["fileName"] = this.TmpDataMap["fileName"]
	this.TmpDataMap["fileServerPath"] = this.TmpDataMap["fileServerPath"]
	this.TmpDataMap["errorMsg"] = this.TmpDataMap["errorMsg?"]
	this.DataStore.CacheDatas(iworkconst.DO_RESPONSE_RECEIVE_FILE, map[string]interface{}{iworkconst.DO_RESPONSE_RECEIVE_FILE: this.TmpDataMap})
	// 直接提交 dataStore
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, this.TmpDataMap)
}

func (this *DoResponseReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {"fileName", "最终的上传文件名称"},
		2: {"fileServerPath", "最终的服务器地址"},
		3: {"errorMsg?", "异常信息"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *DoResponseReceiveFileNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{"fileName", "fileServerPath", "errorMsg"})
}
