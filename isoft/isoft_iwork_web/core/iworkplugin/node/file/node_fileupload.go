package file

import (
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"path"
	"strings"
)

type DoReceiveFileNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoReceiveFileNode) Execute(trackingId string) {
	fileUpload := this.Dispatcher.TmpDataMap[iworkconst.HTTP_REQUEST_IFILE_UPLOAD].(interfaces.IFileUpload)
	suffixStr := this.TmpDataMap[iworkconst.STRING_PREFIX+"suffixs"].(string)
	// fileServerAddr := "http://localhost:8086/api/files/"
	fileServerAddr := this.TmpDataMap[iworkconst.STRING_PREFIX+"fileServerAddr"].(string)
	suffixs := strings.Split(suffixStr, ",")
	tempFileName, fileName, tempFilePath := fileUpload.SaveFile(suffixs)
	paramMap := map[string]interface{}{
		"fileName":       fileName,
		"tempFileName":   tempFileName,
		"fileExt":        path.Ext(fileName),
		"tempFilePath":   tempFilePath,
		"fileServerAddr": fileServerAddr,
	}
	if calHash := this.TmpDataMap[iworkconst.BOOL_PREFIX+"calHash?"].(string); calHash == "true" {
		hash, _ := hashutil.CalculateHashWithFile(tempFilePath)
		paramMap["hash"] = hash
		paramMap["handleSpecialHash"], _ = stringutil.ReplaceAllString(hash, "/", "-")
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *DoReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "calHash?", "是否计算hash值"},
		2: {iworkconst.STRING_PREFIX + "suffixs", "上传文件支持的后缀名,*表示支持任意类型的后缀,多个后缀用逗号分隔"},
		3: {iworkconst.STRING_PREFIX + "fileServerAddr", "上传文件服务器访问路径"},
	}
	choiceMap := map[string][]string{iworkconst.BOOL_PREFIX + "calHash?": {"`true`", "`false`"}}
	return this.BPIS1(paramMap, choiceMap)
}

func (this *DoReceiveFileNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{"fileName", "tempFileName", "fileExt", "tempFilePath", "fileServerAddr"})
}

func (this *DoReceiveFileNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{}
	calHash := param.GetStaticParamValueWithStep(iworkconst.BOOL_PREFIX+"calHash?", this.WorkStep).(string)
	if calHash == "true" || calHash == "`true`" {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{
			ParamName: "hash",
		})
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{
			ParamName: "handleSpecialHash",
		})
	}
	return pos
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
	return this.BPIS1(paramMap)
}

func (this *DoResponseReceiveFileNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{"fileName", "fileServerPath", "errorMsg"})
}
