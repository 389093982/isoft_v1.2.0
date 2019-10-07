package iworkpool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
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
	models.RegisterOpenConnFunc(OpenDBConn)
	// 初始化所有数据库连接
	resources := models.QueryAllResource("db")
	for _, resource := range resources {
		OpenDBConn("mysql", resource.ResourceDsn)
	}
}

func GetDBConn(driverName, dataSourceName string) (*sql.DB, error) {
	m.RLock()
	defer m.RUnlock()
	// $RESOURCE 格式
	if db, ok := dbMap[driverName+"_"+dataSourceName]; ok {
		return db, nil
	}
	// 一般格式
	return openConn(driverName, dataSourceName)
}

func OpenDBConn(driverName, dataSourceName string) (err error) {
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
