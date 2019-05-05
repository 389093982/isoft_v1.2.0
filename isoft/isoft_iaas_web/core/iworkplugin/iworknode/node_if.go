package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type IFNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(trackingId string, logwriter *iworklog.CacheLoggerWriter, blockStep *block.BlockStep, datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *IFNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.BOOL_PREFIX + "expression": expression})

	if expression && this.BlockStep.HasChildren {
		order := this.getChildBlockStepExecuteOrder()
		for _, blockStep := range order {
			this.BlockStepRunFunc(trackingId, this.LogWriter, blockStep, this.DataStore, nil)
		}
	} else {
		this.LogWriter.Write(trackingId, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *IFNode) getChildBlockStepExecuteOrder() []*block.BlockStep {
	order := make([]*block.BlockStep, 0)
	deferOrder := make([]*block.BlockStep, 0)
	for _, blockStep := range this.BlockStep.ChildBlockSteps {
		if blockStep.Step.IsDefer == "true" {
			deferOrder = append(deferOrder, blockStep)
		} else {
			order = append(order, blockStep)
		}
	}
	order = append(order, datatypeutil.ReverseSlice(deferOrder).([]*block.BlockStep)...)
	return order
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
	BlockStepRunFunc func(trackingId string, logwriter *iworklog.CacheLoggerWriter, blockStep *block.BlockStep, datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *ElIfNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.BOOL_PREFIX + "expression": expression})

	if expression && this.BlockStep.HasChildren {
		order := this.getChildBlockStepExecuteOrder()
		for _, blockStep := range order {
			this.BlockStepRunFunc(trackingId, this.LogWriter, blockStep, this.DataStore, nil)
		}
	} else {
		this.LogWriter.Write(trackingId, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *ElIfNode) getChildBlockStepExecuteOrder() []*block.BlockStep {
	order := make([]*block.BlockStep, 0)
	deferOrder := make([]*block.BlockStep, 0)
	for _, blockStep := range this.BlockStep.ChildBlockSteps {
		if blockStep.Step.IsDefer == "true" {
			deferOrder = append(deferOrder, blockStep)
		} else {
			order = append(order, blockStep)
		}
	}
	order = append(order, datatypeutil.ReverseSlice(deferOrder).([]*block.BlockStep)...)
	return order
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
	BlockStepRunFunc func(trackingId string, logwriter *iworklog.CacheLoggerWriter, blockStep *block.BlockStep, datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *ElseNode) Execute(trackingId string) {
	// 节点中间数据
	//tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	if this.BlockStep.HasChildren {
		order := this.getChildBlockStepExecuteOrder()
		for _, blockStep := range order {
			this.BlockStepRunFunc(trackingId, this.LogWriter, blockStep, this.DataStore, nil)
		}
	}
}

func (this *ElseNode) getChildBlockStepExecuteOrder() []*block.BlockStep {
	order := make([]*block.BlockStep, 0)
	deferOrder := make([]*block.BlockStep, 0)
	for _, blockStep := range this.BlockStep.ChildBlockSteps {
		if blockStep.Step.IsDefer == "true" {
			deferOrder = append(deferOrder, blockStep)
		} else {
			order = append(order, blockStep)
		}
	}
	order = append(order, datatypeutil.ReverseSlice(deferOrder).([]*block.BlockStep)...)
	return order
}
