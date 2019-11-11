package startup

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"isoft/isoft/common/fileutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/models"
	"path/filepath"
	"runtime"
	"time"
)

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
}

func truncateDB() {
	orm.NewOrm().QueryTable("sql_migrate").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("work").Filter("id__gt", 0).Delete()
	orm.NewOrm().QueryTable("work_step").Filter("id__gt", 0).Delete()
}

func persistentToDB(dirPath string, persistentFunc func(filepath string)) {
	filepaths, _, _ := fileutil.GetAllSubFile(dirPath)
	for _, filepath := range filepaths {
		persistentFunc(filepath)
	}
}

func persistentWorkFileToDB(filepath string) {
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

func persistentSqlMigrateFileToDB(filepath string) {
	sqlMigrate := models.SqlMigrate{}
	bytes, _ := ioutil.ReadFile(filepath)
	xml.Unmarshal(bytes, &sqlMigrate)
	_, err := orm.NewOrm().Insert(&sqlMigrate)
	errorutil.CheckError(err)
}

func Persistent() {
	if persistent_initial, _ := beego.AppConfig.Bool("persistent.initial"); persistent_initial == true {
		backupDB()
		truncateDB()

		// 获取 persistent 目录
		_, file, _, _ := runtime.Caller(0)
		persistentPath := fmt.Sprintf("%s/persistent", filepath.Dir(filepath.Dir(file)))

		persistentToDB(fmt.Sprintf("%s/works", persistentPath), persistentWorkFileToDB)
		persistentToDB(fmt.Sprintf("%s/migrates", persistentPath), persistentSqlMigrateFileToDB)
	}
}
