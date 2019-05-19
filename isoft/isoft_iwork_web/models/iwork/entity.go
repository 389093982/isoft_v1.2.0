package iwork

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Entity struct {
	Id              int64     `json:"id"`
	EntityName      string    `json:"entity_name"`
	EntityType      string    `json:"entity_type"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func QueryEntity(condArr map[string]string, page int, offset int) (entities []Entity, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("entity")
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		qs = qs.Filter("entity_name__contains", search)
	}
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&entities)
	return
}

func InsertOrUpdateEntity(entity *Entity) (id int64, err error) {
	o := orm.NewOrm()
	if entity.Id > 0 {
		id, err = o.Update(entity)
	} else {
		id, err = o.Insert(entity)
	}
	return
}

func QueryEntityByName(entity_name string) (entity Entity, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("entity").Filter("entity_name", entity_name).One(&entity)
	return
}

func DeleteEntityById(entity_id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("entity").Filter("id", entity_id).Delete()
	return err
}

func QueryAllEntityInfo() (entities []Entity) {
	o := orm.NewOrm()
	o.QueryTable("entity").OrderBy("id").All(&entities)
	return
}
