package iworkrun

import (
	"fmt"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/interfaces"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"time"
)

func GetCacheParamInputSchemaFunc(workStep *models.WorkStep) *iworkmodels.ParamInputSchema {
	parser := schema.WorkStepFactorySchemaParser{WorkStep: workStep, ParamSchemaParser: &node.WorkStepFactory{WorkStep: workStep}}
	return parser.GetCacheParamInputSchema()
}

// dispatcher 为父流程遗传下来的参数
func RunOneWork(work_id int64, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	logwriter := new(iworklog.CacheLoggerWriter)
	defer logwriter.Close()
	workCache, err := iworkcache.GetWorkCache(work_id, GetCacheParamInputSchemaFunc)
	// 为当前流程创建新的 trackingId, 前提条件 cacheContext.Work 一定存在
	trackingId := createNewTrackingIdForWork(dispatcher, workCache.Work)
	if err != nil {
		logwriter.Write(trackingId, fmt.Sprintf("<span style='color:red;'>internal error:%s</span>", err.Error()))
	}
	defer logwriter.RecordCostTimeLog("execute work", trackingId, time.Now())
	// 记录日志详细
	logwriter.Write(trackingId, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", workCache.Work.WorkName))

	bsoRunner := node.BlockStepOrdersRunner{
		ParentStepId: -1,
		WorkCache:    workCache,
		TrackingId:   trackingId,
		LogWriter:    logwriter,
		Store:        datastore.InitDataStore(trackingId, logwriter), // 获取数据中心
		Dispatcher:   dispatcher,
		RunOneStep:   RunOneStep,
	}
	receiver = bsoRunner.Run()
	logwriter.Write(trackingId, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", workCache.Work.WorkName))
	return
}

// 执行单个 BlockStep
func RunOneStep(args *interfaces.RunOneStepArgs) (receiver *entry.Receiver) {
	// 统计耗费时间
	defer args.Logwriter.RecordCostTimeLog(args.BlockStep.Step.WorkStepName, args.TrackingId, time.Now())
	// 记录开始执行日志
	log := "start execute blockStep: >>>>>>>>>> [[<span style='color:blue;'>%s<span>]]"
	args.Logwriter.Write(args.TrackingId, fmt.Sprintf(log, args.BlockStep.Step.WorkStepName))
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
	args.Logwriter.Write(args.TrackingId, fmt.Sprintf("end execute blockStep: >>>>>>>>>> [[%s]]", args.BlockStep.Step.WorkStepName))
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	return factory.Receiver
}
