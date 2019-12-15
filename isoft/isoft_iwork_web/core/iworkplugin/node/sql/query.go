package sql

import (
	"errors"
	"fmt"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkfunc"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkpool"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
	"reflect"
	"strings"
)

type SQLQueryNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	// 需要存储的中间数据
	paramMap := make(map[string]interface{}, 0)
	// 三种 sql
	sql, namings := parseNamingSql(this.TmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string))
	total_sql := fmt.Sprintf(`select count(*) as count from (%s) ttt`, sql)
	limit_sql := fmt.Sprintf(`%s limit ?,?`, sql)
	// 数据源
	dataSourceName := this.TmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(this.TmpDataMap, namings)

	var (
		totalcount int64                    // 分页查询时总数据量
		datacounts int64                    // 实际查询时的总数据量
		rowDatas   []map[string]interface{} // 查询出来的数据
	)
	// 判断是简单查询还是分页查询
	current_page := this.TmpDataMap[iworkconst.NUMBER_PREFIX+"current_page?"]
	page_size := this.TmpDataMap[iworkconst.NUMBER_PREFIX+"page_size?"]

	isPage := false
	if current_page != nil && page_size != nil {
		_current_page := datatypeutil.InterfaceConvertToInt(current_page)
		_page_size := datatypeutil.InterfaceConvertToInt(page_size)
		if _current_page > 0 && _page_size > 0 { // 正数才表示分页
			totalcount = sqlutil.QuerySelectCount(total_sql, sql_binding, dataSourceName)
			sql_binding = append(sql_binding, (_current_page-1)*_page_size, _page_size)
			datacounts, rowDatas = sqlutil.Query(limit_sql, sql_binding, dataSourceName)
			// 存储分页信息
			paginator := pageutil.Paginator(_current_page, _page_size, totalcount)
			paramMap[iworkconst.COMPLEX_PREFIX+"paginator"] = paginator
			for key, value := range paginator {
				paramMap[iworkconst.COMPLEX_PREFIX+"paginator."+key] = value
			}
			isPage = true
		}
	}
	if !isPage {
		datacounts, rowDatas = sqlutil.Query(sql, sql_binding, dataSourceName)
	}
	// 将数据数据存储到数据中心
	// 存储 datacounts
	paramMap[iworkconst.NUMBER_PREFIX+"datacounts"] = datacounts
	// 数组对象整体存储在 rows 里面
	paramMap["rows"] = rowDatas
	if len(rowDatas) > 0 {
		paramMap["row"] = rowDatas[0]
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
	this.checkPanicNoDataCount(datacounts)
}

// 当影响条数为 0 时,是否报出异常信息
func (this *SQLQueryNode) checkPanicNoDataCount(datacounts int64) {
	if panicNoAffected := this.TmpDataMap[iworkconst.STRING_PREFIX+"panic_no_datacounts?"]; panicNoAffected != nil {
		panicNoAffectedMsg := panicNoAffected.(string)
		if datacounts == 0 && panicNoAffectedMsg != "" {
			panic(&interfaces.InsensitiveError{Error: errors.New(panicNoAffectedMsg)})
		}
	}
}

func (this *SQLQueryNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sql", "查询sql语句,带分页条件的sql"},
		2: {iworkconst.STRING_PREFIX + "columnNames?", "查询结果集列名列表,以逗号分隔,动态sql需要提供"},
		3: {iworkconst.STRING_PREFIX + "metadata_sql?", "查询 metadata 所需 sql"},
		4: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数和sql中的?数量相同"},
		5: {iworkconst.NUMBER_PREFIX + "current_page?", "当前页数"},
		6: {iworkconst.NUMBER_PREFIX + "page_size?", "每页数据量"},
		7: {iworkconst.STRING_PREFIX + "panic_no_datacounts?", "查询数据量为 0 时,抛出的异常信息,为空时不抛出异常!"},
		8: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return this.BPIS1(paramMap)
}

func (this *SQLQueryNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *SQLQueryNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.NUMBER_PREFIX + "datacounts", "rows", "row"})
}

