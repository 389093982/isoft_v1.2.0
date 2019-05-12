package iworknode

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkcache"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkplugin/iworkprotocol"
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

type BlockStepOrdersRunner struct {
	ParentStepId int64
	CacheContext *iworkcache.CacheContext
	TrackingId   string
	Logwriter    *iworklog.CacheLoggerWriter
	Store        *datastore.DataStore
	Dispatcher   *entry.Dispatcher
	RunOneStep   iworkprotocol.RunOneStep
}

func (this *BlockStepOrdersRunner) Run() (receiver *entry.Receiver) {
	// 存储前置步骤 afterJudgeInterrupt 属性
	afterJudgeInterrupt := false
	for _, blockStep := range this.CacheContext.BlockStepOrdersMap[this.ParentStepId] {
		if blockStep.Step.WorkStepType == "empty" {
			continue
		}

		args := &iworkprotocol.RunOneStepArgs{
			TrackingId:   this.TrackingId,
			Logwriter:    this.Logwriter,
			BlockStep:    blockStep,
			Datastore:    this.Store,
			Dispatcher:   this.Dispatcher,
			CacheContext: this.CacheContext,
		}

		if blockStep.Step.WorkStepType == "if" { // 遇到 if 必定可以执行
			receiver = this.RunOneStep(args)
			afterJudgeInterrupt = blockStep.AfterJudgeInterrupt
		} else if stringutil.CheckContains(blockStep.Step.WorkStepType, []string{"elif", "else"}) { // 遇到 elif 和 else
			if !afterJudgeInterrupt {
				receiver = this.RunOneStep(args)
				afterJudgeInterrupt = blockStep.AfterJudgeInterrupt
			}
		} else { // 非 if、elif、else 节点必定执行
			receiver = this.RunOneStep(args)
			afterJudgeInterrupt = false
		}
	}
	return receiver
}
