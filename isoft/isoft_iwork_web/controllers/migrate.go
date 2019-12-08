package controllers

import (
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/pkg/errors"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/core/iworkutil/migrateutil"
	"isoft/isoft_iwork_web/models"
	"regexp"
	"strings"
	"time"
)

var (
	MIGRATE_NAME_FORMAT      = "^(CREATE|UPDATE|DELETE|INSERT|ALTER|DROP)_[a-zA-Z0-9_]+\\.sql$"
	DATE_MIGRATE_NAME_FORMAT = "^[0-9]{14}_(CREATE|UPDATE|DELETE|INSERT|ALTER|DROP)_[a-zA-Z0-9_]+\\.sql$"
)

func (this *WorkController) GetLastMigrateLogs() {
	defer this.ServeJSON()
	trackingId := this.GetString("trackingId")
	logs, _ := models.QueryAllSqlMigrateLog(trackingId)
	if len(logs) > 0 && strings.Contains(logs[len(logs)-1].TrackingDetail, "__OVER__") {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "logs": logs, "over": true}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "logs": logs}
	}
}

func (this *WorkController) ExecuteMigrate() {
	resource_name := this.GetString("resource_name")
	forceClean, _ := this.GetBool("forceClean", false)
	resource, _ := models.QueryResourceByName(resource_name)
	trackingId := stringutil.RandomUUID()
	go migrateutil.MigrateToDB(trackingId, resource.ResourceDsn, forceClean)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "trackingId": trackingId}
	this.ServeJSON()
}

// .*匹配除 \n 以外的任何字符
func checkMigrateSqlFormat(sql string) string {
	sql, _ = stringutil.ReplaceAllString(sql, "\\/\\*.*\\*\\/;", "") // 去除注释 /**/;
	sql, _ = stringutil.ReplaceAllString(sql, "\\/\\*.*\\*\\/", "")  // 去除注释 /**/
	sql = stringutil.TrimEmptyLines(sql)                             // 去除所有空行
	return sql
}

func (this *WorkController) EditSqlMigrate() {
	defer this.ServeJSON()
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": errorutil.ToError(err).Error()}
		}
	}()
	id, _ := this.GetInt64("id")
	migrate := new(models.SqlMigrate)
	migrate_name := this.GetString("migrate_name")
	migrate_name = stringutil.AppendSuffix(migrate_name, ".sql")
	if match, _ := regexp.MatchString(MIGRATE_NAME_FORMAT, migrate_name); match {
		migrate_name = fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), migrate_name)
	} else if match, _ := regexp.MatchString(DATE_MIGRATE_NAME_FORMAT, migrate_name); !match {
		panic(errors.New(fmt.Sprintf("migrate_name must match with %s", MIGRATE_NAME_FORMAT)))
	}
	migrate_sql := this.GetString("migrate_sql")
	migrate_sql = checkMigrateSqlFormat(migrate_sql) // 检查 sql 格式
	migrate_hash := hashutil.CalculateHashWithString(migrate_sql)
	if id > 0 {
		*migrate, _ = models.QuerySqlMigrateById(id)
	} else {
		migrate.CreatedBy = "SYSTEM"
		migrate.CreatedTime = time.Now()
		migrate.Effective = true
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
}

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

func (this *WorkController) GetSqlMigrateInfo() {
	defer this.ServeJSON()
	id, _ := this.GetInt64("id")
	migrate, err := models.QuerySqlMigrateInfo(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "migrate": migrate}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
}

func (this *WorkController) ToggleSqlMigrateEffective() {
	defer this.ServeJSON()
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": errorutil.ToError(err).Error()}
		}
	}()
	id, _ := this.GetInt64("id")
	migrate, err := models.QuerySqlMigrateInfo(id)
	checkError(err)
	migrate.Effective = !migrate.Effective
	_, err = models.InsertOrUpdateSqlMigrate(&migrate)
	checkError(err)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
}
