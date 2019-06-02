package iworknode

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
)

type BlockStepOrdersRunner struct {
	ParentStepId int64
	WorkCache    *iworkcache.WorkCache
	TrackingId   string
	Logwriter    *iworklog.CacheLoggerWriter
	Store        *datastore.DataStore
	Dispatcher   *entry.Dispatcher
	RunOneStep   iworkprotocol.RunOneStep
}

func (this *BlockStepOrdersRunner) Run() (receiver *entry.Receiver) {
	parentStepId := this.ParentStepId // 记录当前的 parentStepId
	defer func() {
		if err := recover(); err != nil {
			// 记录 4 kb大小的堆栈信息
			this.Logwriter.Write(this.TrackingId, "~~~~~~~~~~~~~~~~~~~~~~~~ internal error trace stack ~~~~~~~~~~~~~~~~~~~~~~~~~~")
			this.Logwriter.Write(this.TrackingId, string(errorutil.PanicTrace(4)))
			this.Logwriter.Write(this.TrackingId, fmt.Sprintf("<span style='color:red;'>internal error:%s</span>", err))
			this.Logwriter.Write(this.TrackingId, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			// 重置 parentStepId,并执行 end 节点
			this.ParentStepId = parentStepId
			receiver = this.runDetail(true)
		}
	}()
	return this.runDetail()
}

func (this *BlockStepOrdersRunner) runDetail(runEnd ...bool) (receiver *entry.Receiver) {
	// 存储前置步骤 afterJudgeInterrupt 属性
	afterJudgeInterrupt := false
	for _, blockStep := range this.WorkCache.BlockStepOrdersMap[this.ParentStepId] {
		if len(runEnd) > 0 && runEnd[0] == true && blockStep.Step.WorkStepType != "work_end" { // 不满足 runEnd 条件
			continue
		}
		if blockStep.Step.WorkStepType == "empty" {
			continue
		}
		args := &iworkprotocol.RunOneStepArgs{
			TrackingId: this.TrackingId,
			Logwriter:  this.Logwriter,
			BlockStep:  blockStep,
			Datastore:  this.Store,
			Dispatcher: this.Dispatcher,
			WorkCache:  this.WorkCache,
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
