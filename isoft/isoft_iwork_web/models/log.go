package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type RunLogRecord struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	WorkId          int64     `json:"work_id"`
	WorkName        string    `json:"work_name"`
	LogLevel        string    `json:"log_level"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type RunLogDetail struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id" orm:"null"`
	WorkStepName    string    `json:"work_step_name" orm:"null"`
	LogLevel        string    `json:"log_level" orm:"null"` // INFO、SUCCESS、ERROR
	Detail          string    `json:"detail" orm:"type(text);null"`
	LogOrder        int64     `json:"log_order" orm:"null"`
	CreatedBy       string    `json:"created_by" orm:"null"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime);null"`
	LastUpdatedBy   string    `json:"last_updated_by" orm:"null"`
	LastUpdatedTime time.Time `json:"last_updated_time" orm:"null"`
}

func InsertRunLogRecord(record *RunLogRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(record)
	return
}

func InsertMultiRunLogDetail(details []*RunLogDetail) (num int64, err error) {
	o := orm.NewOrm()
	for _, detail := range details {
		if detail.LogLevel == "ERROR" {
			o.QueryTable("run_log_record").Filter("tracking_id", detail.TrackingId).Update(orm.Params{"LogLevel": "ERROR"})
			break
		}
	}
	num, err = o.InsertMulti(len(details), &details)
	return
}

func QueryRunLogRecordWithTracking(tracking_id string) (runLogRecord RunLogRecord, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("run_log_record").Filter("tracking_id", tracking_id).One(&runLogRecord)
	return
}

func QueryRunLogRecordCount(workId int64, log_level string) (count int64, err error) {
	qs := orm.NewOrm().QueryTable("run_log_record").Filter("work_id", workId)
	if log_level != "" {
		qs = qs.Filter("log_level", log_level)
	}
	count, err = qs.Count()
	return
}

func QueryRunLogRecord(work_id int64, logLevel string, page int, offset int) (runLogRecords []RunLogRecord, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("run_log_record")
	if work_id > 0 {
		work, _ := QueryWorkById(work_id, o)
		qs = qs.Filter("work_name", work.WorkName)
	}
	if logLevel != "" {
		qs = qs.Filter("log_level", logLevel)
	}
	qs = qs.OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&runLogRecords)
	return
}

func QueryLastRunLogDetail(tracking_id string) (runLogDetails []RunLogDetail, err error) {
	o := orm.NewOrm()
	// __startswith 多级 tracking_id 也查出来
	_, err = o.QueryTable("run_log_detail").Filter("tracking_id__startswith", tracking_id).OrderBy("log_order").All(&runLogDetails)
	return
}

type ValidateLogRecord struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	WorkId          int64     `json:"work_id"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type ValidateLogDetail struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	WorkId          int64     `json:"work_id"`
	WorkStepId      int64     `json:"work_step_id"`
	WorkName        string    `json:"work_name"`
	WorkStepName    string    `json:"work_step_name"`
	ParamName       string    `json:"param_name"`
	Detail          string    `json:"detail" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertValidateLogRecord(record *ValidateLogRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(record)
	return
}

func InsertMultiValidateLogDetail(details []*ValidateLogDetail) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.InsertMulti(len(details), &details)
	return
}

func QueryLastValidateLogRecord(work_id int64) (record ValidateLogRecord, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("validate_log_record")
	if work_id > 0 {
		qs = qs.Filter("work_id", work_id)
	}
	err = qs.OrderBy("-last_updated_time").One(&record)
	return
}

func QueryLastValidateLogDetail(work_id int64) (details []ValidateLogDetail, err error) {
	if record, err := QueryLastValidateLogRecord(work_id); err == nil {
		o := orm.NewOrm()
		_, err = o.QueryTable("validate_log_detail").
			Filter("tracking_id", record.TrackingId).OrderBy("-work_id", "work_step_id").All(&details)
	}
	return
}
