package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"isoft/isoft/common/fileutils"
	"isoft/isoft/common/xmlutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var persistentDirPath string

func init() {
	// 获取 persistent 目录
	_, file, _, _ := runtime.Caller(0)
	persistentDirPath = fileutils.ChangeToLinuxSeparator(fmt.Sprintf("%s/persistent", filepath.Dir(filepath.Dir(file))))
}

func persistentToFile() {
	// 进行备份操作
	fileutils.CopyDir(persistentDirPath, fmt.Sprintf(`%s_backup/%s`, persistentDirPath, time.Now().Format("20060102150405")))
	// 进行清理操作
	fileutils.RemoveFileOrDirectory(persistentDirPath)
	persistentFiltersToFile()
	persistentModulesToFile()
	persistentGlobalVarsToFile()
	persistentQuartzsToFile()
	persistentResourcesToFile()
	persistentMigratesToFile()
	persistentWorkCahcesToFile()
	persistentAuditTasksToFile()
}

func persistentAuditTasksToFile() {
	tasks, _ := models.QueryAllAuditTasks(orm.NewOrm())
	for _, task := range tasks {
		filepath := path.Join(persistentDirPath, "audits", fmt.Sprintf(`%s.audit`, task.TaskName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(task)), false)
	}
}

func persistentModulesToFile() {
	modules, _ := models.QueryAllModules()
	for _, module := range modules {
		filepath := path.Join(persistentDirPath, "modules", fmt.Sprintf(`%s.module`, module.ModuleName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(module)), false)
	}
}

func persistentFiltersToFile() {
	filters, _ := models.QueryAllFilters()
	for _, filter := range filters {
		filepath := path.Join(persistentDirPath, "filters", fmt.Sprintf(`%s.filter`, filter.FilterWorkName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(filter)), false)
	}
}

func persistentGlobalVarsToFile() {
	globalVars := models.QueryAllGlobalVar()
	for _, globalVar := range globalVars {
		filepath := path.Join(persistentDirPath, "globalVars", fmt.Sprintf(`%s.globalVar`, globalVar.Name))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(globalVar)), false)
	}
}

func persistentQuartzsToFile() {
	metas, _ := models.QueryAllCronMeta()
	for _, meta := range metas {
		filepath := path.Join(persistentDirPath, "quartzs", fmt.Sprintf(`%s.quartz`, meta.TaskName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(meta)), false)
	}
}

func persistentResourcesToFile() {
	resources := models.QueryAllResource()
	for _, resource := range resources {
		filepath := path.Join(persistentDirPath, "resources", fmt.Sprintf(`%s.resource`, resource.ResourceName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(resource)), false)
	}
}

func persistentMigratesToFile() {
	migrates, _ := models.QueryAllSqlMigrate()
	for _, migrate := range migrates {
		filepath := path.Join(persistentDirPath, "migrates", fmt.Sprintf(`%s.migrate`, migrate.MigrateName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(migrate)), false)
	}
}

func persistentWorkCahcesToFile() {
	workCahces := iworkcache.GetAllWorkCache()
	for _, workCahce := range workCahces {
		filepath := path.Join(persistentDirPath, "works", fmt.Sprintf(`%s.work`, workCahce.Work.WorkName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(workCahce)), false)
	}
}

//**********************************************************************************************************************
func backupTable(tableName, backupName string) {
	p1, err := orm.NewOrm().Raw(fmt.Sprintf(`CREATE TABLE %s LIKE %s`, backupName, tableName)).Prepare()
	errorutil.CheckError(err)
	_, err = p1.Exec()
	errorutil.CheckError(err)
	p2, err := orm.NewOrm().Raw(fmt.Sprintf(`INSERT INTO %s SELECT * FROM %s`, backupName, tableName)).Prepare()
	errorutil.CheckError(err)
	_, err = p2.Exec()
	errorutil.CheckError(err)
}

