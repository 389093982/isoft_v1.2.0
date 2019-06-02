package schema

import (
	"encoding/json"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iwork_web/models/iwork"
	"strings"
)

// 获取缓存的出参 schema,即从 DB 中读取
func GetCacheParamOutputSchema(step *iwork.WorkStep) *iworkmodels.ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepOutput) != "" {
		var paramOutputSchema *iworkmodels.ParamOutputSchema
		if err := json.Unmarshal([]byte(step.WorkStepOutput), &paramOutputSchema); err == nil {
			return paramOutputSchema
		}
	}
	return &iworkmodels.ParamOutputSchema{}
}

// 获取出参 schema
func GetRuntimeParamOutputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamOutputSchema {
	return paramSchemaParser.GetRuntimeParamOutputSchema()
}

func GetDefaultParamOutputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamOutputSchema {
	return paramSchemaParser.GetDefaultParamOutputSchema()
}

// 获取入参 schema
func GetCacheParamInputSchema(step *iwork.WorkStep, paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepInput) != "" {
		var paramInputSchema *iworkmodels.ParamInputSchema
		if err := json.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err == nil {
			return paramInputSchema
		}
	}
	// 获取当前 work_step 对应的 paramInputSchema
	return paramSchemaParser.GetDefaultParamInputSchema()
}

// 获取默认入参 schema
func GetDefaultParamInputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamInputSchema {
	return paramSchemaParser.GetDefaultParamInputSchema()
}

// 获取入参 schema
func GetRuntimeParamInputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamInputSchema {
	return paramSchemaParser.GetRuntimeParamInputSchema()
}
