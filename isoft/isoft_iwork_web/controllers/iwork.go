package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/service"
	"time"
)

type WorkController struct {
	beego.Controller
}

func (this *WorkController) BuildIWorkDL() {
	//dls := make([]*IWorkDL,0)
	//works := iwork.GetAllWorkInfo()
	//for _, work := range works{
	//	dl := &IWorkDL{}
	//	steps, _ := iwork.GetAllWorkStepInfo(work.Id)
	//	for _, step := range steps{
	//		if step.WorkStepType == "work_start"{
	//			dl.RequestInfo = step.WorkStepInput
	//		}
	//		if step.WorkStepType == "work_end"{
	//			dl.ResponseInfo = step.WorkStepOutput
	//		}
	//	}
	//	dls = append(dls, dl)
	//}
}

func (this *WorkController) GetRelativeWork() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.GetRelativeWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":      "SUCCESS",
			"parentWorks": result["parentWorks"],
			"subworks":    result["subworks"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetLastRunLogDetail() {
	tracking_id := this.GetString("tracking_id")
	runLogRecord, _ := models.QueryRunLogRecordWithTracking(tracking_id)
	runLogDetails, err := models.QueryLastRunLogDetail(tracking_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "runLogDetails": runLogDetails, "runLogRecord": runLogRecord}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageLogRecord() {
	work_id, _ := this.GetInt64("work_id")
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	serviceArgs := map[string]interface{}{"work_id": work_id, "offset": offset, "current_page": current_page, "ctx": this.Ctx}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.FilterPageLogRecord); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":        "SUCCESS",
			"runLogRecords": result["runLogRecords"],
			"paginator":     result["paginator"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) RunWork() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if err := service.ExecuteServiceWithTx(serviceArgs, service.RunWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWork() {
	// 将请求参数封装成 work
	var work models.Work
	work_id, err := this.GetInt64("work_id", -1)
	if err == nil && work_id > 0 {
		work.Id = work_id
	}
	work.WorkName = this.GetString("work_name")
	work.WorkDesc = this.GetString("work_desc")
	work.CreatedBy = "SYSTEM"
	work.CreatedTime = time.Now()
	work.LastUpdatedBy = "SYSTEM"
	work.LastUpdatedTime = time.Now()
	serviceArgs := map[string]interface{}{"work": work}
	if err := service.ExecuteServiceWithTx(serviceArgs, service.EditWorkService); err == nil {
		work, _ := models.QueryWorkByName(work.WorkName, orm.NewOrm())
		go flushCache(work.Id)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageWork() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	serviceArgs := map[string]interface{}{"condArr": condArr, "offset": offset, "current_page": current_page, "ctx": this.Ctx}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.FilterPageWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":    "SUCCESS",
			"works":     result["works"],
			"paginator": result["paginator"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkById() {
	id, _ := this.GetInt64("id")
	serviceArgs := map[string]interface{}{"id": id}
	if err := service.ExecuteServiceWithTx(serviceArgs, service.DeleteWorkByIdService); err == nil {
		go flushCache(id)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func flushCache(work_id ...int64) (err error) {
	works := make([]models.Work, 0)
	if len(work_id) > 0 {
		work, err := models.QueryWorkById(work_id[0], orm.NewOrm())
		if err != nil && errors.As(err, &orm.ErrNoRows) {
			iworkcache.DeleteWorkCache(work_id[0]) // 不存在则删除
			return nil
		} else {
			serviceArgs := map[string]interface{}{"work_id": work_id[0]}
			if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.GetRelativeWorkService); err == nil {
				works = append(works, result["parentWorks"].([]models.Work)...)
				works = append(works, result["subworks"].([]models.Work)...)
			}
			works = append(works, work)
		}
	} else {
		works = models.QueryAllWorkInfo(orm.NewOrm())
	}
	for _, work := range works {
		parser := schema.WorkStepFactoryParamSchemaParser{}
		if err = iworkcache.UpdateWorkCache(work.Id, &parser); err != nil {
			break
		}
		go saveHistory(work.Id)
	}
	return
}
