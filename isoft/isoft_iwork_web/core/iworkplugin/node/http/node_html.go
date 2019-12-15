package http

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/htmlutil"
	"isoft/isoft_iwork_web/models"
)

type HrefParserNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *HrefParserNode) Execute(trackingId string) {
	hrefs := make([]interface{}, 0)
	if url, ok := this.TmpDataMap[iworkconst.STRING_PREFIX+"url"].(string); ok {
		if _hrefs := htmlutil.GetAllHref(url); len(_hrefs) > 0 {
			// 将 []string 转换成 []interface{}
			hrefs = stringutil.ChangeStringsToInterfaces(_hrefs)
		}
	}
	// 放在外面保证条件不满足时也是零值,不报空指针异常
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{
		iworkconst.MULTI_PREFIX + "hrefs":         hrefs,
		iworkconst.NUMBER_PREFIX + "href_amounts": len(hrefs),
	})
}

func (this *HrefParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.STRING_PREFIX + "url", "需要分析资源的url地址"},
	}
	return this.BPIS1(paramMap)
}

func (this *HrefParserNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.MULTI_PREFIX + "hrefs", iworkconst.NUMBER_PREFIX + "href_amounts"})
}
