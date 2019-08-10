package controllers

import (
	"github.com/astaxie/beego/context"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/models"
	"strings"
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
	trackingId, receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: mapData})
	if receiver != nil {
		receiver.TmpDataMap[iworkconst.TRACKING_ID] = trackingId
		this.Data["json"] = &receiver.TmpDataMap
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", iworkconst.TRACKING_ID: trackingId, "errorMsg": "Empty Response"}
	}
	this.ServeJSON()
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
				if paramValue := ctx.Input.Query(item.ParamName); strings.TrimSpace(paramValue) != "" {
					mapData[item.ParamName] = paramValue
				}
			}
			break
		}
	}
	return mapData
}
