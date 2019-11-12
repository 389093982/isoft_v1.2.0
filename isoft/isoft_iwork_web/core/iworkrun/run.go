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
		logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_ERROR, fmt.Sprintf("<span style='color:red;'>internal error:%s</span>", err.Error()))
	}
	defer logwriter.RecordCostTimeLog("execute work", trackingId, time.Now())

	// 记录前置 filterTrackingIds 信息
	if filterTrackingIds := getFilterTrackingIds(dispatcher); filterTrackingIds != "" {
		logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("filter stack:%s", filterTrackingIds))
	}
	// 记录日志详细
	logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", workCache.Work.WorkName))

	// 初始化数据中心
	initDataStore := datastore.InitDataStore(trackingId, logwriter, workCache)

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
	logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", workCache.Work.WorkName))
	return
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
