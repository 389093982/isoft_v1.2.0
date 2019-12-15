package models

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"strings"
	"time"
)

type Placement struct {
	Id              int64     `json:"id"`
	PlacementName   string    `json:"placement_name"`
	PlacementLabel  string    `json:"placement_label"`
	PlacementDesc   string    `json:"placement_desc" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time" orm:"auto_now_add;type(datetime)"`
	ElementLimit    int64     `json:"element_limit"`
}

type Element struct {
	Id                 int64     `json:"id"`
	Placement          string    `json:"placement"`
	Title              string    `json:"title"`
	Content            string    `json:"content" orm:"type(text)"`
	LinkedRefer        string    `json:"linked_refer"`
	ImgPath            string    `json:"img_path"`
	Status             int64     `json:"status"`
	NavigationLevel    int64     `json:"navigation_level"`
	NavigationParentId int64     `json:"navigation_parent_id"`
	MdContent          string    `json:"md_content" orm:"type(text)"`
	CreatedBy          string    `json:"created_by"`
	CreatedTime        time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy      string    `json:"last_updated_by"`
	LastUpdatedTime    time.Time `json:"last_updated_time" orm:"auto_now_add;type(datetime)"`
}

func QueryPagePlacement(condArr map[string]string, page int, offset int, o orm.Ormer) (placements []Placement, counts int64, err error) {
	qs := o.QueryTable("placement")
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		qs = qs.Filter("placement_name__contains", search) // 或者描述搜索
	}
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&placements)
	return
}

func InsertOrUpdatePlacement(placement *Placement) (id int64, err error) {
	o := orm.NewOrm()
	if placement.Id > 0 {
		id, err = o.Update(placement)
	} else {
		id, err = o.Insert(placement)
	}
	return
}

func QueryPlacementByName(placement_name string) (placement Placement, err error) {
	err = orm.NewOrm().QueryTable("placement").Filter("placement_name", placement_name).One(&placement)
	return
}

func QueryPlacementById(id int64) (placement Placement, err error) {
	err = orm.NewOrm().QueryTable("placement").Filter("id", id).One(&placement)
	return
}

func DeletePlacementById(id int64) (err error) {
	_, err = orm.NewOrm().QueryTable("placement").Filter("id", id).Delete()
	return
}

func CopyPlacement(id int64) (err error) {
	var placement Placement
	err1 := orm.NewOrm().QueryTable("placement").Filter("id", id).One(&placement)
	placement.Id = 0 // 重置 id 为 0
	_, err2 := InsertOrUpdatePlacement(&placement)
	return errorutil.GetFirstError2(err1, err2)
}

func GetAllPlacements() (placements []Placement, err error) {
	_, err = orm.NewOrm().QueryTable("placement").All(&placements)
	return
}

func FilterPageElement(condArr map[string]string, page int, offset int, o orm.Ormer) (elements []Element, counts int64, err error) {
	qs := o.QueryTable("element")
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		qs = qs.Filter("placement__contains", search) // ...
	}
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&elements)
	return
}

func InsertOrUpdateElement(element *Element) (id int64, err error) {
	o := orm.NewOrm()
	if element.Id > 0 {
		id, err = o.Update(element)
	} else {
		id, err = o.Insert(element)
	}
	return
}

func QueryElementById(id int64) (element Element, err error) {
	err = orm.NewOrm().QueryTable("element").Filter("id", id).One(&element)
	return
}

func QueryLimitValidElementByPlaceName(placement_name string, limit int64) (elements []Element, err error) {
	_, err = orm.NewOrm().QueryTable("element").Filter("placement", placement_name).
		Filter("status", 1).OrderBy("-created_time").Limit(limit).All(&elements)
	return
}

func CopyElement(id int64) (err error) {
	var element Element
	err1 := orm.NewOrm().QueryTable("element").Filter("id", id).One(&element)
	element.Id = 0 // 重置 id 为 0
	_, err2 := InsertOrUpdateElement(&element)
	return errorutil.GetFirstError2(err1, err2)
}
