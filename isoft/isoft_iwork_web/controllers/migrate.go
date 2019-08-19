package controllers

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkutil/migrateutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

// @router /api/iwork/executeMigrate [post]
func (this *WorkController) ExecuteMigrate() {
	resource_name := this.GetString("resource_name")
	forceClean, _ := this.GetBool("forceClean", false)
	resource, _ := models.QueryResourceByName(resource_name)
	if err := migrateutil.MigrateToDB(resource.ResourceDsn, forceClean); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

// @router /api/iwork/editSqlMigrate [post]
func (this *WorkController) EditSqlMigrate() {
	id, _ := this.GetInt64("id")
	migrate := new(models.SqlMigrate)
	migrate_name := this.GetString("migrate_name")
	migrate_sql := this.GetString("migrate_sql")
	migrate_hash := hashutil.CalculateHashWithString(migrate_sql)
	if id > 0 {
		*migrate, _ = models.QuerySqlMigrateById(id)
	} else {
		migrate.CreatedBy = "SYSTEM"
		migrate.CreatedTime = time.Now()
	}
	migrate.LastUpdatedBy = "SYSTEM"
	migrate.LastUpdatedTime = time.Now()
	migrate.MigrateName = migrate_name
	migrate.MigrateSql = migrate_sql
	migrate.MigrateHash = migrate_hash
	if _, err := models.InsertOrUpdateSqlMigrate(migrate); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

// @router /api/iwork/filterPageSqlMigrate [post]
func (this *WorkController) FilterPageSqlMigrate() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	migrates, count, err := models.QuerySqlMigrate(current_page, offset)
	if err == nil {
		resources := models.QueryAllResource("db")
		paginator := pagination.SetPaginator(this.Ctx, offset, count)
		this.Data["json"] = &map[string]interface{}{
			"status":    "SUCCESS",
			"migrates":  migrates,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums()),
			"resources": resources,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

// @router /api/iwork/getSqlMigrateInfo [post]
func (this *WorkController) GetSqlMigrateInfo() {
	id, _ := this.GetInt64("id")
	migrate, err := models.QuerySqlMigrateInfo(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "migrate": migrate}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
