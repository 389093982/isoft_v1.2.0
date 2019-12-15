package file

import (
	"io/ioutil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/core/iworkutil/stringutil"
	"isoft/isoft_iwork_web/models"
	"os"
	"path"
	"strings"
)

type FileReadNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *FileReadNode) Execute(trackingId string) {
	paramMap := make(map[string]interface{}, 0)
	file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	paramMap[iworkconst.STRING_PREFIX+"file_path"] = file_path
	if bytes, err := ioutil.ReadFile(file_path); err == nil {
		paramMap[iworkconst.STRING_PREFIX+"data"] = string(bytes)
	} else {
		panic(err)
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *FileReadNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "读取文件的绝对路径"},
	}
	return this.BPIS1(paramMap)
}

func (this *FileReadNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "file_path", "data"})
}

type FileWriteNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func checkAppend(tmpDataMap map[string]interface{}) bool {
	if append, ok := tmpDataMap[iworkconst.BOOL_PREFIX+"append?"].(string); ok && strings.TrimSpace(append) != "" {
		return true
	}
	return false
}

func (this *FileWriteNode) Execute(trackingId string) {
	file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	var strdata string
	// 写字符串
	if data, ok := this.TmpDataMap[iworkconst.STRING_PREFIX+"data?"].(string); ok {
		strdata = data
	}
	// 写字节数组
	if bytes, ok := this.TmpDataMap[iworkconst.BYTE_ARRAY_PREFIX+"data?"].([]byte); ok {
		strdata = string(bytes)
	}
	// 判断是否需要添加行分隔符
	if linesep, ok := this.TmpDataMap[iworkconst.BOOL_PREFIX+"linesep?"].(string); ok && strings.TrimSpace(linesep) != "" {
		strdata += "\n"
	}
	if err := fileutil.WriteFile(file_path, []byte(strdata), checkAppend(this.TmpDataMap)); err != nil {
		panic(err)
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "file_path": file_path})
}

func (this *FileWriteNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "写入文件的绝对路径,文件不存在时会自动创建"},
		2: {iworkconst.STRING_PREFIX + "data?", "可选参数,写入文件的字符数据"},
		3: {iworkconst.BYTE_ARRAY_PREFIX + "data?", "可选参数,写入文件的二进制字节数据"},
		4: {iworkconst.BOOL_PREFIX + "append?", "可选参数,文件追加模式,值为空表示覆盖,有值表示追加"},
		5: {iworkconst.BOOL_PREFIX + "linesep?", "可选参数,行分隔符,默认没有分割符,有值表示使用换行符进行分割"},
	}
	return this.BPIS1(paramMap)
}

func (this *FileWriteNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "file_path"})
}

type FileSyncNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *FileSyncNode) Execute(trackingId string) {
	sync_mod := stringutil.GetString(this.TmpDataMap[iworkconst.STRING_PREFIX+"sync_mod?"], "copy", true)
	file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	new_file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"new_file_path"].(string)
	var err error
	if sync_mod == "copy" {
		err = fileutil.CopyFile(file_path, new_file_path)
	} else if sync_mod == "rename" {
		err = os.Rename(file_path, new_file_path)
	}
	if err == nil {
		paramMap := map[string]interface{}{
			"file_path":     new_file_path,
			"new_file_name": path.Base(new_file_path),
			"new_file_ext":  path.Ext(new_file_path),
		}
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
	} else {
		panic(err)
	}
}

func (this *FileSyncNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sync_mod?", "文件同步的策略,支持拷贝重命名和移动重命名(copy、rename),默认是 copy"},
		2: {iworkconst.STRING_PREFIX + "file_path", "需要进行同步操作的文件路径"},
		3: {iworkconst.STRING_PREFIX + "new_file_path", "同步操作后的文件路径"},
	}
	return this.BPIS1(paramMap)
}

func (this *FileSyncNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{"new_file_path", "new_file_name", "new_file_ext"})
}

type FileDeleteNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *FileDeleteNode) Execute(trackingId string) {
	delete_file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"delete_file_path"].(string)
	err := os.RemoveAll(delete_file_path)
	if err != nil {
		panic(err)
	}
}

func (this *FileDeleteNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "delete_file_path", "待删除的文件或文件夹路径"},
	}
	return this.BPIS1(paramMap)
}
