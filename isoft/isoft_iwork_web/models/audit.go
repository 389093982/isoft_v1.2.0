package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AuditTask struct {
	Id              int64     `json:"id"`
	TaskName        string    `json:"task_name"`
	TaskDesc        string    `json:"task_desc" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateAuditTask(task *AuditTask, o orm.Ormer) (id int64, err error) {
	if task.Id > 0 {
		id, err = o.Update(task)
	} else {
		id, err = o.Insert(task)
	}
	return
}

func QueryPageAuditTask(page int, offset int, o orm.Ormer) (tasks []AuditTask, counts int64, err error) {
	qs := o.QueryTable("audit_task")
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&tasks)
	return
}
