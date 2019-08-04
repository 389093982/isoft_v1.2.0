package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Filters struct {
	Id              int64     `json:"id"`
	FilterWorkId    int64     `json:"filter_work_id"`
	WorkName        string    `json:"work_name"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertMultiFilters(filter_id int64, filters []*Filters) (num int64, err error) {
	o := orm.NewOrm()
	o.QueryTable("filters").Filter("filter_work_id", filter_id).Delete()
	num, err = o.InsertMulti(len(filters), &filters)
	return
}

func QueryAllFilters() (filters []Filters, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("filters").All(&filters)
	return
}
