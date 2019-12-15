package node

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"strings"
)

type BlockStepOrdersRunner struct {
	ParentStepId int64
	WorkCache    *iworkcache.WorkCache
	TrackingId   string
	LogWriter    *iworklog.CacheLoggerWriter
	Store        *datastore.DataStore
	Dispatcher   *entry.Dispatcher
	RunOneStep   interfaces.RunOneStep
}

func (this *BlockStepOrdersRunner) recordLog(err interface{}) {
	var (
		_err         interface{}
		workStepName string
		errorMsg     string
	)
	if wsError, ok := err.(interfaces.WorkStepError); ok {
		_err = wsError.Err
		workStepName = wsError.WorkStepName
	} else {
		_err = err
	}

	errorMsg = strings.Join([]string{
		"~~~~~~~~~~~~~~~~~~~~~~~~ internal error trace stack ~~~~~~~~~~~~~~~~~~~~~~~~~~",
		errorutil.PanicTraceForHtml(4), // 记录 4 kb大小的堆栈信息
		errorutil.FormatInternalError(_err),
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~",
	},
		"<br/>")
	this.LogWriter.Write(this.TrackingId, workStepName, iworkconst.LOG_LEVEL_ERROR, errorMsg)
}

func (this *BlockStepOrdersRunner) Run() (receiver *entry.Receiver) {
	defer func() {
		if err := recover(); err != nil {
			this.recordLog(err)
			// 重置 parentStepId,并执行 end 节点
			this.ParentStepId = iworkconst.PARENT_STEP_ID_FOR_START_END
			receiver = this.runDetail(true)
		}
	}()
	return this.runDetail()
}

func (this *BlockStepOrdersRunner) runDetail(runEnd ...bool) (receiver *entry.Receiver) {
	if len(runEnd) > 0 { // end 节点异常暂不抛出
		defer func() {
			if err := recover(); err != nil {
				this.recordLog(err)
			}
		}()
	}
	// 存储前置步骤 afterJudgeInterrupt 属性
	afterJudgeInterrupt := false
	for _, blockStep := range this.WorkCache.BlockStepOrdersMap[this.ParentStepId] {
		if this.skippable(blockStep, runEnd...) {
			continue
		}
		args := &interfaces.RunOneStepArgs{
			TrackingId: this.TrackingId,
			Logwriter:  this.LogWriter,
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

// 当前 blockStep 是否可以跳过
func (this *BlockStepOrdersRunner) skippable(blockStep *block.BlockStep, runEnd ...bool) bool {
	if len(runEnd) > 0 && runEnd[0] == true && blockStep.Step.WorkStepType != "work_end" { // 不满足 runEnd 条件
		return true
	}
	if blockStep.Step.WorkStepType == "empty" {
		return true
	}
	return false
}
