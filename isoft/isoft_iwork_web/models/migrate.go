package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SqlMigrate struct {
	Id              int64     `json:"id"`
	MigrateName     string    `json:"migrate_name"`
	MigrateSql      string    `json:"migrate_sql" orm:"type(text)"`
	MigrateHash     string    `json:"migrate_hash"`
	Effective       bool      `json:"effective"` // 是否生效
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type SqlMigrateLog struct {
	Id             int64  `json:"id"`
	TrackingId     string `json:"tracking_id"`
	Status         bool   `json:"status"`
	MigrateName    string `json:"migrate_name"`
	TrackingDetail string `json:"tracking_detail" orm:"type(text)"`
}

func InsertSqlMigrateLog(sml *SqlMigrateLog) (id int64, err error) {
	_, err = orm.NewOrm().Insert(sml)
	return
}

func QueryAllSqlMigrateLog(trackingId string) (logs []SqlMigrateLog, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("sql_migrate_log").Filter("tracking_id", trackingId).OrderBy("id").All(&logs)
	return
}

func QueryAllSqlMigrate() (migrates []SqlMigrate, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("sql_migrate").Filter("effective", true).OrderBy("id").All(&migrates)
	return
}

func QuerySqlMigrateInfo(id int64) (migrate SqlMigrate, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("sql_migrate").Filter("id", id).One(&migrate)
	return
}

func QuerySqlMigrate(current_page, offset int) (migrates []SqlMigrate, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("sql_migrate")
	counts, _ = qs.Count()
	qs = qs.OrderBy("-id").Limit(offset, (current_page-1)*offset)
	_, err = qs.All(&migrates)
	return
}

func InsertOrUpdateSqlMigrate(sm *SqlMigrate) (id int64, err error) {
	o := orm.NewOrm()
	if sm.Id > 0 {
		id, err = o.Update(sm)
	} else {
		id, err = o.Insert(sm)
	}
	return
}

func QuerySqlMigrateById(id int64) (migrate SqlMigrate, err error) {
	err = orm.NewOrm().QueryTable("sql_migrate").Filter("id", id).One(&migrate)
	return
}
