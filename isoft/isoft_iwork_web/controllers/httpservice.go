package controllers

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
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
	var (
		work  models.Work
		steps []models.WorkStep
		err   error
	)
	work, err = models.QueryWorkByName(work_name, orm.NewOrm())
	checkError(err)
	steps, err = models.QueryAllWorkStepByWorkName(work_name, orm.NewOrm())
	checkError(err)
	mapData := this.ParseParam(steps)
	receiver := iworkrun.RunOneWork(work.Id, &entry.Dispatcher{TmpDataMap: mapData})
	if receiver != nil {
		this.Data["json"] = &receiver.TmpDataMap
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "Empty Response"}
	}
	this.ServeJSON()
}

func (this *WorkController) ParseParam(steps []models.WorkStep) map[string]interface{} {
	mapData := map[string]interface{}{}
	for _, step := range steps {
		if step.WorkStepType == "work_start" {
			parser := schema.WorkStepFactoryParamSchemaParser{WorkStep: &step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step}}
			inputSchema := parser.GetCacheParamInputSchema()
			for _, item := range inputSchema.ParamInputSchemaItems {
				// 默认参数类型都当成 string 类型
				if paramValue := this.Input().Get(item.ParamName); strings.TrimSpace(paramValue) != "" {
					mapData[item.ParamName] = paramValue
				}
			}
			break
		}
	}
	return mapData
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
