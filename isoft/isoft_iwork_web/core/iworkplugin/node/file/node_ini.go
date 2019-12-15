package file

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
)

type IniReadNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *IniReadNode) Execute(trackingId string) {
	file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	var section_name string
	if _section_name, ok := this.TmpDataMap[iworkconst.STRING_PREFIX+"section_name?"].(string); ok {
		section_name = _section_name
	}
	key := this.TmpDataMap[iworkconst.STRING_PREFIX+"key"].(string)
	value, err := fileutil.ReadBeegoIniFile(file_path, section_name, key)
	if err == nil {
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "value": value})
	} else {
		panic(err)
	}
}

func (this *IniReadNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "需要读取的文件路径"},
		2: {iworkconst.STRING_PREFIX + "section_name?", "section 名称,可为空"},
		3: {iworkconst.STRING_PREFIX + "key", "key 值"},
	}
	return this.BPIS1(paramMap)
}

func (this *IniReadNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "value"})
}

type IniWriteNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *IniWriteNode) Execute(trackingId string) {
	file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	var section_name string
	if _section_name, ok := this.TmpDataMap[iworkconst.STRING_PREFIX+"section_name?"].(string); ok {
		section_name = _section_name
	}
	key := this.TmpDataMap[iworkconst.STRING_PREFIX+"key"].(string)
	value := this.TmpDataMap[iworkconst.STRING_PREFIX+"value"].(string)
	err := fileutil.WriteBeegoIniFile(file_path, section_name, key, value)
	if err != nil {
		panic(err)
	}
}

func (this *IniWriteNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "需要读取的文件路径"},
		2: {iworkconst.STRING_PREFIX + "section_name?", "section 名称,可为空"},
		3: {iworkconst.STRING_PREFIX + "key", "key 值"},
		4: {iworkconst.STRING_PREFIX + "value", "value 值"},
	}
	return this.BPIS1(paramMap)
}
