package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
)

func GetBlockStepExecuteOrder(blockSteps []*block.BlockStep) []*block.BlockStep {
	order := make([]*block.BlockStep, 0)
	deferOrder := make([]*block.BlockStep, 0)
	var end *block.BlockStep
	for _, blockStep := range blockSteps {
		if blockStep.Step.IsDefer == "true" {
			deferOrder = append(deferOrder, blockStep)
		} else if blockStep.Step.WorkStepType == "work_end" {
			end = blockStep
		} else {
			order = append(order, blockStep)
		}
	}
	// is_defer 和 work_end 都是需要延迟执行
	order = append(order, datatypeutil.ReverseSlice(deferOrder).([]*block.BlockStep)...)
	if end != nil {
		order = append(order, end)
	}
	return order
}

type RunOneStepArgs struct {
	TrackingId string
	Logwriter  *iworklog.CacheLoggerWriter
	BlockStep  *block.BlockStep
	Datastore  *datastore.DataStore
	Dispatcher *entry.Dispatcher
}

type RunOneStep func(args *RunOneStepArgs) (receiver *entry.Receiver)

func BlockStepOrdersRunnerWarpper(blockStepOrders []*block.BlockStep, trackingId string, logwriter *iworklog.CacheLoggerWriter,
	store *datastore.DataStore, dispatcher *entry.Dispatcher, runOneStep RunOneStep) (receiver *entry.Receiver) {
	// 存储前置步骤 afterJudgeInterrupt 属性
	afterJudgeInterrupt := false
	for _, blockStep := range blockStepOrders {
		if blockStep.Step.WorkStepType == "empty" {
			continue
		}

		args := &RunOneStepArgs{
			TrackingId: trackingId,
			Logwriter:  logwriter,
			BlockStep:  blockStep,
			Datastore:  store,
			Dispatcher: dispatcher,
		}

		if blockStep.Step.WorkStepType == "elif" || blockStep.Step.WorkStepType == "else" {
			if !afterJudgeInterrupt {
				receiver = runOneStep(args)
				afterJudgeInterrupt = blockStep.AfterJudgeInterrupt
			}
		} else {
			// 当前步骤不是 elif 或者 else
			receiver = runOneStep(args)
			afterJudgeInterrupt = false
		}
	}
	return receiver
}
