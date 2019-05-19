package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"os"
)

type GetEnvNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *GetEnvNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	env_var_value := os.Getenv(tmpDataMap[iworkconst.STRING_PREFIX+"env_var_name"].(string))
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "env_var_value": env_var_value})
}

func (this *GetEnvNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "env_var_name", "环境变量名称"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *GetEnvNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "env_var_value"})
}

type SetEnvNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SetEnvNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	env_var_name := tmpDataMap[iworkconst.STRING_PREFIX+"env_var_name"].(string)
	env_var_value := tmpDataMap[iworkconst.STRING_PREFIX+"env_var_value"].(string)
	if err := os.Setenv(env_var_name, env_var_value); err != nil {
		panic(err)
	}
}

func (this *SetEnvNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "env_var_name", "环境变量名称"},
		2: {iworkconst.STRING_PREFIX + "env_var_value", "环境变量值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}
