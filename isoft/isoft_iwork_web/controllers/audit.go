package controllers

import (
	"encoding/json"
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/jsonutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft/common/xmlutil"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

func (this *WorkController) EditAuditTask() {
	id, err := this.GetInt64("id", -1)
	taskName := this.GetString("task_name")
	taskDesc := this.GetString("task_desc")
	task := &models.AuditTask{
		TaskName:        taskName,
		TaskDesc:        taskDesc,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	if err == nil && id > 0 {
		task.Id = id
	}
	_, err = models.InsertOrUpdateAuditTask(task, orm.NewOrm())
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) QueryPageAuditTask() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	tasks, count, err := models.QueryPageAuditTask(current_page, offset, orm.NewOrm())
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "tasks": tasks,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditAuditTaskTarget() {
	taskName := this.GetString("task_name")
	update_cases := this.GetString("update_cases")
	task, _ := models.QueryAuditTaskByTaskName(taskName, orm.NewOrm())
	var taskDetail models.TaskDetail
	xml.Unmarshal([]byte(task.TaskDetail), &taskDetail)
	json.Unmarshal([]byte(update_cases), &taskDetail.UpdateCases)
	task.TaskDetail = xmlutil.RenderToString(taskDetail)
	// 配置入库
	_, err := models.InsertOrUpdateAuditTask(&task, orm.NewOrm())
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditAuditTaskSource() {
	taskName := this.GetString("task_name")
	resourceName := this.GetString("resource_name")
	querySql := this.GetString("query_sql")
	resource, _ := models.QueryResourceByName(resourceName)
	colNames := sqlutil.GetMetaDatas(querySql, resource.ResourceDsn)
	task, _ := models.QueryAuditTaskByTaskName(taskName, orm.NewOrm())
	task.TaskDetail = xmlutil.RenderToString(&models.TaskDetail{
		ResourceName: resourceName,
		QuerySql:     querySql,
		ColNames:     jsonutil.RenderToNoIndentJson(colNames),
	})
	// 配置入库
	_, err := models.InsertOrUpdateAuditTask(&task, orm.NewOrm())
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) QueryTaskDetail() {
	taskName := this.GetString("task_name")
	task, err := models.QueryAuditTaskByTaskName(taskName, orm.NewOrm())
	if err == nil {
		var taskDetail models.TaskDetail
		xml.Unmarshal([]byte(task.TaskDetail), &taskDetail)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "task": task, "taskDetail": taskDetail}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
