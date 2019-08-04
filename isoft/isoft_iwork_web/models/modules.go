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
