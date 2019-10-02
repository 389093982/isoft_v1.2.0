package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/models"
	"path"
)

// 示例地址: http://localhost:8086/api/iwork/httpservice/test_iblog_table_migrate?author0=admin1234567
func (this *WorkController) PublishSerivce() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{
				"status":   "ERROR",
				"errorMsg": err.(error).Error(),
			}
			this.ServeJSON()
		}
	}()
	work_name := this.Ctx.Input.Param(":work_name")
	workCache, err := iworkcache.GetWorkCacheWithName(work_name)
	checkError(err)
	mapData := ParseParam(this.Ctx, workCache.Steps)
	mapData[iworkconst.HTTP_REQUEST_OBJECT] = this.Ctx.Request // 传递 request 对象
	mapData[iworkconst.HTTP_REQUEST_IFILE_UPLOAD] = this
	trackingId, receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: mapData})
	if receiver != nil {
		receiver.TmpDataMap[iworkconst.TRACKING_ID] = trackingId
		this.ResponseUploadFile(receiver)
		this.Data["json"] = &receiver.TmpDataMap
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", iworkconst.TRACKING_ID: trackingId, "errorMsg": "Empty Response"}
	}
	this.ServeJSON()
}

func (this *WorkController) ResponseUploadFile(receiver *entry.Receiver) {
	tempDataMap := receiver.TmpDataMap
	if data, ok := tempDataMap[iworkconst.DO_RESPONSE_RECEIVE_FILE]; ok {
		tmpDataMap := data.(map[string]interface{})
		receiver.TmpDataMap["fileName"] = tmpDataMap["fileName"].(string) // 将临时文件的数据刷新成正式数据
		receiver.TmpDataMap["fileServerPath"] = tmpDataMap["fileServerPath"].(string)
		receiver.TmpDataMap["status"] = "SUCCESS"
		if errorMsg, ok := tmpDataMap["errorMsg?"].(string); ok {
			receiver.TmpDataMap["errorMsg"] = errorMsg
			receiver.TmpDataMap["status"] = "ERROR"
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseParam(ctx *context.Context, steps []models.WorkStep) map[string]interface{} {
	mapData := map[string]interface{}{}
	for _, step := range steps {
		if step.WorkStepType == iworkconst.NODE_TYPE_WORK_START {
			inputSchema := node.GetCacheParamInputSchema(&step)
			for _, item := range inputSchema.ParamInputSchemaItems {
				// 默认参数类型都当成 string 类型
				mapData[item.ParamName] = ctx.Input.Query(item.ParamName) // 传递参数允许为空串
			}
			break
		}
	}
	return mapData
}

func (this *WorkController) SaveFile(suffixs []string) (tempFileName, fileName, tempFilePath string) {
	// 判断是否是文件上传
	f, h, err := this.GetFile("file")
	checkError(err)
	if !stringutil.AnyOf("*", suffixs) && !stringutil.AnyOf(path.Ext(h.Filename), suffixs) {
		panic("check upload file suffix error!")
	}
	tempFileName = stringutil.RandomUUID() + path.Ext(h.Filename)
	//得到文件的名称
	fileName = h.Filename
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer f.Close()
	//保存文件到指定的位置,static/uploadfile,这个是文件的地址,第一个static前面不要有/
	tempFilePath = path.Join(beego.AppConfig.String("file.server"), tempFileName)
	err = this.SaveToFile("file", tempFilePath)
	checkError(err)
	return tempFileName, fileName, tempFilePath
}
