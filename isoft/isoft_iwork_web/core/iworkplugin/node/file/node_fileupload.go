package file

import (
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type DoReceiveFileNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoReceiveFileNode) Execute(trackingId string) {
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{
		"filename": this.Dispatcher.TmpDataMap["__filename"],
		"fileExt":  this.Dispatcher.TmpDataMap["__fileExt"],
		"filepath": this.Dispatcher.TmpDataMap["__filepath"],
	})
}

func (this *DoReceiveFileNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: make([]iworkmodels.ParamInputSchemaItem, 0)}
}

func (this *DoReceiveFileNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{"filename", "fileExt", "filepath"})
}
