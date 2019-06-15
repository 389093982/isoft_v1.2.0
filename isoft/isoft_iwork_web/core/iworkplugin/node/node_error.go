package node

import (
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/models/iwork"
)

type PanicErrorNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *PanicErrorNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"panic_expression"].(bool)
	if expression {
		errorMsg := tmpDataMap[iworkconst.STRING_PREFIX+"panic_errorMsg?"].(string)
		this.DataStore.CacheDatas("Error", map[string]interface{}{"errorMsg": errorMsg})
		panic(errors.New(errorMsg))
	}
}

func (this *PanicErrorNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.BOOL_PREFIX + "panic_expression", "抛出异常的条件,值为 bool 类型!"},
		2: {iworkconst.STRING_PREFIX + "panic_errorMsg?", "抛出异常的信息,值为字符串类型!"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}
