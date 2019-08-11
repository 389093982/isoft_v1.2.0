package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"path"
	"strings"
)

func (this *WorkController) FileUpload() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]string{"status": "ERROR", "errorMsg": errorutil.ToError(err).Error()}
			this.ServeJSON()
		}
	}()
	var (
		workCache *iworkcache.WorkCache
		err       error
	)
	work_name := this.Ctx.Input.Param(":work_name")
	if work_name != "default" {
		if workCache, err = iworkcache.GetWorkCacheWithName(work_name); err != nil {
			panic(errors.New(fmt.Sprintf("处理流程 %s 不存在!", work_name)))
		}
	}
	fileName, filePath := this.saveFile()
	tmpDataMap := map[string]interface{}{
		"filename": fileName,
		"filepath": filePath,
	}
	if workCache != nil {
		// 调度流程进行处理
		trackingId, receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: tmpDataMap})
		this.Ctx.ResponseWriter.Header().Add(iworkconst.TRACKING_ID, trackingId)
		if receiver != nil {
			fmt.Println(receiver)
		}
	}
	this.Data["json"] = &map[string]interface{}{
		"status":   "SUCCESS",
		"filename": fileName,
		"filepath": "http://localhost:8086/api/files/" + fileName,
	}
	this.ServeJSON()
}

func (this *WorkController) saveFile() (string, string) {
	// 判断是否是文件上传
	f, h, err := this.GetFile("file")
	checkError(err)
	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置,static/uploadfile,这个是文件的地址,第一个static前面不要有/
	filePath := path.Join(beego.AppConfig.String("file.server"), fileName)
	err = this.SaveToFile("file", filePath)
	checkError(err)
	return fileName, filePath
}
