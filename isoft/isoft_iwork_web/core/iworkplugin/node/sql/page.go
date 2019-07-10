package sql

import (
	"fmt"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkfunc"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkpool"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
	"strconv"
	"strings"
)

type SQLQueryPageNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *SQLQueryPageNode) Execute(trackingId string) {
	// 需要存储的中间数据
	paramMap := make(map[string]interface{}, 0)
	sql := this.TmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string)
	var total_sql string
	if _total_sql, ok := this.TmpDataMap[iworkconst.STRING_PREFIX+"total_sql?"].(string); ok && strings.TrimSpace(_total_sql) != "" {
		total_sql = _total_sql
	} else {
		total_sql = getTotalSqlFromQuery(sql)
	}
	dataSourceName := this.TmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(this.TmpDataMap)
	totalcount := sqlutil.QuerySelectCount(total_sql, sql_binding, dataSourceName)
	sql = fmt.Sprintf(`%s limit ?,?`, sql) // 追加 limit 语句
	current_page := this.TmpDataMap[iworkconst.NUMBER_PREFIX+"current_page"].(string)
	page_size := this.TmpDataMap[iworkconst.NUMBER_PREFIX+"page_size"].(string)
	_current_page, _ := strconv.Atoi(current_page)
	_page_size, _ := strconv.Atoi(page_size)
	sql_binding = append(sql_binding, (_current_page-1)*_page_size, _page_size)
	datacounts, rowDatas := sqlutil.Query(sql, sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 datacounts
	paramMap[iworkconst.NUMBER_PREFIX+"datacounts"] = datacounts
	// 数组对象整体存储在 rows 里面
	paramMap["rows"] = rowDatas
	if len(rowDatas) > 0 {
		paramMap["row"] = rowDatas[0]
	}
	// 存储分页信息
	pageIndex, pageSize := getPageIndexAndPageSize(this.TmpDataMap)
	paginator := pageutil.Paginator(pageIndex, pageSize, totalcount)
	paramMap[iworkconst.COMPLEX_PREFIX+"paginator"] = paginator

	for key, value := range paginator {
		paramMap[iworkconst.FIELD_PREFIX+"paginator."+key] = value
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func getPageIndexAndPageSize(tmpDataMap map[string]interface{}) (currentPage int, pageSize int) {
	var convert = func(data interface{}) (result int) {
		if _data, ok := data.(string); ok {
			result, _ = strconv.Atoi(_data)
		} else if _data, ok := data.(int); ok {
			result = _data
		} else if _data, ok := data.(int64); ok {
			result = int(_data)
		}
		return
	}
	currentPage = convert(tmpDataMap[iworkconst.NUMBER_PREFIX+"current_page"])
	pageSize = convert(tmpDataMap[iworkconst.NUMBER_PREFIX+"page_size"])
	return
}

func (this *SQLQueryPageNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sql", "查询sql语句,带分页条件的sql,等价于 ${total_sql} limit ?,?"},
		2: {iworkconst.STRING_PREFIX + "columnNames?", "查询结果集列名列表,以逗号分隔,动态sql需要提供"},
		3: {iworkconst.STRING_PREFIX + "total_sql?", "统计总数sql,返回N页总数据量,格式参考select count(*) as count from blog where xxx"},
		4: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数和sql中的?数量相同,前N-2位参数和total_sql中的?数量相同"},
		5: {iworkconst.NUMBER_PREFIX + "current_page", "当前页数"},
		6: {iworkconst.NUMBER_PREFIX + "page_size", "每页数据量"},
		7: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLQueryPageNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *SQLQueryPageNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, paginatorField := range pageutil.GetPaginatorFields() {
		items = append(items, iworkmodels.ParamOutputSchemaItem{
			ParentPath: iworkconst.COMPLEX_PREFIX + "paginator",
			ParamName:  paginatorField,
		})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *SQLQueryPageNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return getMetaDataQuietlyForQuery(this.WorkStep)
}

func (this *SQLQueryPageNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	validateTotalSqlBindingParamCount(this.WorkStep)
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

func getTotalSqlFromQuery(querySql string) string {
	totalSql := fmt.Sprintf(`select count(*) as count from (%s)`, querySql)
	return totalSql
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

func getMetaDataQuietlyForQuery(step *models.WorkStep) *iworkmodels.ParamOutputSchema {
	var columnNames []string
	columnNamesStr := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"columnNames?", step).(string)
	if columnNamesStr != "" {
		columnNames = strings.Split(columnNamesStr, ",")
	} else {
		metadataSql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"sql", step).(string)
		dataSourceName := validateAndGetDataStoreName(step)
		columnNames = sqlutil.GetMetaDatas(metadataSql, dataSourceName)
	}
	return renderMetaData(columnNames)
}
