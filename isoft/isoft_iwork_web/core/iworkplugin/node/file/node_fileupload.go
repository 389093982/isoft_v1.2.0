package file

import (
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"path"
)

type DoReceiveFileNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoReceiveFileNode) Execute(trackingId string) {
	fileUpload := this.Dispatcher.TmpDataMap[iworkconst.HTTP_REQUEST_IFILE_UPLOAD].(interfaces.IFileUpload)
	tempFileName, fileName, tempFilePath := fileUpload.SaveFile()
	tempFileServerPath := "http://localhost:8086/api/files/" + tempFileName
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{
		"fileName":           fileName,
		"tempFileName":       tempFileName,
		"fileExt":            path.Ext(fileName),
		"tempFilePath":       tempFilePath,
		"tempFileServerPath": tempFileServerPath,
	})
}

func (this *DoReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "calHash?", "是否计算hash值"},
	}
	choiceMap := map[string][]string{iworkconst.BOOL_PREFIX + "calHash?": {"`true`", "`false`"}}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap, choiceMap)
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
