package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"strings"
	"testing"
)

func Test_sql(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/isoft_iaas")

	rows, _ := db.Query("select * from share where id = ? and link_href = ? and last_updated_time = ?", nil, nil, nil)
	colNames, _ := rows.Columns()
	fmt.Println(colNames)
}

func Test_sql2(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/isoft_iaas")

	stmt, _ := db.Prepare("select * from share where id = ? and link_href = ? and last_updated_time = ?")
	rows, _ := stmt.Query(nil, nil, nil)
	colNames, _ := rows.Columns()
	fmt.Println(colNames)
}

func Test_sql3(t *testing.T) {
	fmt.Println(strings.Count("a?b?c?d", "?"))
}

func Test_sql4(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/isoft_iwork_backup0929")
	errorutil.CheckError(err)
	err = db.Ping()
	errorutil.CheckError(err)
}
