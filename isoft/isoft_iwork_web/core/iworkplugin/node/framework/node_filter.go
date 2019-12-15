package framework

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type DoErrorFilterNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DoErrorFilterNode) Execute(trackingId string) {
	this.DataStore.CacheDatas(iworkconst.DO_ERROR_FILTER, map[string]interface{}{iworkconst.DO_ERROR_FILTER: this.TmpDataMap})
}

func (this *DoErrorFilterNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {"headerCode", "响应头 code"},
	}
	return this.BPIS1(paramMap)
}