func backupDB() {
	sql_migrate_backup := fmt.Sprintf(`backup_sql_migrate_%s`, time.Now().Format("20060102150405"))
	backupTable("sql_migrate", sql_migrate_backup)
	work_backup := fmt.Sprintf(`backup_work_%s`, time.Now().Format("20060102150405"))
	backupTable("work", work_backup)
	workstep_backup := fmt.Sprintf(`backup_work_step_%s`, time.Now().Format("20060102150405"))
	backupTable("work_step", workstep_backup)
	filter_backup := fmt.Sprintf(`backup_filters_%s`, time.Now().Format("20060102150405"))
	backupTable("filters", filter_backup)
	cron_meta_backup := fmt.Sprintf(`backup_cron_meta_%s`, time.Now().Format("20060102150405"))
	backupTable("cron_meta", cron_meta_backup)
	resource_backup := fmt.Sprintf(`backup_resource_%s`, time.Now().Format("20060102150405"))
	backupTable("resource", resource_backup)
	module_backup := fmt.Sprintf(`backup_module_%s`, time.Now().Format("20060102150405"))
	backupTable("module", module_backup)
	globalVar_backup := fmt.Sprintf(`backup_globalVar_%s`, time.Now().Format("20060102150405"))
	backupTable("global_var", globalVar_backup)
	audit_task_backup := fmt.Sprintf(`backup_audit_task_%s`, time.Now().Format("20060102150405"))
	backupTable("audit_task", audit_task_backup)
}

func truncateDB() {
	orm.NewOrm().QueryTable("filters").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("cron_meta").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("resource").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("module").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("global_var").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("sql_migrate").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("work").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("work_step").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("audit_task").Filter("id__gt", 0).Delete()
}

func persistentToDB(dirPath string, persistentFunc func(v interface{}, filepath string), v interface{}) {
	filepaths, _, _ := fileutils.GetAllSubFile(dirPath)
	for _, filepath := range filepaths {
		persistentFunc(v, filepath)
	}
}

func persistentModelToDB(v interface{}, filepath string) {
	bytes, _ := ioutil.ReadFile(filepath)
	xml.Unmarshal(bytes, v)
	_, err := orm.NewOrm().Insert(v)
	errorutil.CheckError(err)
}

func persistentWorksFileToDB(v interface{}, filepath string) {
	workCache := iworkcache.WorkCache{}
	bytes, _ := ioutil.ReadFile(filepath)
	err := xml.Unmarshal(bytes, &workCache)
	errorutil.CheckError(err)
	work := workCache.Work
	work.CreatedTime = time.Now()
	work.LastUpdatedTime = time.Now()
	_, err = orm.NewOrm().Insert(&work)
	errorutil.CheckError(err)
	for _, step := range workCache.Steps {
		step.CreatedTime = time.Now()
		step.LastUpdatedTime = time.Now()
		_, err := orm.NewOrm().Insert(&step)
		errorutil.CheckError(err)
	}
}

func importProject() {
	if persistent_initial, _ := beego.AppConfig.Bool("persistent.initial"); persistent_initial == true {
		// backupDB()
		truncateDB()
		persistentToDB(fmt.Sprintf("%s/filters", persistentDirPath), persistentModelToDB, &models.Filters{})
		persistentToDB(fmt.Sprintf("%s/quartzs", persistentDirPath), persistentModelToDB, &models.CronMeta{})
		persistentToDB(fmt.Sprintf("%s/resources", persistentDirPath), persistentModelToDB, &models.Resource{})
		persistentToDB(fmt.Sprintf("%s/modules", persistentDirPath), persistentModelToDB, &models.Module{})
		persistentToDB(fmt.Sprintf("%s/globalVars", persistentDirPath), persistentModelToDB, &models.GlobalVar{})
		persistentToDB(fmt.Sprintf("%s/works", persistentDirPath), persistentWorksFileToDB, nil)
		persistentToDB(fmt.Sprintf("%s/migrates", persistentDirPath), persistentModelToDB, &models.SqlMigrate{})
		persistentToDB(fmt.Sprintf("%s/audits", persistentDirPath), persistentModelToDB, &models.AuditTask{})
	}
}
