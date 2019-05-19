package iworknode

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iaas_web/models/iwork"
)

type IFNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *iworkprotocol.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *IFNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.BOOL_PREFIX + "expression": expression})

	if expression && this.BlockStep.HasChildren {
		this.BlockStep.AfterJudgeInterrupt = true // if 条件满足, AfterJudgeInterrupt 属性变为 true
		bsoRunner := BlockStepOrdersRunner{
			ParentStepId: this.WorkStep.WorkStepId,
			CacheContext: this.CacheContext,
			TrackingId:   trackingId,
			Logwriter:    this.LogWriter,
			Store:        this.DataStore, // 获取数据中心
			Dispatcher:   nil,
			RunOneStep:   this.BlockStepRunFunc,
		}
		bsoRunner.Run()
	} else {
		this.LogWriter.Write(trackingId, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *IFNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "expression", "if条件表达式,值为 bool 类型!"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *IFNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.BOOL_PREFIX + "expression"})
}

type ElIfNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *iworkprotocol.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *ElIfNode) Execute(trackingId string) {
	if this.BlockStep.PreviousBlockStep.Step == nil ||
		!stringutil.CheckContains(this.BlockStep.PreviousBlockStep.Step.WorkStepType, []string{"if", "elif"}) {
		panic(errors.New(fmt.Sprintf(`previous step is not if or elif node for %s`, this.BlockStep.Step.WorkStepName)))
	}
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.BOOL_PREFIX + "expression": expression})

	if expression && this.BlockStep.HasChildren {
		this.BlockStep.AfterJudgeInterrupt = true // if 条件满足, AfterJudgeInterrupt 属性变为 true
		bsoRunner := BlockStepOrdersRunner{
			ParentStepId: this.WorkStep.WorkStepId,
			CacheContext: this.CacheContext,
			TrackingId:   trackingId,
			Logwriter:    this.LogWriter,
			Store:        this.DataStore, // 获取数据中心
			Dispatcher:   nil,
			RunOneStep:   this.BlockStepRunFunc,
		}
		bsoRunner.Run()
	} else {
		this.LogWriter.Write(trackingId, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *ElIfNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "expression", "if条件表达式,值为 bool 类型!"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *ElIfNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.BOOL_PREFIX + "expression"})
}

type ElseNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *iworkprotocol.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *ElseNode) Execute(trackingId string) {
	if this.BlockStep.PreviousBlockStep.Step == nil ||
		!stringutil.CheckContains(this.BlockStep.PreviousBlockStep.Step.WorkStepType, []string{"if", "elif"}) {
		panic(errors.New(fmt.Sprintf(`previous step is not if or elif node for %s`, this.BlockStep.Step.WorkStepName)))
	}
	if this.BlockStep.HasChildren {
		bsoRunner := BlockStepOrdersRunner{
			ParentStepId: this.WorkStep.WorkStepId,
			CacheContext: this.CacheContext,
			TrackingId:   trackingId,
			Logwriter:    this.LogWriter,
			Store:        this.DataStore, // 获取数据中心
			Dispatcher:   nil,
			RunOneStep:   this.BlockStepRunFunc,
		}
		bsoRunner.Run()
	}
}
