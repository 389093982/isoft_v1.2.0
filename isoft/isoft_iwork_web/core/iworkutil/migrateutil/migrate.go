package migrateutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"isoft/isoft/common/hashutil"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
	"strconv"
	"strings"
	"time"
)

type MigrateExecutor struct {
	Dsn        string // dsn 连接串
	db         *sql.DB
	TrackingId string
	ForceClean bool
	migrates   []models.SqlMigrate
}

func (this *MigrateExecutor) ping() (err error) {
	if this.Dsn == "" {
		return errors.New("empty dsn error...")
	}
	// 建立连接
	this.db, err = sql.Open("mysql", this.Dsn)
	return nil
}

func (this *MigrateExecutor) executeForceClean() error {
	if this.ForceClean == false {
		return nil
	}
	if !strings.Contains(this.Dsn, "_test") {
		// 强制清理功能只适用于 _test 库
		return errors.New("ForceClean only can used by *_test database!")
	}

	tableNames := sqlutil.GetAllTableNames(this.Dsn)
	if len(tableNames) > 0 {
		for _, tableName := range tableNames {
			if _, err := this.ExecSQL(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, tableName)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *MigrateExecutor) getAllMigrateVersions() ([]string, []string) {
	sql := `SELECT migrate_name, migrate_hash FROM migrate_version order by created_time asc`
	rows, err := this.db.Query(sql)
	errorutil.CheckError(err)
	defer rows.Close()
	migrateNames := make([]string, 0)
	migrateHashs := make([]string, 0)
	for rows.Next() {
		var migrate_name, migrate_hash string
		err = rows.Scan(&migrate_name, &migrate_hash)
		errorutil.CheckError(err)
		migrateNames = append(migrateNames, migrate_name)
		migrateHashs = append(migrateHashs, migrate_hash)
	}
	return migrateNames, migrateHashs
}

// 检查执行历史
// 1、文件是否被删除
// 2、文件是否被篡改
// 3、文件执行顺序是否被更改
func (this *MigrateExecutor) checkHistory() {
	migrateNames, migrateHashs := this.getAllMigrateVersions()
	for index, migrateName := range migrateNames {
		if migrateName == this.migrates[index].MigrateName {
			if migrateHashs[index] != this.migrates[index].MigrateHash {
				errorMsg := `fatal match error, history migrate_name for %s's hash is %s, but %s was found at migrate_index %d, please check why the file has been modified!`
				panic(errors.New(fmt.Sprintf(errorMsg, migrateName, migrateHashs[index], this.migrates[index].MigrateHash, index)))
			}
		} else {
			panic(errors.New(fmt.Sprintf("fatal match error, history migrate_name is %s, but %s was found at at migrate_index %d",
				migrateName, this.migrates[index].MigrateName, index)))
		}
	}
	return
}

// 建立迁移文件版本管理表
func (this *MigrateExecutor) initial() {
	if err := this.executeForceClean(); err != nil {
		panic(err)
	}
	versionTable := `CREATE TABLE IF NOT EXISTS migrate_version (id INT(20) PRIMARY KEY AUTO_INCREMENT, 
	MIGRATE_NAME CHAR(200), MIGRATE_HASH CHAR(200), CREATED_TIME DATETIME, SUCCESS BOOLEAN);`
	if _, err := this.ExecSQL(versionTable); err != nil {
		panic(err)
	}
	return
}

func (this *MigrateExecutor) ExecSQL(sql string, args ...interface{}) (rs sql.Result, err error) {
	stmt, err := this.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rs, err = stmt.Exec(args...)
	return
}

func (this *MigrateExecutor) QueryRowSQL(sql string, args ...interface{}) (row *sql.Row, err error) {
	if stmt, err := this.db.Prepare(sql); err == nil {
		row = stmt.QueryRow(args...)
	}
	return
}

func (this *MigrateExecutor) insertOrUpdateMigrateVersion(migrate_name, migrate_hash string, successFlag bool) error {
	if this.db != nil {
		var (
			recordLog string
			err       error
		)
		if !successFlag {
			recordLog = `INSERT INTO migrate_version(migrate_name,migrate_hash,created_time, success) VALUES (?,?,NOW(), false);`
			_, err = this.ExecSQL(recordLog, migrate_name, migrate_hash)
		} else {
			recordLog = `UPDATE migrate_version SET success = true where migrate_name = ?;`
			_, err = this.ExecSQL(recordLog, migrate_name)
		}

		return err
	}
	return nil
}

func (this *MigrateExecutor) loadAllMigrate() {
	this.migrates, _ = models.QueryAllSqlMigrate()
}

func (this *MigrateExecutor) migrate() {
	for _, migrate := range this.migrates {
		if err := this.migrateOne(migrate); err != nil {
			return
		}
	}
}

func (this *MigrateExecutor) checkExecuted(migrate_name, migrate_hash string) bool {
	sql := `SELECT COUNT(*) FROM migrate_version WHERE migrate_name = ? and migrate_hash = ?`
	if row, err := this.QueryRowSQL(sql, migrate_name, migrate_hash); err == nil {
		var datacount int64
		if err := row.Scan(&datacount); err == nil && datacount > 0 {
			return true
		}
	}
	return false
}

func (this *MigrateExecutor) InsertSqlMigrateLog(migrate_name, detail string, status bool) {
	log := &models.SqlMigrateLog{
		TrackingId:     this.TrackingId,
		MigrateName:    migrate_name,
		Status:         status,
		TrackingDetail: fmt.Sprintf("%s [%s]", detail, time.Now().Format("2006-01-02 15:04:05")),
	}
	models.InsertSqlMigrateLog(log)
}

func (this *MigrateExecutor) migrateOne(migrate models.SqlMigrate) error {
	hash := hashutil.CalculateHashWithString(migrate.MigrateSql)
	// 已经执行过则忽略
	if this.checkExecuted(migrate.MigrateName, hash) {
		detail := fmt.Sprintf(`%s was migrated and skip it...`, migrate.MigrateName)
		this.InsertSqlMigrateLog(migrate.MigrateName, detail, true)
		return nil
	}
	// 每次迁移都有可能有多个执行 sql
	executeSqls := strings.Split(migrate.MigrateSql, ";")
	executeSqls = datatypeutil.FilterSlice(executeSqls, datatypeutil.CheckNotEmpty)
	// 开启事务前先插入一条 false 记录
	this.insertOrUpdateMigrateVersion(migrate.MigrateName, hash, false)
	tx, err := this.db.Begin()
	errorutil.CheckError(err)
	for _, executeSql := range executeSqls {
		if _, err := this.ExecSQL(executeSql); err != nil {
			tx.Rollback()
			detail := fmt.Sprintf("[%s] - [%s] - [%s] : %s", strconv.FormatInt(migrate.Id, 10), migrate.MigrateName, executeSql, err.Error())
			this.InsertSqlMigrateLog(migrate.MigrateName, detail, false)

			detail = fmt.Sprintf(`%s was migrated failed and rollback ...`, migrate.MigrateName)
			this.InsertSqlMigrateLog(migrate.MigrateName, detail, false)
			return err
		}
	}
	// 控制在同一个事务中执行 migrate_version 表
	err = this.insertOrUpdateMigrateVersion(migrate.MigrateName, hash, true)
	tx.Commit()
	detail := fmt.Sprintf(`%s was migrated success ...`, migrate.MigrateName)
	this.InsertSqlMigrateLog(migrate.MigrateName, detail, true)
	return nil
}

func MigrateToDB(trackingId, dsn string, forceClean bool) (err error) {
	executor := &MigrateExecutor{
		Dsn:        dsn,
		TrackingId: trackingId,
		ForceClean: forceClean,
	}
	defer executor.InsertSqlMigrateLog("", "__OVER__", true)
	defer func() {
		if err := recover(); err != nil {
			detail := errorutil.ToError(err).Error()
			executor.InsertSqlMigrateLog("", detail, false)
		}
	}()

	if err = executor.ping(); err == nil {
		executor.loadAllMigrate()
		executor.initial()
		executor.checkHistory()
		executor.migrate()
	}
	return
}
