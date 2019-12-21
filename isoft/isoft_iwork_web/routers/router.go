package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"isoft/isoft_iwork_web/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Include(&controllers.WorkController{})

	// 通用代理路由
	beego.Router("/api/iwork/proxyCall", &controllers.WorkController{}, "get,post:ProxyCall")
	loadloadIWorkerRouterDetail()
}

func loadloadIWorkerRouterDetail() {
	wc := &controllers.WorkController{}
	registRouter("/api/iwork/addQuartz", wc, wc.AddQuartz, "get,post:AddQuartz")
	registRouter("/api/iwork/filterPageQuartz", wc, wc.FilterPageQuartz, "get,post:FilterPageQuartz")
	registRouter("/api/iwork/editQuartz", wc, wc.EditQuartz, "get,post:EditQuartz")
	registRouter("/api/iwork/editResource", wc, wc.EditResource, "get,post:EditResource")
	registRouter("/api/iwork/filterPageResource", wc, wc.FilterPageResource, "get,post:FilterPageResource")
	registRouter("/api/iwork/deleteResource", wc, wc.DeleteResource, "get,post:DeleteResource")
	registRouter("/api/iwork/validateResource", wc, wc.ValidateResource, "get,post:ValidateResource")
	registRouter("/api/iwork/getAllResource", wc, wc.GetAllResource, "get,post:GetAllResource")
	registRouter("/api/iwork/getResourceById", wc, wc.GetResourceById, "get,post:GetResourceById")

	registRouter("/api/iwork/filterPageWorkHistory", wc, wc.FilterPageWorkHistory, "get,post:FilterPageWorkHistory")
	registRouter("/api/iwork/restoreFromWorkHistory", wc, wc.RestoreFromWorkHistory, "get,post:RestoreFromWorkHistory")
	registRouter("/api/iwork/filterPageWorks", wc, wc.FilterPageWorks, "get,post:FilterPageWorks")
	registRouter("/api/iwork/editWork", wc, wc.EditWork, "get,post:EditWork")
	registRouter("/api/iwork/deleteOrCopyWorkById", wc, wc.DeleteOrCopyWorkById, "get,post:DeleteOrCopyWorkById")
	registRouter("/api/iwork/addWorkStep", wc, wc.AddWorkStep, "get,post:AddWorkStep")
	registRouter("/api/iwork/editWorkStepBaseInfo", wc, wc.EditWorkStepBaseInfo, "get,post:EditWorkStepBaseInfo")
	registRouter("/api/iwork/deleteWorkStepByWorkStepId", wc, wc.DeleteWorkStepByWorkStepId, "get,post:DeleteWorkStepByWorkStepId")
	registRouter("/api/iwork/copyWorkStepByWorkStepId", wc, wc.CopyWorkStepByWorkStepId, "get,post:CopyWorkStepByWorkStepId")
	registRouter("/api/iwork/loadWorkStepInfo", wc, wc.LoadWorkStepInfo, "get,post:LoadWorkStepInfo")
	registRouter("/api/iwork/getAllWorkStepInfo", wc, wc.GetAllWorkStepInfo, "get,post:GetAllWorkStepInfo")
	registRouter("/api/iwork/changeWorkStepOrder", wc, wc.ChangeWorkStepOrder, "get,post:ChangeWorkStepOrder")
	registRouter("/api/iwork/runWork", wc, wc.RunWork, "get,post:RunWork")
	registRouter("/api/iwork/loadPreNodeOutput", wc, wc.LoadPreNodeOutput, "get,post:LoadPreNodeOutput")
	registRouter("/api/iwork/filterPageLogRecord", wc, wc.FilterPageLogRecord, "get,post:FilterPageLogRecord")
	registRouter("/api/iwork/getLastRunLogDetail", wc, wc.GetLastRunLogDetail, "get,post:GetLastRunLogDetail")
	registRouter("/api/iwork/httpservice/:work_name", wc, wc.PublishSerivce, "get,post:PublishSerivce")
	registRouter("/api/iwork/getRelativeWork", wc, wc.GetRelativeWork, "get,post:GetRelativeWork")
	registRouter("/api/iwork/buildIWorkDL", wc, wc.BuildIWorkDL, "get,post:BuildIWorkDL")
	registRouter("/api/iwork/validateWork", wc, wc.ValidateWork, "get,post:ValidateWork")
	registRouter("/api/iwork/refactorWorkStepInfo", wc, wc.RefactorWorkStepInfo, "get,post:RefactorWorkStepInfo")
	registRouter("/api/iwork/batchChangeIndent", wc, wc.BatchChangeIndent, "get,post:BatchChangeIndent")
	registRouter("/api/iwork/parseToMultiValue", wc, wc.ParseToMultiValue, "get,post:ParseToMultiValue")
	registRouter("/api/iwork/workStepList", wc, wc.WorkStepList, "get,post:WorkStepList")
	registRouter("/api/iwork/deleteModuleById", wc, wc.DeleteModuleById, "get,post:DeleteModuleById")
	registRouter("/api/iwork/getAllModules", wc, wc.GetAllModules, "get,post:GetAllModules")
	registRouter("/api/iwork/getLastMigrateLogs", wc, wc.GetLastMigrateLogs, "get,post:GetLastMigrateLogs")
	registRouter("/api/iwork/executeMigrate", wc, wc.ExecuteMigrate, "get,post:ExecuteMigrate")
	registRouter("/api/iwork/editSqlMigrate", wc, wc.EditSqlMigrate, "get,post:EditSqlMigrate")
	registRouter("/api/iwork/filterPageSqlMigrate", wc, wc.FilterPageSqlMigrate, "get,post:FilterPageSqlMigrate")
	registRouter("/api/iwork/getSqlMigrateInfo", wc, wc.GetSqlMigrateInfo, "get,post:GetSqlMigrateInfo")
	registRouter("/api/iwork/toggleSqlMigrateEffective", wc, wc.ToggleSqlMigrateEffective, "get,post:ToggleSqlMigrateEffective")
	registRouter("/api/iwork/editWorkStepParamInfo", wc, wc.EditWorkStepParamInfo, "get,post:EditWorkStepParamInfo")

	registRouter("/api/iwork/loadQuickSqlMeta", wc, wc.LoadQuickSqlMeta, "get,post:LoadQuickSqlMeta")

	registRouter("/api/iwork/globalVarList", wc, wc.GlobalVarList, "get,post:GlobalVarList")
	registRouter("/api/iwork/editGlobalVar", wc, wc.EditGlobalVar, "get,post:EditGlobalVar")
	registRouter("/api/iwork/deleteGlobalVarById", wc, wc.DeleteGlobalVarById, "get,post:DeleteGlobalVarById")

	registRouter("/api/iwork/editTemplate", wc, wc.EditTemplate, "get,post:EditTemplate")
	registRouter("/api/iwork/templateList", wc, wc.TemplateList, "get,post:TemplateList")
	registRouter("/api/iwork/deleteTemplateById", wc, wc.DeleteTemplateById, "get,post:DeleteTemplateById")

	registRouter("/api/iwork/download/:work_id", wc, wc.Download, "get,post:Download")

	registRouter("/api/iwork/moduleList", wc, wc.ModuleList, "get,post:ModuleList")
	registRouter("/api/iwork/editModule", wc, wc.EditModule, "get,post:EditModule")
	registRouter("/api/iwork/getAllFiltersAndWorks", wc, wc.GetAllFiltersAndWorks, "get,post:GetAllFiltersAndWorks")
	registRouter("/api/iwork/saveFilters", wc, wc.SaveFilters, "get,post:SaveFilters")
	registRouter("/api/iwork/getMetaInfo", wc, wc.GetMetaInfo, "get,post:GetMetaInfo")
	registRouter("/api/iwork/queryWorkDetail", wc, wc.QueryWorkDetail, "get,post:QueryWorkDetail")
	registRouter("/api/iwork/loadValidateResult", wc, wc.LoadValidateResult, "get,post:LoadValidateResult")
	registRouter("/api/iwork/saveProject", wc, wc.SaveProject, "get,post:SaveProject")
	registRouter("/api/iwork/importProject", wc, wc.ImportProject, "get,post:ImportProject")

	registRouter("/api/iwork/editAuditTask", wc, wc.EditAuditTask, "get,post:EditAuditTask")
	registRouter("/api/iwork/queryPageAuditTask", wc, wc.QueryPageAuditTask, "get,post:QueryPageAuditTask")
	registRouter("/api/iwork/editAuditTaskSource", wc, wc.EditAuditTaskSource, "get,post:EditAuditTaskSource")
	registRouter("/api/iwork/queryTaskDetail", wc, wc.QueryTaskDetail, "get,post:QueryTaskDetail")
	registRouter("/api/iwork/editAuditTaskTarget", wc, wc.EditAuditTaskTarget, "get,post:EditAuditTaskTarget")
	registRouter("/api/iwork/getAuditHandleData", wc, wc.GetAuditHandleData, "get,post:GetAuditHandleData")
	registRouter("/api/iwork/executeAuditTask", wc, wc.ExecuteAuditTask, "get,post:ExecuteAuditTask")
	registRouter("/api/iwork/deleteAuditTask", wc, wc.DeleteAuditTask, "get,post:DeleteAuditTask")

	registRouter("/api/iwork/filterPagePlacement", wc, wc.FilterPagePlacement, "get,post:FilterPagePlacement")
	registRouter("/api/iwork/editPlacement", wc, wc.EditPlacement, "get,post:EditPlacement")
	registRouter("/api/iwork/queryPlacementById", wc, wc.QueryPlacementById, "get,post:QueryPlacementById")
	registRouter("/api/iwork/queryPlacementByName", wc, wc.QueryPlacementByName, "get,post:QueryPlacementByName")
	registRouter("/api/iwork/deletePlacementById", wc, wc.DeletePlacementById, "get,post:DeletePlacementById")
	registRouter("/api/iwork/copyPlacement", wc, wc.CopyPlacement, "get,post:CopyPlacement")
	registRouter("/api/iwork/filterPageElement", wc, wc.FilterPageElement, "get,post:FilterPageElement")
	registRouter("/api/iwork/editElement", wc, wc.EditElement, "get,post:EditElement")
	registRouter("/api/iwork/updateElementStatus", wc, wc.UpdateElementStatus, "get,post:UpdateElementStatus")
	registRouter("/api/iwork/queryElementById", wc, wc.QueryElementById, "get,post:QueryElementById")
	registRouter("/api/iwork/copyElement", wc, wc.CopyElement, "get,post:CopyElement")
	registRouter("/api/iwork/filterElementByPlacement", wc, wc.FilterElementByPlacement, "get,post:FilterElementByPlacement")

	registRouter("/api/iwork/import", wc, wc.Import, "get,post:Import")
}

func registRouter(rootpath string, c beego.ControllerInterface, callFunc func(), mappingMethods ...string) *beego.App {
	return beego.Router(rootpath, c, mappingMethods...)
}
