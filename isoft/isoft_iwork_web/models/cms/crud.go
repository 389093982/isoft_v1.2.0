package cms

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

func FilterConfigurations(condArr map[string]string, page int, offset int) (configurations []Configuration, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("configuration")
	var cond = orm.NewCondition()
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("configuration_name__contains", search).Or("configuration_value__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&configurations)
	return
}

func QueryAllConfigurations(configuration_name string, parent_id int64) (configurations []*Configuration, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("configuration").Filter("configuration_name", configuration_name).Filter("parent_id", parent_id).All(&configurations)
	if err == nil && len(configurations) > 0 {
		for _, configuration := range configurations {
			sub, err := QueryAllConfigurations(configuration_name, configuration.Id)
			if err == nil && len(sub) > 0 {
				configuration.SubConfigurations = sub
			}
		}
	}
	return
}

func AddConfiguration(configuration *Configuration) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(configuration)
	return
}

func QueryRandomCommonLink(link_type string) (commonLinks []*CommonLink, err error) {
	o := orm.NewOrm()
	_, err = o.Raw(fmt.Sprintf("SELECT link_name,link_addr FROM COMMON_LINK WHERE link_type = '%s' ORDER BY RAND() limit 50", link_type)).QueryRows(&commonLinks)
	return
}

func FilterCommonLinks(condArr map[string]string, page int, offset int) (commonLinks []CommonLink, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("common_link")
	var cond = orm.NewCondition()
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("link_name__contains", search).Or("link_addr__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	qs = qs.OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&commonLinks)
	return
}

func AddCommonLink(commonLink *CommonLink) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(commonLink)
	return
}
