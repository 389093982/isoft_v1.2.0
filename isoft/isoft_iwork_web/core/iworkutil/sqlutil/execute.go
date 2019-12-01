package sqlutil

import (
	"isoft/isoft_iwork_web/core/iworkpool"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
)

func Execute(sqlstring string, sql_binding []interface{}, dataSourceName string) (lastInsertId, affected int64) {
	db, err := iworkpool.GetDBConn("mysql", dataSourceName)
	errorutil.CheckError(err)
	// 使用预编译 sql 防止 sql 注入
	stmt, err := db.Prepare(sqlstring)
	errorutil.CheckError(err)
	result, err := stmt.Exec(sql_binding...)
	errorutil.CheckError(err)
	_affected, err := result.RowsAffected()
	errorutil.CheckError(err)
	affected = _affected
	lastInsertId, _ = result.LastInsertId()
	return
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
