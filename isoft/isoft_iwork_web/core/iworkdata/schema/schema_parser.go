package schema

import (
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/interfaces"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type WorkStepFactoryParamSchemaParser struct {
	WorkStep          *models.WorkStep
	ParamSchemaParser interfaces.IParamSchemaParser
}

// 获取缓存的出参 schema,即从 DB 中读取
func (this *WorkStepFactoryParamSchemaParser) GetCacheParamOutputSchema() *iworkmodels.ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(this.WorkStep.WorkStepOutput) != "" {
		if paramOutputSchema, err := iworkmodels.ParseToParamOutputSchema(this.WorkStep.WorkStepOutput); err == nil {
			return paramOutputSchema
		}
	}
	return &iworkmodels.ParamOutputSchema{}
}

// 获取出参 schema
func (this *WorkStepFactoryParamSchemaParser) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.ParamSchemaParser.GetRuntimeParamOutputSchema()
}

func (this *WorkStepFactoryParamSchemaParser) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.ParamSchemaParser.GetDefaultParamOutputSchema()
}

// 获取入参 schema
func (this *WorkStepFactoryParamSchemaParser) GetCacheParamInputSchema() *iworkmodels.ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(this.WorkStep.WorkStepInput) != "" {
		if paramInputSchema, err := iworkmodels.ParseToParamInputSchema(this.WorkStep.WorkStepInput); err == nil {
			return paramInputSchema
		}
	}

	// 获取当前 work_step 对应的 paramInputSchema
	return this.ParamSchemaParser.GetDefaultParamInputSchema()
}

// 获取默认入参 schema
func (this *WorkStepFactoryParamSchemaParser) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return this.ParamSchemaParser.GetDefaultParamInputSchema()
}

// 获取入参 schema
func (this *WorkStepFactoryParamSchemaParser) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return this.ParamSchemaParser.GetRuntimeParamInputSchema()
}
