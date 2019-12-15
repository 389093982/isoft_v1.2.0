package sql

import (
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/sqlutil"
	"isoft/isoft_iwork_web/models"
)

type DBParserNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *DBParserNode) Execute(trackingId string) {
	dataSourceName := this.TmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	tableNames := sqlutil.GetAllTableNames(dataSourceName)
	paramMap := map[string]interface{}{"tableNames": tableNames}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *DBParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return this.BPIS1(paramMap)
}

func (this *DBParserNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{"tableNames"})
}
