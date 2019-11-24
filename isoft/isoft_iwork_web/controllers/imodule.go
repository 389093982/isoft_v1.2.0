package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

func (this *WorkController) ModuleList() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	modules, count, err := models.QueryPageModuleList(condArr, current_page, offset, orm.NewOrm())
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "modules": modules,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditModule() {
	var module models.Module
	module_id, err := this.GetInt64("module_id", -1)
	if err == nil && module_id > 0 {
		module.Id = module_id
	}
	module.ModuleName = this.GetString("module_name")
	module.ModuleDesc = this.GetString("module_desc")
	module.CreatedBy = "SYSTEM"
	module.CreatedTime = time.Now()
	module.LastUpdatedBy = "SYSTEM"
	module.LastUpdatedTime = time.Now()
	_, err = models.InsertOrUpdateModule(&module, orm.NewOrm())
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteModuleById() {
	id, _ := this.GetInt64("id")
	_, err := models.DeleteModuleById(id, orm.NewOrm())
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetAllModules() {
	moudles, err := models.QueryAllModules()
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "moudles": moudles}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
