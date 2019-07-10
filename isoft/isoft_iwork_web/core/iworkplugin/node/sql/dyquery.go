package sql

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
)

type SQLDynamicQueryNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *SQLDynamicQueryNode) Execute(trackingId string) {
	paramMap := make(map[string]interface{}, 0)
	sql := this.TmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string)
	dataSourceName := this.TmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(this.TmpDataMap)
	datacounts, rowDatas := sqlutil.Query(sql, sql_binding, dataSourceName)
	// 存储 datacounts
	paramMap[iworkconst.NUMBER_PREFIX+"datacounts"] = datacounts
	// 数组对象整体存储在 rows 里面
	paramMap["rows"] = rowDatas
	if len(rowDatas) > 0 {
		paramMap["row"] = rowDatas[0]
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *SQLDynamicQueryNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sql", "查询sql语句"},
		2: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数必须和当前执行sql语句中的占位符参数个数相同"},
		3: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLDynamicQueryNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{iworkconst.NUMBER_PREFIX + "datacounts"})
}

func (this *SQLDynamicQueryNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return getMetaDataQuietlyForQuery(this.WorkStep)
}

func (this *SQLDynamicQueryNode) ValidateCustom() (checkResult []string) {
	return []string{}
}
