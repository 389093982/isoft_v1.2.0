package framework

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type IFNode struct {
	node.BaseNode
	WorkStep         *models.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *interfaces.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *IFNode) Execute(trackingId string) {
	expression := this.TmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.BOOL_PREFIX + "expression": expression})

	if expression && this.BlockStep.HasChildren {
		this.BlockStep.AfterJudgeInterrupt = true // if 条件满足, AfterJudgeInterrupt 属性变为 true
		bsoRunner := node.BlockStepOrdersRunner{
			ParentStepId: this.WorkStep.WorkStepId,
			WorkCache:    this.WorkCache,
			TrackingId:   trackingId,
			LogWriter:    this.LogWriter,
			Store:        this.DataStore, // 获取数据中心
			Dispatcher:   this.Dispatcher,
			RunOneStep:   this.BlockStepRunFunc,
		}
		bsoRunner.Run()
	} else {
		this.BlockStep.AfterJudgeInterrupt = false
		this.LogWriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *IFNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "expression", "if条件表达式,值为 bool 类型!"},
	}
	return this.BPIS1(paramMap)
}

func (this *IFNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.BOOL_PREFIX + "expression"})
}

type ElIfNode struct {
	node.BaseNode
	WorkStep         *models.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *interfaces.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *ElIfNode) Execute(trackingId string) {
	if this.BlockStep.PreviousBlockStep.Step == nil ||
		!stringutil.CheckContains(this.BlockStep.PreviousBlockStep.Step.WorkStepType, []string{"if", "elif"}) {
		panic(errors.New(fmt.Sprintf(`previous step is not if or elif node for %s`, this.BlockStep.Step.WorkStepName)))
	}
	expression := this.TmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.BOOL_PREFIX + "expression": expression})

	if expression && this.BlockStep.HasChildren {
		this.BlockStep.AfterJudgeInterrupt = true // if 条件满足, AfterJudgeInterrupt 属性变为 true
		bsoRunner := node.BlockStepOrdersRunner{
			ParentStepId: this.WorkStep.WorkStepId,
			WorkCache:    this.WorkCache,
			TrackingId:   trackingId,
			LogWriter:    this.LogWriter,
			Store:        this.DataStore, // 获取数据中心
			Dispatcher:   this.Dispatcher,
			RunOneStep:   this.BlockStepRunFunc,
		}
		bsoRunner.Run()
	} else {
		this.BlockStep.AfterJudgeInterrupt = false
		this.LogWriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *ElIfNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "expression", "if条件表达式,值为 bool 类型!"},
	}
	return this.BPIS1(paramMap)
}

func (this *ElIfNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.BOOL_PREFIX + "expression"})
}

type ElseNode struct {
	node.BaseNode
	WorkStep         *models.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *interfaces.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *ElseNode) Execute(trackingId string) {
	if this.BlockStep.PreviousBlockStep.Step == nil ||
		!stringutil.CheckContains(this.BlockStep.PreviousBlockStep.Step.WorkStepType, []string{"if", "elif"}) {
		panic(errors.New(fmt.Sprintf(`previous step is not if or elif node for %s`, this.BlockStep.Step.WorkStepName)))
	}
	if this.BlockStep.HasChildren {
		bsoRunner := node.BlockStepOrdersRunner{
			ParentStepId: this.WorkStep.WorkStepId,
			WorkCache:    this.WorkCache,
			TrackingId:   trackingId,
			LogWriter:    this.LogWriter,
			Store:        this.DataStore, // 获取数据中心
			Dispatcher:   this.Dispatcher,
			RunOneStep:   this.BlockStepRunFunc,
		}
		bsoRunner.Run()
	}
}
