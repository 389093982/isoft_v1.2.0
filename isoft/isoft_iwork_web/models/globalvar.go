package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type GlobalVar struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name" orm:"unique"`
	Value           string    `json:"value" orm:"type(text)"`
	Type            int       `json:"type"` // 类型：0 表示不可删除
	Desc            string    `json:"desc"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func QueryGlobalVarByName(name string) (gv GlobalVar, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("global_var")
	err = qs.Filter("name", name).One(&gv)
	return
}

func DeleteGlobalVarById(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("global_var").Filter("id", id).Delete()
	return err
}

func InsertOrUpdateGlobalVar(globalVar *GlobalVar, o orm.Ormer) (id int64, err error) {
	if globalVar.Id > 0 {
		id, err = o.Update(globalVar)
	} else {
		id, err = o.Insert(globalVar)
	}
	return
}

func QueryAllGlobalVar() (globalVars []GlobalVar) {
	o := orm.NewOrm()
	o.QueryTable("global_var").All(&globalVars)
	return
}

func QueryGlobalVar(condArr map[string]string, page int, offset int, o orm.Ormer) (globalVars []GlobalVar, counts int64, err error) {
	qs := o.QueryTable("global_var")
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		qs = qs.Filter("name__contains", search)
	}
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&globalVars)
	return
}
