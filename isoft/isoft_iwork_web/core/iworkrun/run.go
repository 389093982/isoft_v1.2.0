package iworkrun

import (
	"fmt"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"net/http"
	"time"
)

// dispatcher 为父流程或者调用者传递下来的参数
func RunOneWork(work_id int64, dispatcher *entry.Dispatcher) (trackingId string, receiver *entry.Receiver) {
	// 缓冲日志写入对象
	logwriter := createNewLoggerWriter(dispatcher)
	defer logwriter.Close()
	workCache, err := iworkcache.GetWorkCache(work_id)
	// 为当前流程创建新的 trackingId, 前提条件 cacheContext.Work 一定存在
	trackingId = createNewTrackingIdForWork(dispatcher, workCache.Work)
	if err != nil {
		logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_ERROR, errorutil.FormatInternalError(err))
	}
	defer logwriter.RecordCostTimeLog("execute work", trackingId, time.Now())

	// 记录继承下来的日志: 如前置 filterTrackingIds 信息
	recordExtendLog(dispatcher, logwriter, trackingId)

	// 记录日志: 标记流程执行开始和结束
	recordStartEndLog(trackingId, logwriter, workCache, "start")
	defer recordStartEndLog(trackingId, logwriter, workCache, "end")

	// 初始化数据中心
	initDataStore := datastore.InitDataStore(trackingId, logwriter, workCache, dispatcher, &node.TxManager{})
	if dispatcher == nil || dispatcher.TxManger == nil { // dispatcher 不含事务管理器,则代表是最外层流程
		defer func() {
			tsm := initDataStore.TxManger.(*node.TxManager)
			tsm.Commit(true) // 提交事务,暂不支持事务回退,只支持事务提交
		}()
	}
	bsoRunner := node.BlockStepOrdersRunner{
		ParentStepId: iworkconst.PARENT_STEP_ID_FOR_START_END,
		WorkCache:    workCache,
		TrackingId:   trackingId,
		LogWriter:    logwriter,
		Store:        initDataStore, // 获取数据中心
		Dispatcher:   dispatcher,    // dispatcher 是全流程共享的
		RunOneStep:   RunOneStep,
	}
	receiver = bsoRunner.Run()
	return
}

// 记录前置 filterTrackingIds 信息
func recordExtendLog(dispatcher *entry.Dispatcher, logwriter *iworklog.CacheLoggerWriter, trackingId string) {
	if filterTrackingIds := getFilterTrackingIds(dispatcher); filterTrackingIds != "" {
		msg := fmt.Sprintf("filter stack:%s", filterTrackingIds)
		logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, msg)
	}
}

func recordStartEndLog(trackingId string, logwriter *iworklog.CacheLoggerWriter, workCache *iworkcache.WorkCache, pattern string) {
	msg := fmt.Sprintf("~~~~~~~~~~%s execute work:%s~~~~~~~~~~", pattern, workCache.Work.WorkName)
	logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, msg)

}

func createNewLoggerWriter(dispatcher *entry.Dispatcher) *iworklog.CacheLoggerWriter {
	var logwriter *iworklog.CacheLoggerWriter
	// 调度者不为空时代表有父级流程
	if dispatcher != nil && dispatcher.TmpDataMap != nil && dispatcher.TmpDataMap["logwriter"] != nil {
		logwriter = dispatcher.TmpDataMap["logwriter"].(*iworklog.CacheLoggerWriter)
	} else {
		logwriter = new(iworklog.CacheLoggerWriter)
	}
	return logwriter
}

// 执行单个 BlockStep
func RunOneStep(args *interfaces.RunOneStepArgs) (receiver *entry.Receiver) {
	// 统计耗费时间
	defer args.Logwriter.RecordCostTimeLog(args.BlockStep.Step.WorkStepName, args.TrackingId, time.Now())
	// 记录开始执行日志
	startLogStr := fmt.Sprintf("start execute blockStep: >>>>>>>>>> [[<span style='color:blue;'>%s<span>]]", args.BlockStep.Step.WorkStepName)
	args.Logwriter.Write(args.TrackingId, "", iworkconst.LOG_LEVEL_INFO, startLogStr)
	// 记录结束执行日志
	endLogStr := fmt.Sprintf("end execute blockStep: >>>>>>>>>> [[%s]]", args.BlockStep.Step.WorkStepName)
	defer args.Logwriter.Write(args.TrackingId, "", iworkconst.LOG_LEVEL_INFO, endLogStr)

	// 由工厂代为执行步骤
	factory := &node.WorkStepFactory{
		WorkStep:         args.BlockStep.Step,
		Dispatcher:       args.Dispatcher,
		Receiver:         receiver,
		BlockStep:        args.BlockStep,
		DataStore:        args.Datastore,
		LogWriter:        args.Logwriter,
		BlockStepRunFunc: RunOneStep,
		WorkSubRunFunc:   RunOneWork,
		WorkCache:        args.WorkCache,
	}
	factory.Execute(args.TrackingId)
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	return factory.Receiver
}

func getFilterTrackingIds(dispatcher *entry.Dispatcher) string {
	if dispatcher != nil {
		if request, ok := dispatcher.TmpDataMap[iworkconst.HTTP_REQUEST_OBJECT].(*http.Request); ok && request != nil {
			if filterTrackingIds := request.Header.Get(iworkconst.FILTER_TRACKING_ID_STACK); filterTrackingIds != "" {
				request.Header.Del(iworkconst.FILTER_TRACKING_ID_STACK)
				return filterTrackingIds
			}
		}
	}
	return ""
}
