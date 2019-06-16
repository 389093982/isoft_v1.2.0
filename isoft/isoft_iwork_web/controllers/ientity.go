package controllers

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

func (this *WorkController) DeleteEntity() {
	entity_id, _ := this.GetInt64("entity_id", -1)
	if err := models.DeleteEntityById(entity_id); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditEntity() {
	var entity models.Entity
	entity_id, err := this.GetInt64("entity_id", -1)
	if err == nil && entity_id > 0 {
		entity.Id = entity_id
	}
	entity.EntityName = this.GetString("entity_name")
	entity.EntityType = this.GetString("entity_type")
	entity.CreatedBy = "SYSTEM"
	entity.CreatedTime = time.Now()
	entity.LastUpdatedBy = "SYSTEM"
	entity.LastUpdatedTime = time.Now()
	if _, err := models.InsertOrUpdateEntity(&entity); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageEntity() {
	search := this.GetString("search")
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr := map[string]string{"search": search}
	entities, count, err := models.QueryEntity(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "entities": entities,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
