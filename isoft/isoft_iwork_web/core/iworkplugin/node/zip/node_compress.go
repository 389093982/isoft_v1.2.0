package zip

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/compressutil"
	"isoft/isoft_iwork_web/models"
)

type TarGzUnCompressNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *TarGzUnCompressNode) Execute(trackingId string) {
	targz_file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"targz_file_path"].(string)
	dest_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"dest_dir_path"].(string)
	if err := compressutil.DeCompress(targz_file_path, dest_path); err == nil {
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "dest_dir_path": dest_path})
	} else {
		panic(err)
	}
}

func (this *TarGzUnCompressNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "targz_file_path", "targz 文件路径"},
		2: {iworkconst.STRING_PREFIX + "dest_dir_path", "解压后的文件夹路径"},
	}
	return this.BPIS1(paramMap)
}

func (this *TarGzUnCompressNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "dest_dir_path"})
}

type TarGzCompressNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *TarGzCompressNode) Execute(trackingId string) {
	dir_file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"src_dir_path"].(string)
	dest_file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"dest_file_path"].(string)
	if err := compressutil.CompressDir(dir_file_path, dest_file_path); err == nil {
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "dest_file_path": dest_file_path})
	} else {
		panic(err)
	}
}

func (this *TarGzCompressNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "src_dir_path", "待压缩的文件夹路径"},
		2: {iworkconst.STRING_PREFIX + "dest_file_path", "压缩后的targz文件路径"},
	}
	return this.BPIS1(paramMap)
}

func (this *TarGzCompressNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "dest_file_path"})
}
