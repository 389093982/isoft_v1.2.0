package datastore

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/startup/sysconfig"
	"strings"
)

type DataNodeStore struct {
	NodeOutputDataMap map[string]interface{} // 当前节点的输出参数 map
}

// 当前流程数据存储仓库,只在当前流程有效,父子流程之间通信使用 Dispatcher 和 Receiver
type DataStore struct {
	TrackingId       string
	wc               *iworkcache.WorkCache
	logwriter        *iworklog.CacheLoggerWriter
	DataNodeStoreMap map[string]*DataNodeStore
	TxManger         interface{} // 事务管理器
}

// 向数据中心缓存数据
func (this *DataStore) CacheDatas(nodeName string, paramMap map[string]interface{}, byteParamNames ...string) {
	logs := make([]string, 0)
	this.cacheMemory(nodeName, "__output__", paramMap)
	for paramName, paramValue := range paramMap {
		if !sysconfig.SYSCONFIG_VARS_USAGE_LOGGABLE && !this.isReferUsage(nodeName, paramName) {
			continue
		}
		this.cacheMemory(nodeName, paramName, paramValue)
		if !stringutil.CheckContains(paramName, byteParamNames) {
			// 记录日志并存储到 db
			log := fmt.Sprintf("<span style='color:#FF99FF;'> [%s] </span>"+
				"<span style='color:#6633FF;'> cache data for $%s.%s: </span>"+
				"<span style='color:#19be6b;'> %v </span>", this.TrackingId, nodeName, paramName, paramValue)
			logs = append(logs, log)
		}
	}
	this.logwriter.Write(this.TrackingId, nodeName, iworkconst.LOG_LEVEL_SUCCESS, strings.Join(logs, "<br/>"))
}

func (this *DataStore) isReferUsage(nodeName, paramName string) bool {
	if nodeName == "start" || nodeName == "end" {
		return true
	}
	for _, usages := range this.wc.Usage.UsageMap {
		for _, usage := range usages {
			if usage == fmt.Sprintf(`$%s.%s`, nodeName, paramName) {
				return true
			}
		}
	}
	return false
}

// 存储字节数据,不用记录日志
func (this *DataStore) cacheMemory(nodeName, paramName string, paramValue interface{}) {
	// 为当前 nodeName 绑定 DataNodeStore 数据空间
	if _, ok := this.DataNodeStoreMap[nodeName]; !ok {
		this.DataNodeStoreMap[nodeName] = &DataNodeStore{
			NodeOutputDataMap: make(map[string]interface{}, 0),
		}
	}
	// 存数据
	dataNodeStore := this.DataNodeStoreMap[nodeName]
	dataNodeStore.NodeOutputDataMap[paramName] = paramValue
}

// 从数据中心获取数据
func (this *DataStore) GetData(nodeName, paramName string) interface{} {
	store := this.DataNodeStoreMap[nodeName]
	if store == nil {
		return nil
	}
	return this.DataNodeStoreMap[nodeName].NodeOutputDataMap[paramName]
}

// 获取数据中心
func InitDataStore(trackingId string, logwriter *iworklog.CacheLoggerWriter, wc *iworkcache.WorkCache,
	dispatcher *entry.Dispatcher, tx interface{}) *DataStore {
	dataStore := &DataStore{
		TrackingId:       trackingId,
		logwriter:        logwriter,
		wc:               wc,
		DataNodeStoreMap: make(map[string]*DataNodeStore, 0),
		TxManger:         getTxManager(dispatcher, tx),
	}
	initDefaultNodeData(dataStore)
	return dataStore
}

// 获取事务控制器, 获取顺序: 先从 dispatcher 中获取,没有再使用新值
func getTxManager(dispatcher *entry.Dispatcher, tx interface{}) interface{} {
	if dispatcher != nil && dispatcher.TxManger != nil {
		tx = dispatcher.TxManger
	}
	return tx
}

func initDefaultNodeData(dataStore *DataStore) {
	dataStore.DataNodeStoreMap["Error"] = &DataNodeStore{
		// 初始化数据中心中的 isNoError 值,出错时会被覆盖
		NodeOutputDataMap: map[string]interface{}{"isNoError": true},
	}
}
