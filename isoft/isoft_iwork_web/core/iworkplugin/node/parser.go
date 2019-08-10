package node

import (
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil/reflectutil"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type ParamSchemaParser struct {
	WorkStep          *models.WorkStep
	ParamSchemaParser interfaces.IParamSchemaParser
	interfaces.IParamSchemaCacheParser
}

// 获取缓存的出参 schema,即从 DB 中读取
func (this *ParamSchemaParser) GetCacheParamOutputSchema() *iworkmodels.ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(this.WorkStep.WorkStepOutput) != "" {
		if paramOutputSchema, err := iworkmodels.ParseToParamOutputSchema(this.WorkStep.WorkStepOutput); err == nil {
			return paramOutputSchema
		}
	}
	return &iworkmodels.ParamOutputSchema{}
}

// 获取出参 schema
func (this *ParamSchemaParser) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.ParamSchemaParser.GetRuntimeParamOutputSchema()
}

func (this *ParamSchemaParser) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.ParamSchemaParser.GetDefaultParamOutputSchema()
}

// 获取入参 schema
func (this *ParamSchemaParser) GetCacheParamInputSchema(replaceStep ...*models.WorkStep) *iworkmodels.ParamInputSchema {
	if len(replaceStep) > 0 {
		this.WorkStep = replaceStep[0]
	}
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(this.WorkStep.WorkStepInput) != "" {
		if paramInputSchema, err := iworkmodels.ParseToParamInputSchema(this.WorkStep.WorkStepInput); err == nil {
			return paramInputSchema
		}
	}
	if this.ParamSchemaParser != nil {
		// 将当前 paramSchemaParser 拷贝副本
		paramSchemaParser := this.ParamSchemaParser
		if len(replaceStep) > 0 {
			// 填充新参数至 ParamSchemaParser
			reflectutil.FillFieldValueToStruct(this.ParamSchemaParser, map[string]interface{}{"WorkStep": replaceStep[0]})
		}
		// 获取当前 work_step 对应的 paramInputSchema
		return paramSchemaParser.GetDefaultParamInputSchema()
	} else {
		return &iworkmodels.ParamInputSchema{}
	}
}

// 获取默认入参 schema
func (this *ParamSchemaParser) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return this.ParamSchemaParser.GetDefaultParamInputSchema()
}

// 获取入参 schema
func (this *ParamSchemaParser) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return this.ParamSchemaParser.GetRuntimeParamInputSchema()
}
