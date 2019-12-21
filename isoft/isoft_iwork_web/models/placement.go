package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"strings"
	"time"
)

type PlacementElementMppaer struct {
	Placement Placement `json:"placement"`
	Elements  []Element `json:"elements"`
}

type Placement struct {
	Id              int64     `json:"id"`
	PlacementName   string    `json:"placement_name"`
	PlacementLabel  string    `json:"placement_label"`
	PlacementType   string    `json:"placement_type"`
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
	ElementName        string    `json:"element_name"`
	ElementLabel       string    `json:"element_label"`
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
	var cond = orm.NewCondition()
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("placement_name__contains", search).Or("placement_label__contains", search).Or("placement_desc__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&placements)
	return
}

func InsertOrUpdatePlacement(placement *Placement) (id int64, err error) {
	o := orm.NewOrm()
	if placement.Id > 0 {
		old, _ := QueryPlacementById(placement.Id)
		if old.PlacementName != placement.PlacementName {
			// 更新 element
			_, err = o.QueryTable("element").Filter("placement", old.PlacementName).Update(orm.Params{"placement": placement.PlacementName})
		}
		id, err = o.Update(placement)
	} else {
		_, err := QueryPlacementByName(placement.PlacementName)
		if err == nil {
			return -1, errors.New("占位符已存在,请修改占位符名称！")
		}
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
	placement, _ := QueryPlacementById(id)
	count, _ := orm.NewOrm().QueryTable("element").Filter("placement", placement.PlacementName).Count()
	if count > 0 {
		return errors.New("请先删除 placement 关联的 element 元素，然后再进行删除操作！")
	}
	_, err = orm.NewOrm().QueryTable("placement").Filter("id", id).Delete()
	return
}

func CopyPlacement(id int64) (err error) {
	var placement Placement
	err1 := orm.NewOrm().QueryTable("placement").Filter("id", id).One(&placement)
	placement.Id = 0 // 重置 id 为 0
	placement.PlacementName = placement.PlacementName + "_copy"
	_, err2 := InsertOrUpdatePlacement(&placement)
	return errorutil.GetFirstError2(err1, err2)
}

func GetAllPlacements() (placements []Placement, err error) {
	_, err = orm.NewOrm().QueryTable("placement").All(&placements)
	return
}

func FilterPageElement(condArr map[string]string, placement_name string, page int, offset int, o orm.Ormer) (elements []Element, counts int64, err error) {
	qs := o.QueryTable("element")
	var cond = orm.NewCondition()
	cond = cond.And("placement", placement_name)
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("element_label__contains", search).Or("element_name__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.OrderBy("placement", "navigation_level", "navigation_parent_id", "-last_updated_time").Limit(offset, (page-1)*offset)
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

func QueryElementByPlacementAndElementName(placement, element_name string) (element Element, err error) {
	err = orm.NewOrm().QueryTable("element").Filter("placement", placement).Filter("element_name", element_name).One(&element)
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
	element.Id = 0      // 重置 id 为 0
	element.Status = -1 // 复制的默认失效
	_, err2 := InsertOrUpdateElement(&element)
	return errorutil.GetFirstError2(err1, err2)
}

func DeleteElementById(id int64) (err error) {
	_, err = orm.NewOrm().QueryTable("element").Filter("id", id).Delete()
	return
}

func QueryElementsByPlacename(placement_name string) (elements []Element, err error) {
	_, err = orm.NewOrm().QueryTable("element").Filter("placement", placement_name).All(&elements)
	return
}
