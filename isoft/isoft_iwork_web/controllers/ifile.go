package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
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
		errorMsg  string
		err       error
	)
	work_name := this.Ctx.Input.Param(":work_name")
	if work_name != "default" {
		if workCache, err = iworkcache.GetWorkCacheWithName(work_name); err != nil {
			panic(errors.New(fmt.Sprintf("处理流程 %s 不存在!", work_name)))
		}
	}
	tempFileName, fileName, tempFilePath := this.saveFile()
	tempFileServerPath := "http://localhost:8086/api/files/" + tempFileName
	if workCache != nil {
		mapData := ParseParam(this.Ctx, workCache.Steps)
		mapData[iworkconst.HTTP_REQUEST_OBJECT] = this.Ctx.Request // 传递 request 对象
		mapData["__tempFileName"] = tempFileName
		mapData["__fileName"] = fileName
		mapData["__fileExt"] = path.Ext(fileName)
		mapData["__tempFilePath"] = tempFilePath
		mapData["__tempFileServerPath"] = tempFileServerPath
		// 调度流程进行处理
		trackingId, receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: mapData})
		this.Ctx.ResponseWriter.Header().Add(iworkconst.TRACKING_ID, trackingId)
		if receiver != nil {
			tempDataMap := receiver.TmpDataMap
			if errorMsg, ok := tempDataMap["errorMsg"]; ok && errorMsg != nil && errorMsg.(string) != "" {
				panic(errors.New(errorMsg.(string)))
			}
			if data, ok := tempDataMap[iworkconst.DO_RESPONSE_RECEIVE_FILE]; ok {
				tmpDataMap := data.(map[string]interface{})
				tempFileName = tmpDataMap["fileName"].(string) // 将临时文件的数据刷新成正式数据
				tempFileServerPath = tmpDataMap["fileServerPath"].(string)
			}
		}
	}
	this.Data["json"] = &map[string]interface{}{
		"status":         "SUCCESS",
		"filename":       tempFileName,
		"fileServerPath": tempFileServerPath,
		"errorMsg":       errorMsg,
	}
	this.ServeJSON()
}

func (this *WorkController) saveFile() (tempFileName, fileName, tempFilePath string) {
	tempFileName = stringutil.RandomUUID()
	// 判断是否是文件上传
	f, h, err := this.GetFile("file")
	checkError(err)
	//得到文件的名称
	fileName = h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置,static/uploadfile,这个是文件的地址,第一个static前面不要有/
	tempFilePath = path.Join(beego.AppConfig.String("file.server"), tempFileName)
	err = this.SaveToFile("file", tempFilePath)
	checkError(err)
	return tempFileName, fileName, tempFilePath
}
