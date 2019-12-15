package os

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"os"
)

type GetEnvNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *GetEnvNode) Execute(trackingId string) {
	env_var_value := os.Getenv(this.TmpDataMap[iworkconst.STRING_PREFIX+"env_var_name"].(string))
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "env_var_value": env_var_value})
}

func (this *GetEnvNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "env_var_name", "环境变量名称"},
	}
	return this.BPIS1(paramMap)
}

func (this *GetEnvNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "env_var_value"})
}

type SetEnvNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *SetEnvNode) Execute(trackingId string) {
	env_var_name := this.TmpDataMap[iworkconst.STRING_PREFIX+"env_var_name"].(string)
	env_var_value := this.TmpDataMap[iworkconst.STRING_PREFIX+"env_var_value"].(string)
	if err := os.Setenv(env_var_name, env_var_value); err != nil {
		panic(err)
	}
}

func (this *SetEnvNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "env_var_name", "环境变量名称"},
		2: {iworkconst.STRING_PREFIX + "env_var_value", "环境变量值"},
	}
	return this.BPIS1(paramMap)
}
