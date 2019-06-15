package controllers

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/models/iwork"
	"strings"
	"time"
)

// 示例地址: http://localhost:8086/api/iwork/httpservice/test_iblog_table_migrate?author0=admin1234567
func (this *WorkController) PublishSerivce() {
	start := time.Now()
	defer func() {
		if err := recover(); err != nil {
			costTime := time.Now().Sub(start).Nanoseconds() / 1e6
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.(error).Error(), "cost(ms)": costTime}
			this.ServeJSON()
		}
	}()
	work_name := this.Ctx.Input.Param(":work_name")
	work, err := iwork.QueryWorkByName(work_name, orm.NewOrm())
	if err != nil {
		panic(err)
	}
	steps, err := iwork.QueryAllWorkStepByWorkName(work_name, orm.NewOrm())
	if err != nil {
		panic(err)
	}
	mapData := this.ParseParam(steps)
	receiver := iworkrun.RunOneWork(work.Id, &entry.Dispatcher{TmpDataMap: mapData})
	costTime := time.Now().Sub(start).Nanoseconds() / 1e6
	if receiver != nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "result": receiver.TmpDataMap, "cost_ms": costTime}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "cost(ms)": costTime}
	}
	this.ServeJSON()
}

func (this *WorkController) ParseParam(steps []iwork.WorkStep) map[string]interface{} {
	mapData := map[string]interface{}{}
	for _, step := range steps {
		if step.WorkStepType == "work_start" {
			parser := schema.WorkStepFactorySchemaParser{WorkStep: &step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step}}
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
