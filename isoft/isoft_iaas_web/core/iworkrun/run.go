package iworkrun

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/memory"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkplugin/iworknode"
	"isoft/isoft_iaas_web/core/iworkutil/errorutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

// dispatcher 为父流程遗传下来的参数
func RunOneWork(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	logwriter := new(iworklog.CacheLoggerWriter)
	defer logwriter.Close()

	// 为当前流程创建新的 trackingId
	trackingId := createNewTrackingIdForWork(dispatcher, work)
	defer logwriter.RecordCostTimeLog("execute work", trackingId, time.Now())
	defer func() {
		if err := recover(); err != nil {
			// 记录 4 kb大小的堆栈信息
			logwriter.Write(trackingId, "~~~~~~~~~~~~~~~~~~~~~~~~ internal error trace stack ~~~~~~~~~~~~~~~~~~~~~~~~~~")
			logwriter.Write(trackingId, string(errorutil.PanicTrace(4)))
			logwriter.Write(trackingId, fmt.Sprintf("<span style='color:red;'>internal error:%s</span>", err))
			logwriter.Write(trackingId, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
	}()
	// 记录日志详细
	logwriter.Write(trackingId, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", work.WorkName))
	// 获取数据中心
	store := datastore.InitDataStore(trackingId, logwriter)

	// 将 steps 转换成 BlockSteps, 逐个 block 依次执行
	blockStepOrders := iworknode.GetBlockStepExecuteOrder(block.ParseToBlockStep(steps))
	receiver = iworknode.BlockStepOrdersRunnerWarpper(blockStepOrders, trackingId, logwriter, store, dispatcher, RunOneStep)

	// 注销 MemoryCache,无需注册,不存在时会自动注册
	memory.UnRegistMemoryCache(trackingId)
	logwriter.Write(trackingId, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", work.WorkName))
	return
}

// 执行单个 BlockStep
func RunOneStep(args *iworknode.RunOneStepArgs) (receiver *entry.Receiver) {
	// 统计耗费时间
	defer args.Logwriter.RecordCostTimeLog(args.BlockStep.Step.WorkStepName, args.TrackingId, time.Now())
	// 记录开始执行日志
	log := "start execute blockStep: >>>>>>>>>> [[<span style='color:blue;'>%s<span>]]"
	args.Logwriter.Write(args.TrackingId, fmt.Sprintf(log, args.BlockStep.Step.WorkStepName))
	// 由工厂代为执行步骤
	factory := &iworknode.WorkStepFactory{
		WorkStep:         args.BlockStep.Step,
		Dispatcher:       args.Dispatcher,
		Receiver:         receiver,
		BlockStep:        args.BlockStep,
		DataStore:        args.Datastore,
		LogWriter:        args.Logwriter,
		BlockStepRunFunc: RunOneStep,
		WorkSubRunFunc:   RunOneWork,
	}
	factory.Execute(args.TrackingId)
	// 记录结束执行日志
	args.Logwriter.Write(args.TrackingId, fmt.Sprintf("end execute blockStep: >>>>>>>>>> [[%s]]", args.BlockStep.Step.WorkStepName))
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	return factory.Receiver
}