func (this *SQLQueryNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	// 输出 metadata
	pos, _ := getMetaDataQuietlyForQuery(this.WorkStep)
	// 输出分页信息
	current_page := param.GetStaticParamValueWithStep(iworkconst.NUMBER_PREFIX+"current_page?", this.WorkStep).(string)
	page_size := param.GetStaticParamValueWithStep(iworkconst.NUMBER_PREFIX+"page_size?", this.WorkStep).(string)
	if current_page != "" && page_size != "" {
		items := make([]iworkmodels.ParamOutputSchemaItem, 0)
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: iworkconst.COMPLEX_PREFIX + "paginator"})
		for _, paginatorField := range pageutil.GetPaginatorFields() {
			items = append(items, iworkmodels.ParamOutputSchemaItem{
				ParentPath: iworkconst.COMPLEX_PREFIX + "paginator",
				ParamName:  paginatorField,
			})
		}
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, items...)
	}
	return pos
}

func (this *SQLQueryNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	//validateSqlBindingParamCount(this.WorkStep)
	//validateTotalSqlBindingParamCount(this.WorkStep)
	return
}

func validateTotalSqlBindingParamCount(step *models.WorkStep) {
	total_sql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"total_sql?", step).(string)
	sql_binding := param.GetStaticParamValueWithStep(iworkconst.MULTI_PREFIX+"sql_binding?", step).(string)
	if strings.Count(total_sql, "?")+2 != strings.Count(iworkfunc.EncodeSpecialForParamVaule(sql_binding), ";") {
		panic("Number of ? in total_sql and number of ; in sql_binding is mismatch!")
	}
}

func validateSqlBindingParamCount(step *models.WorkStep) {
	sql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"sql", step).(string)
	sql_binding := param.GetStaticParamValueWithStep(iworkconst.MULTI_PREFIX+"sql_binding?", step).(string)
	if strings.Count(sql, "?") != strings.Count(iworkfunc.EncodeSpecialForParamVaule(sql_binding), ";") {
		panic("Number of ? in SQL and number of ; in sql_binding is unequal!")
	}
}

func validateAndGetDataStoreName(step *models.WorkStep) string {
	dataSourceName := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"db_conn", step).(string)
	if strings.TrimSpace(dataSourceName) == "" {
		panic("Invalid param for db_conn! Can't resolve it!")
	}
	_, err := iworkpool.GetDBConn("mysql", dataSourceName) // 全局 db 不能 Close
	if err != nil {
		panic(fmt.Sprintf("Can't get DB connection for %s!", dataSourceName))
	}
	return dataSourceName
}

func renderMetaData(columnNames []string) *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, columnName := range columnNames {
		items = append(items, iworkmodels.ParamOutputSchemaItem{
			ParentPath: "rows",
			ParamName:  columnName,
		})
		items = append(items, iworkmodels.ParamOutputSchemaItem{
			ParentPath: "row",
			ParamName:  columnName,
		})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func getMetaDataQuietlyForQuery(step *models.WorkStep) (*iworkmodels.ParamOutputSchema, []string) {
	var columnNames, namings []string
	columnNamesStr := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"columnNames?", step).(string)
	if columnNamesStr != "" {
		columnNames = strings.Split(columnNamesStr, ",")
	} else {
		metadataSql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"metadata_sql?", step).(string)
		if strings.TrimSpace(metadataSql) == "" {
			metadataSql = param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"sql", step).(string)
		}
		metadataSql, namings = parseNamingSql(metadataSql)
		dataSourceName := validateAndGetDataStoreName(step)
		columnNames = sqlutil.GetMetaDatas(metadataSql, dataSourceName)
	}
	return renderMetaData(columnNames), namings
}

// 从 tmpDataMap 获取 sql_binding 数据
func getSqlBinding(tmpDataMap map[string]interface{}, namings []string) []interface{} {
	result := make([]interface{}, 0)
	sql_binding := tmpDataMap[iworkconst.MULTI_PREFIX+"sql_binding?"]
	if sql_binding == nil {
		return result
	}

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
	} else if t1.Kind() == reflect.Map {
		if datas, ok := sql_binding.(map[string]interface{}); ok {
			for key, value := range datas {
				for _, name := range namings {
					if name == ":"+key {
						result = append(result, value)
					}
				}
			}
		}
	} else {
		result = append(result, sql_binding)
	}
	return result
}
