package iworkpool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/models"
	"sync"
)

var dbMap map[string]*sql.DB
var m *sync.RWMutex

func init() {
	m = new(sync.RWMutex)
	dbMap = make(map[string]*sql.DB, 0)

}

func LoadAndCachePool() {
	// 初始化所有数据库连接
	resources := models.QueryAllResource("db")
	for _, resource := range resources {
		openDBConn("mysql", resource.ResourceDsn)
	}
}

func GetDBConn(driverName, dataSourceName string) (*sql.DB, error) {
	m.RLock()
	defer m.RUnlock()
	if db, ok := dbMap[driverName+"_"+dataSourceName]; ok {
		return db, nil
	} else {
		return nil, errors.New("miss sql.DB for " + dataSourceName)
	}
}

func openDBConn(driverName, dataSourceName string) (err error) {
	m.Lock()
	defer m.Unlock()
	if db, err := openConn(driverName, dataSourceName); err == nil {
		dbMap[driverName+"_"+dataSourceName] = db
	}
	return
}

func openConn(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err == nil {
		err = db.Ping()
	}
	return db, err
}

func GetConnForMysql(driverName, dataSourceName string) (db *sql.DB, err error) {
	db, err = sql.Open(driverName, dataSourceName)
	if err == nil {
		err = db.Ping()
	}
	return
}
