package chiper

import (
	"isoft/isoft/common/hashutil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type CalHashNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *CalHashNode) Execute(trackingId string) {
	str_data := this.TmpDataMap[iworkconst.STRING_PREFIX+"str_data"].(string)
	hash := hashutil.CalculateHashWithString(str_data)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "hash": hash})
}

func (this *CalHashNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "str_data", "需要计算hash值的字符串"},
	}
	return this.BPIS1(paramMap)
}

func (this *CalHashNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "hash"})
}
