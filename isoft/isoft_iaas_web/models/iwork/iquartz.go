package iwork

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type CronMeta struct {
	Id              int64     `json:"id"`
	TaskName        string    `json:"task_name"` // 任务名称
	TaskType        string    `json:"task_type"` // 任务类型
	CronStr         string    `json:"cron_str"`
	Enable          bool      `json:"enable"` // 是否启用
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func DeleteCronMetaByTaskName(taskName string, o orm.Ormer) error {
	_, err := o.QueryTable("cron_meta").Filter("task_name", taskName).Delete()
	return err
}

func InsertOrUpdateCronMeta(meta *CronMeta, o orm.Ormer) (id int64, err error) {
	if meta.Id > 0 {
		id, err = o.Update(meta)
	} else {
		id, err = o.Insert(meta)
	}
	return
}

func QueryCronMetaByName(taskName string) (meta CronMeta, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("cron_meta")
	err = qs.Filter("task_name", taskName).One(&meta)
	return
}

func QueryCronMeta(condArr map[string]string, page int, offset int) (metas []CronMeta, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("cron_meta")
	var cond = orm.NewCondition()
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("task_name__contains", search).Or("task_type__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&metas)
	return
}

func QueryAllCronMeta() (metas []CronMeta, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("cron_meta").All(&metas)
	return
}
