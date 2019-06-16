package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Template struct {
	Id              int64     `json:"id"`
	TemplateTheme   string    `json:"template_theme"`
	TemplateName    string    `json:"template_name"`
	TemplateValue   string    `json:"template_value"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateTemplate(template *Template, o orm.Ormer) (id int64, err error) {
	if template.Id > 0 {
		id, err = o.Update(template)
	} else {
		id, err = o.Insert(template)
	}
	return
}

func QueryTemplate(condArr map[string]string, page int, offset int, o orm.Ormer) (templates []Template, counts int64, err error) {
	qs := o.QueryTable("template")
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		qs = qs.Filter("template_value__contains", search)
	}
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&templates)
	return
}

func DeleteTemplateById(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("template").Filter("id", id).Delete()
	return err
}
