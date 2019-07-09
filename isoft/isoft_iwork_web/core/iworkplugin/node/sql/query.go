package sql

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
	"reflect"
)

type SQLQueryNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	paramMap := make(map[string]interface{}, 0)
	sql := this.TmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string)
	dataSourceName := this.TmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(this.TmpDataMap)
	//sql = strings.ReplaceAll(sql, "{{", "")
	//sql = strings.ReplaceAll(sql, "}}", "")
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

func (this *SQLQueryNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sql", "查询sql语句,可以使用 {{}} 标识出 metadata 字段名和 tableName 表名,用于自动推断 metadatasql"},
		2: {iworkconst.STRING_PREFIX + "metadata_sql?", "元数据sql语句,针对复杂查询sql,需要提供类似于select * from blog where 1=0的辅助sql用来构建节点输出"},
		3: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数必须和当前执行sql语句中的占位符参数个数相同"},
		4: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLQueryNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{iworkconst.NUMBER_PREFIX + "datacounts"})
}

func (this *SQLQueryNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return getMetaDataQuietlyForQuery(this.WorkStep)
}

func (this *SQLQueryNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	validateAndGetMetaDataSql(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	return []string{}
}

// 从 tmpDataMap 获取 sql_binding 数据
func getSqlBinding(tmpDataMap map[string]interface{}) []interface{} {
	result := make([]interface{}, 0)
	sql_binding := tmpDataMap[iworkconst.MULTI_PREFIX+"sql_binding?"]
	t1 := reflect.TypeOf(sql_binding)
	v1 := reflect.ValueOf(sql_binding)
	// 支持单层切片和双层切片
	if t1.Kind() == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			v2 := v1.Index(i)
			if v2.Kind() == reflect.Slice {
				for j := 0; j < v2.Len(); j++ {
					result = append(result, v2.Index(j))
				}
			} else if v2.Kind() == reflect.Interface {
				if datas, ok := v2.Interface().([]interface{}); ok {
					for _, data := range datas {
						result = append(result, data)
					}
				} else {
					result = append(result, v2.Interface())
				}
			} else {
				result = append(result, v2)
			}
		}
	} else if t1.Kind() == reflect.Interface {
		if datas, ok := sql_binding.([]interface{}); ok {
			for _, data := range datas {
				result = append(result, data)
			}
		} else {
			result = append(result, sql_binding)
		}
	} else {
		result = append(result, sql_binding)
	}
	return result
}
