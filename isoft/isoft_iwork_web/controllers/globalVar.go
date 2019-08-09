package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup/memory"
	"time"
)

func (this *WorkController) GlobalVarList() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr := map[string]string{"search": this.GetString("search")}
	globalVars, count, err := models.QueryGlobalVar(condArr, current_page, offset, orm.NewOrm())
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "globalVars": globalVars,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditGlobalVar() {
	id, err := this.GetInt64("id", -1)
	globalVarName := this.GetString("globalVarName")
	globalVarValue := this.GetString("globalVarValue")
	globalVar := &models.GlobalVar{
		Name:            globalVarName,
		Value:           globalVarValue,
		Type:            1,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	if err == nil && id > 0 {
		globalVar.Id = id
	}
	_, err = models.InsertOrUpdateGlobalVar(globalVar, orm.NewOrm())
	if err == nil {
		flushMemoryGlobalVar()
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteGlobalVarById() {
	id, _ := this.GetInt64("id")
	err := models.DeleteGlobalVarById(id)
	if err == nil {
		flushMemoryGlobalVar()
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func flushMemoryGlobalVar() {
	memory.FlushMemoryGlobalVar()
}
