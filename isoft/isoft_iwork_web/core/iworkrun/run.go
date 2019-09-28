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
	"time"
)

// dispatcher 为父流程遗传下来的参数
func RunOneWork(work_id int64, dispatcher *entry.Dispatcher) (trackingId string, receiver *entry.Receiver) {
	logwriter := createNewLoggerWriter(dispatcher)
	defer logwriter.Close()
	workCache, err := iworkcache.GetWorkCache(work_id)
	// 为当前流程创建新的 trackingId, 前提条件 cacheContext.Work 一定存在
	trackingId = createNewTrackingIdForWork(dispatcher, workCache.Work)
	if err != nil {
		logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_ERROR, fmt.Sprintf("<span style='color:red;'>internal error:%s</span>", err.Error()))
	}
	defer logwriter.RecordCostTimeLog("execute work", trackingId, time.Now())
	// 记录日志详细
	logwriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", workCache.Work.WorkName))

	// 初始化数据中心
	initDataStore := datastore.InitDataStore(trackingId, logwriter, workCache)

	bsoRunner := node.BlockStepOrdersRunner{
		ParentStepId: -1,
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
	log := "start execute blockStep: >>>>>>>>>> [[<span style='color:blue;'>%s<span>]]"
	args.Logwriter.Write(args.TrackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf(log, args.BlockStep.Step.WorkStepName))
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
	// 记录结束执行日志
	args.Logwriter.Write(args.TrackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("end execute blockStep: >>>>>>>>>> [[%s]]", args.BlockStep.Step.WorkStepName))
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	return factory.Receiver
}
