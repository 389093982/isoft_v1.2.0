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
	beego.Router("/", &controllers.MainController{})

	// 通用代理路由
	beego.Router("/api/iwork/proxyCall", &controllers.WorkController{}, "post:ProxyCall")
	loadloadIWorkerRouterDetail()
}

// egg: // @router /api/iwork/deleteModuleById [post]
func loadloadIWorkerRouterDetail() {
	beego.Router("/api/iwork/addQuartz", &controllers.WorkController{}, "post:AddQuartz")
	beego.Router("/api/iwork/filterPageQuartz", &controllers.WorkController{}, "post:FilterPageQuartz")
	beego.Router("/api/iwork/editQuartz", &controllers.WorkController{}, "post:EditQuartz")
	beego.Router("/api/iwork/addResource", &controllers.WorkController{}, "post:AddResource")
	beego.Router("/api/iwork/filterPageResource", &controllers.WorkController{}, "post:FilterPageResource")
	beego.Router("/api/iwork/deleteResource", &controllers.WorkController{}, "post:DeleteResource")
	beego.Router("/api/iwork/validateResource", &controllers.WorkController{}, "post:ValidateResource")
	beego.Router("/api/iwork/getAllResource", &controllers.WorkController{}, "post:GetAllResource")

	beego.Router("/api/iwork/filterPageWorkHistory", &controllers.WorkController{}, "post:FilterPageWorkHistory")
	beego.Router("/api/iwork/filterPageWork", &controllers.WorkController{}, "post:FilterPageWork")
	beego.Router("/api/iwork/editWork", &controllers.WorkController{}, "post:EditWork")
	beego.Router("/api/iwork/deleteOrCopyWorkById", &controllers.WorkController{}, "post:DeleteOrCopyWorkById")
	beego.Router("/api/iwork/addWorkStep", &controllers.WorkController{}, "post:AddWorkStep")
	beego.Router("/api/iwork/editWorkStepBaseInfo", &controllers.WorkController{}, "post:EditWorkStepBaseInfo")
	beego.Router("/api/iwork/deleteWorkStepByWorkStepId", &controllers.WorkController{}, "post:DeleteWorkStepByWorkStepId")
	beego.Router("/api/iwork/copyWorkStepByWorkStepId", &controllers.WorkController{}, "post:CopyWorkStepByWorkStepId")
	beego.Router("/api/iwork/loadWorkStepInfo", &controllers.WorkController{}, "post:LoadWorkStepInfo")
	beego.Router("/api/iwork/getAllWorkStepInfo", &controllers.WorkController{}, "post:GetAllWorkStepInfo")
	beego.Router("/api/iwork/changeWorkStepOrder", &controllers.WorkController{}, "post:ChangeWorkStepOrder")
	beego.Router("/api/iwork/runWork", &controllers.WorkController{}, "post:RunWork")
	beego.Router("/api/iwork/loadPreNodeOutput", &controllers.WorkController{}, "post:LoadPreNodeOutput")
	beego.Router("/api/iwork/filterPageLogRecord", &controllers.WorkController{}, "post:FilterPageLogRecord")
	beego.Router("/api/iwork/getLastRunLogDetail", &controllers.WorkController{}, "post:GetLastRunLogDetail")
	beego.Router("/api/iwork/httpservice/:work_name", &controllers.WorkController{}, "get,post:PublishSerivce")
	beego.Router("/api/iwork/getRelativeWork", &controllers.WorkController{}, "post:GetRelativeWork")
	beego.Router("/api/iwork/filterPageEntity", &controllers.WorkController{}, "post:FilterPageEntity")
	beego.Router("/api/iwork/editEntity", &controllers.WorkController{}, "post:EditEntity")
	beego.Router("/api/iwork/deleteEntity", &controllers.WorkController{}, "post:DeleteEntity")
	beego.Router("/api/iwork/buildIWorkDL", &controllers.WorkController{}, "post:BuildIWorkDL")
	beego.Router("/api/iwork/validateWork", &controllers.WorkController{}, "post:ValidateWork")
	beego.Router("/api/iwork/refactorWorkStepInfo", &controllers.WorkController{}, "post:RefactorWorkStepInfo")
	beego.Router("/api/iwork/batchChangeIndent", &controllers.WorkController{}, "post:BatchChangeIndent")
	beego.Router("/api/iwork/parseToMultiValue", &controllers.WorkController{}, "post:ParseToMultiValue")

	beego.Router("/api/iwork/loadQuickSqlMeta", &controllers.WorkController{}, "post:LoadQuickSqlMeta")

	beego.Router("/api/iwork/globalVarList", &controllers.WorkController{}, "post:GlobalVarList")
	beego.Router("/api/iwork/editGlobalVar", &controllers.WorkController{}, "post:EditGlobalVar")
	beego.Router("/api/iwork/deleteGlobalVarById", &controllers.WorkController{}, "post:DeleteGlobalVarById")

	beego.Router("/api/iwork/editTemplate", &controllers.WorkController{}, "post:EditTemplate")
	beego.Router("/api/iwork/templateList", &controllers.WorkController{}, "post:TemplateList")
	beego.Router("/api/iwork/deleteTemplateById", &controllers.WorkController{}, "post:DeleteTemplateById")

	beego.Router("/api/iwork/download/:work_id", &controllers.WorkController{}, "get,post:Download")

	beego.Router("/api/iwork/moduleList", &controllers.WorkController{}, "get,post:ModuleList")
	beego.Router("/api/iwork/editModule", &controllers.WorkController{}, "get,post:EditModule")
	beego.Router("/api/iwork/getAllFiltersAndWorks", &controllers.WorkController{}, "get,post:GetAllFiltersAndWorks")
	beego.Router("/api/iwork/saveFilters", &controllers.WorkController{}, "get,post:SaveFilters")
	beego.Router("/api/iwork/getMetaInfo", &controllers.WorkController{}, "get,post:GetMetaInfo")
}
