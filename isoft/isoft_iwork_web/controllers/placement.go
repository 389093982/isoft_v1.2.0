package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

func (this *WorkController) FilterPagePlacement() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr := map[string]string{"search": this.GetString("search")}
	placements, count, err := models.QueryPagePlacement(condArr, current_page, offset, orm.NewOrm())
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "placements": placements,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) EditPlacement() {
	id, _ := this.GetInt64("id", 0)
	placement_name := this.GetString("placement_name")
	placement_desc := this.GetString("placement_desc")
	placement_label := this.GetString("placement_label")
	element_limit, _ := this.GetInt64("element_limit")
	placement_type := this.GetString("placement_type")

	placement := &models.Placement{
		Id:              id,
		PlacementName:   placement_name,
		PlacementDesc:   placement_desc,
		PlacementLabel:  placement_label,
		PlacementType:   placement_type,
		ElementLimit:    element_limit,
		CreatedBy:       `SYSTEM`,
		CreatedTime:     time.Now(),
		LastUpdatedBy:   `SYSTEM`,
		LastUpdatedTime: time.Now(),
	}
	_, err := models.InsertOrUpdatePlacement(placement)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) QueryPlacementByName() {
	placement_name := this.GetString("placement_name")
	placement, err := models.QueryPlacementByName(placement_name)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "placement": placement}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) QueryPlacementById() {
	id, _ := this.GetInt64("id", 0)
	placement, err := models.QueryPlacementById(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "placement": placement}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeletePlacementById() {
	id, _ := this.GetInt64("id", 0)
	err := models.DeletePlacementById(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) CopyPlacement() {
	id, _ := this.GetInt64("id", 0)
	err := models.CopyPlacement(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageElement() {
	placement_name := this.GetString("placement_name")
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr := map[string]string{"search": this.GetString("search")}
	elements, count, err := models.FilterPageElement(condArr, placement_name, current_page, offset, orm.NewOrm())
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "elements": elements,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) EditElement() {
	id, _ := this.GetInt64("id", 0)
	placement := this.GetString("placement")
	element_name := this.GetString("element_name")
	navigation_level, _ := this.GetInt64("navigation_level", 0)
	navigation_parent_id, _ := this.GetInt64("navigation_parent_id", 0)
	element_label := this.GetString("element_label")
	content := this.GetString("content")
	md_content := this.GetString("md_content")
	imgpath := this.GetString("imgpath")
	linked_refer := this.GetString("linked_refer")

	element := &models.Element{
		Id:                 id,
		Placement:          placement,
		ElementName:        element_name,
		NavigationLevel:    navigation_level,
		NavigationParentId: navigation_parent_id,
		ElementLabel:       element_label,
		Content:            content,
		MdContent:          md_content,
		ImgPath:            imgpath,
		LinkedRefer:        linked_refer,
		CreatedBy:          `SYSTEM`,
		CreatedTime:        time.Now(),
		LastUpdatedBy:      `SYSTEM`,
		LastUpdatedTime:    time.Now(),
	}
	_, err := models.InsertOrUpdateElement(element)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) UpdateElementStatus() {
	id, _ := this.GetInt64("id", 0)
	status, _ := this.GetInt64("status", 0)
	var err error
	if status == -2 {
		err = models.DeleteElementById(id)
	} else {
		element, _ := models.QueryElementById(id)
		element.Status = status
		_, err = models.InsertOrUpdateElement(&element)
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) QueryElementById() {
	id, _ := this.GetInt64("id", 0)
	element, err := models.QueryElementById(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "element": element}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) CopyElement() {
	id, _ := this.GetInt64("id", 0)
	err := models.CopyElement(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterElementByPlacement() {
	placement_name := this.GetString("placement")
	placement, _ := models.QueryPlacementByName(placement_name)
	if placement.ElementLimit < 0 {
		placement.ElementLimit = 1000
	}
	elements, err := models.QueryLimitValidElementByPlaceName(placement_name, placement.ElementLimit)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "placement": placement, "elements": elements}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
