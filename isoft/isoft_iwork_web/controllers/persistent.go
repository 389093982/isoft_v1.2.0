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
	"reflect"
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
	persistentPlacementsToFile()
}

func persistentPlacementsToFile() {
	placements, _ := models.GetAllPlacements()
	for _, placement := range placements {
		elements, _ := models.QueryElementsByPlacename(placement.PlacementName)
		filepath := path.Join(persistentDirPath, "placements", fmt.Sprintf(`%s.placement`, placement.PlacementName))
		data := &models.PlacementElementMppaer{
			Placement: placement,
			Elements:  elements,
		}
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(data)), false)
	}
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
		filepath := path.Join(persistentDirPath, "migrates", fmt.Sprintf(`%s`, migrate.MigrateName))
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

// 长度过长可能会批量导入失败,需要进一步拆分
func persistentWorkFilesToDB(dirPath string) {
	filepaths, _, _ := fileutils.GetAllSubFile(dirPath)
	var err error
	works := make([]models.Work, 0)
	workSteps := make([]models.WorkStep, 0)
	for _, filepath := range filepaths {
		works, workSteps = parseWorkFile(filepath, works, workSteps)
	}
	_, err = orm.NewOrm().InsertMulti(len(works), works)
	errorutil.CheckError(err)
	_, err = orm.NewOrm().InsertMulti(len(workSteps), workSteps)
	errorutil.CheckError(err)
}

func parseWorkFile(filepath string, works []models.Work, workSteps []models.WorkStep) ([]models.Work, []models.WorkStep) {
	workCache := iworkcache.WorkCache{}
	bytes, _ := ioutil.ReadFile(filepath)
	err := xml.Unmarshal(bytes, &workCache)
	errorutil.CheckError(err)
	work := workCache.Work
	work.CreatedTime = time.Now()
	work.LastUpdatedTime = time.Now()
	works = append(works, work)
	errorutil.CheckError(err)
	for _, step := range workCache.Steps {
		step.CreatedTime = time.Now()
		step.LastUpdatedTime = time.Now()
		workSteps = append(workSteps, step)
	}
	return works, workSteps
}

func Insert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	appendSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 1, 1)
	appendSlice.Index(0).Set(reflect.ValueOf(value))
	v = reflect.AppendSlice(v.Slice(0, pos), reflect.AppendSlice(appendSlice, v.Slice(pos, v.Len())))
	return v.Interface()
}

func Append(slice interface{}, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	return Insert(slice, v.Len(), value)
}

// 批量插入 DB
func persistentMultiToDB(dirPath string, tp reflect.Type) {
	filepaths, _, _ := fileutils.GetAllSubFile(dirPath)
	// reflect.PtrTo 返回类型t的指针的类型
	// reflect.SliceOf 返回类型t的切片的类型
	rows := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(tp)), 0, 0).Interface()
	for _, filepath := range filepaths {
		bytes, _ := ioutil.ReadFile(filepath)
		// reflect.New 返回一个Value类型值,该值持有一个指向类型为typ的新申请的零值的指针,返回值的Type为PtrTo(typ)
		val := reflect.New(tp).Interface()
		xml.Unmarshal(bytes, val)
		rows = Append(rows, val)
	}
	_, err := orm.NewOrm().InsertMulti(len(filepaths), rows)
	errorutil.CheckError(err)
}

func importProject() {
	if persistent_initial, _ := beego.AppConfig.Bool("persistent.initial"); persistent_initial == true {
		backupDB()
		truncateDB()
		persistentMultiToDB(fmt.Sprintf("%s/filters", persistentDirPath), reflect.TypeOf(models.Filters{}))
		persistentMultiToDB(fmt.Sprintf("%s/quartzs", persistentDirPath), reflect.TypeOf(models.CronMeta{}))
		persistentMultiToDB(fmt.Sprintf("%s/resources", persistentDirPath), reflect.TypeOf(models.Resource{}))
		persistentMultiToDB(fmt.Sprintf("%s/modules", persistentDirPath), reflect.TypeOf(models.Module{}))
		persistentMultiToDB(fmt.Sprintf("%s/globalVars", persistentDirPath), reflect.TypeOf(models.GlobalVar{}))
		persistentMultiToDB(fmt.Sprintf("%s/migrates", persistentDirPath), reflect.TypeOf(models.SqlMigrate{}))
		persistentMultiToDB(fmt.Sprintf("%s/audits", persistentDirPath), reflect.TypeOf(models.AuditTask{}))
		persistentWorkFilesToDB(fmt.Sprintf("%s/works", persistentDirPath))
	}
}
