package node

import (
	"isoft/isoft/common/fileutil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil/sftputil"
	"isoft/isoft_iwork_web/models/iwork"
	"path/filepath"
)

type SftpUploadNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SftpUploadNode) Execute(trackingId string) {
	sftpResource := this.TmpDataMap[iworkconst.STRING_PREFIX+"sftp_conn"].(iwork.Resource)
	local_file_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"local_file_path"].(string)
	remote_dir_path := this.TmpDataMap[iworkconst.STRING_PREFIX+"remote_dir_path"].(string)
	err := sftputil.SFTPFileCopy(sftpResource.ResourceUsername, sftpResource.ResourcePassword, sftpResource.ResourceDsn, 22, local_file_path, remote_dir_path)
	if err == nil {
		remote_file_path := fileutil.ChangeToLinuxSeparator(filepath.Join(remote_dir_path, filepath.Base(local_file_path)))
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "remote_file_path": remote_file_path})
	} else {
		panic(err)
	}
}

func (this *SftpUploadNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sftp_conn", "sftp连接信息,需要使用 $RESOURCE 全局参数"},
		2: {iworkconst.STRING_PREFIX + "local_file_path", "本地文件路径"},
		3: {iworkconst.STRING_PREFIX + "remote_dir_path", "远程文件夹路径"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SftpUploadNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "remote_file_path"})
}
