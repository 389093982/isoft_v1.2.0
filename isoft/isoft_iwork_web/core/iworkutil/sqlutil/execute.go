package sqlutil

import (
	"database/sql"
	"isoft/isoft_iwork_web/core/iworkpool"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
)

func ExecuteWithTx(sqlstring string, sql_binding []interface{}, tx *sql.Tx) (lastInsertId, affected int64) {
	stmt, err := tx.Prepare(sqlstring)
	errorutil.CheckError(err)
	result, err := stmt.Exec(sql_binding...)
	errorutil.CheckError(err)
	_affected, err := result.RowsAffected()
	errorutil.CheckError(err)
	affected = _affected
	lastInsertId, _ = result.LastInsertId()
	return
}

func Execute(sqlstring string, sql_binding []interface{}, dataSourceName string) (lastInsertId, affected int64) {
	db, err := iworkpool.GetDBConn("mysql", dataSourceName)
	errorutil.CheckError(err)
	tx, _ := db.Begin()
	defer tx.Commit()
	return ExecuteWithTx(sqlstring, sql_binding, tx)
}

func ExecuteSql(sqlstring string, sql_binding []interface{}, dataSourceName string) (lastInsertId, affected int64, err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = errorutil.ToError(err1)
		}
	}()
	lastInsertId, affected = Execute(sqlstring, sql_binding, dataSourceName)
	return
}
