package node

import (
	"database/sql"
	"isoft/isoft_iwork_web/core/iworkpool"
	"sync"
)

// 事务管理器类
// 创建 DataStore 时会初始化 TxManager,TxManager 创建方式有两种: 新创建、从 Dispatcher 中获取
type TxManager struct {
	Tx       *sql.Tx
	hasBegin bool // 是否开启了事务
	lock     sync.Mutex
}

func (this *TxManager) FirstBegin(dataSourceName string) {
	db, _ := iworkpool.GetDBConn("mysql", dataSourceName)
	this.lock.Lock() // 并发情况下控制只开启一次事务
	defer this.lock.Unlock()
	if !this.hasBegin {
		this.Tx, _ = db.Begin()
		this.hasBegin = true
	}
}

func (this *TxManager) Commit(flag bool) {
	if this.Tx != nil {
		if flag {
			this.Tx.Commit()
		} else {
			this.Tx.Rollback()
		}
	}
}
