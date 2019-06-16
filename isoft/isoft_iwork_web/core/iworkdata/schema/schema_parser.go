package schema

import (
	"encoding/json"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/interfaces"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type WorkStepFactorySchemaParser struct {
	WorkStep          *models.WorkStep
	ParamSchemaParser interfaces.IParamSchemaParser
}

// 获取缓存的出参 schema,即从 DB 中读取
func (this *WorkStepFactorySchemaParser) GetCacheParamOutputSchema() *iworkmodels.ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(this.WorkStep.WorkStepOutput) != "" {
		var paramOutputSchema *iworkmodels.ParamOutputSchema
		if err := json.Unmarshal([]byte(this.WorkStep.WorkStepOutput), &paramOutputSchema); err == nil {
			return paramOutputSchema
		}
	}
	return &iworkmodels.ParamOutputSchema{}
}

// 获取出参 schema
func (this *WorkStepFactorySchemaParser) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.ParamSchemaParser.GetRuntimeParamOutputSchema()
}

func (this *WorkStepFactorySchemaParser) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.ParamSchemaParser.GetDefaultParamOutputSchema()
}

// 获取入参 schema
func (this *WorkStepFactorySchemaParser) GetCacheParamInputSchema() *iworkmodels.ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(this.WorkStep.WorkStepInput) != "" {
		var paramInputSchema *iworkmodels.ParamInputSchema
		if err := json.Unmarshal([]byte(this.WorkStep.WorkStepInput), &paramInputSchema); err == nil {
			return paramInputSchema
		}
	}
	// 获取当前 work_step 对应的 paramInputSchema
	return this.ParamSchemaParser.GetDefaultParamInputSchema()
}

// 获取默认入参 schema
func (this *WorkStepFactorySchemaParser) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return this.ParamSchemaParser.GetDefaultParamInputSchema()
}

// 获取入参 schema
func (this *WorkStepFactorySchemaParser) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return this.ParamSchemaParser.GetRuntimeParamInputSchema()
}
