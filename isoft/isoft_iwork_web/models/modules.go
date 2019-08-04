package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Module struct {
	Id              int64     `json:"id"`
	ModuleName      string    `json:"module_name" orm:"unique"`
	ModuleDesc      string    `json:"module_desc" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func QueryAllModules() (modules []Module, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("module").All(&modules)
	return
}

func QueryPageModuleList(condArr map[string]string, page int, offset int, o orm.Ormer) (modules []Module, counts int64, err error) {
	qs := o.QueryTable("module")
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		qs = qs.Filter("module_name__contains", search)
	}
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&modules)
	return
}

func InsertOrUpdateModule(module *Module, o orm.Ormer) (id int64, err error) {
	if module.Id > 0 {
		var oldModule Module
		if err := o.QueryTable("module").Filter("id", module.Id).One(&oldModule); err == nil {
			UpdateModuleName(oldModule.ModuleName, module.ModuleName)
		}
		id, err = o.Update(module)
	} else {
		id, err = o.Insert(module)
	}
	return
}

func DeleteModuleById(module_id int64, o orm.Ormer) (id int64, err error) {
	_, err = o.QueryTable("module").Filter("id", module_id).Delete()
	return
}
