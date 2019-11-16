package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"isoft/isoft_iwork_web/controllers"
	"reflect"
	"runtime"
	"strings"
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
	wc := &controllers.WorkController{}
	beego.Router("/api/iwork/addQuartz", wc, "post:AddQuartz")
	beego.Router("/api/iwork/filterPageQuartz", wc, "post:FilterPageQuartz")
	beego.Router("/api/iwork/editQuartz", wc, "post:EditQuartz")
	beego.Router("/api/iwork/addResource", wc, "post:AddResource")
	beego.Router("/api/iwork/filterPageResource", wc, "post:FilterPageResource")
	beego.Router("/api/iwork/deleteResource", wc, "post:DeleteResource")
	beego.Router("/api/iwork/validateResource", wc, "post:ValidateResource")
	beego.Router("/api/iwork/getAllResource", wc, "post:GetAllResource")

	registRouter("/api/iwork/filterPageWorkHistory", wc, wc.FilterPageWorkHistory)
	registRouter("/api/iwork/restoreFromWorkHistory", wc, wc.RestoreFromWorkHistory)
	registRouter("/api/iwork/filterPageWorks", wc, wc.FilterPageWorks)
	beego.Router("/api/iwork/editWork", wc, "post:EditWork")
	beego.Router("/api/iwork/deleteOrCopyWorkById", wc, "post:DeleteOrCopyWorkById")
	beego.Router("/api/iwork/addWorkStep", wc, "post:AddWorkStep")
	beego.Router("/api/iwork/editWorkStepBaseInfo", wc, "post:EditWorkStepBaseInfo")
	beego.Router("/api/iwork/deleteWorkStepByWorkStepId", wc, "post:DeleteWorkStepByWorkStepId")
	beego.Router("/api/iwork/copyWorkStepByWorkStepId", wc, "post:CopyWorkStepByWorkStepId")
	beego.Router("/api/iwork/loadWorkStepInfo", wc, "post:LoadWorkStepInfo")
	beego.Router("/api/iwork/getAllWorkStepInfo", wc, "post:GetAllWorkStepInfo")
	beego.Router("/api/iwork/changeWorkStepOrder", wc, "post:ChangeWorkStepOrder")
	beego.Router("/api/iwork/runWork", wc, "post:RunWork")
	beego.Router("/api/iwork/loadPreNodeOutput", wc, "post:LoadPreNodeOutput")
	beego.Router("/api/iwork/filterPageLogRecord", wc, "post:FilterPageLogRecord")
	beego.Router("/api/iwork/getLastRunLogDetail", wc, "post:GetLastRunLogDetail")
	beego.Router("/api/iwork/httpservice/:work_name", wc, "get,post:PublishSerivce")
	beego.Router("/api/iwork/getRelativeWork", wc, "post:GetRelativeWork")
	beego.Router("/api/iwork/buildIWorkDL", wc, "post:BuildIWorkDL")
	beego.Router("/api/iwork/validateWork", wc, "post:ValidateWork")
	beego.Router("/api/iwork/refactorWorkStepInfo", wc, "post:RefactorWorkStepInfo")
	beego.Router("/api/iwork/batchChangeIndent", wc, "post:BatchChangeIndent")
	beego.Router("/api/iwork/parseToMultiValue", wc, "post:ParseToMultiValue")

	beego.Router("/api/iwork/loadQuickSqlMeta", wc, "post:LoadQuickSqlMeta")

	beego.Router("/api/iwork/globalVarList", wc, "post:GlobalVarList")
	beego.Router("/api/iwork/editGlobalVar", wc, "post:EditGlobalVar")
	beego.Router("/api/iwork/deleteGlobalVarById", wc, "post:DeleteGlobalVarById")

	beego.Router("/api/iwork/editTemplate", wc, "post:EditTemplate")
	beego.Router("/api/iwork/templateList", wc, "post:TemplateList")
	beego.Router("/api/iwork/deleteTemplateById", wc, "post:DeleteTemplateById")

	beego.Router("/api/iwork/download/:work_id", wc, "get,post:Download")

	beego.Router("/api/iwork/moduleList", wc, "get,post:ModuleList")
	beego.Router("/api/iwork/editModule", wc, "get,post:EditModule")
	beego.Router("/api/iwork/getAllFiltersAndWorks", wc, "get,post:GetAllFiltersAndWorks")
	registRouter("/api/iwork/saveFilters", wc, wc.SaveFilters)
	registRouter("/api/iwork/getMetaInfo", wc, wc.GetMetaInfo)
	registRouter("/api/iwork/queryWorkDetail", wc, wc.QueryWorkDetail)
	registRouter("/api/iwork/loadValidateResult", wc, wc.LoadValidateResult)
	registRouter("/api/iwork/saveProject", wc, wc.SaveProject)
	registRouter("/api/iwork/importProject", wc, wc.ImportProject)
}

func registRouter(rootpath string, c beego.ControllerInterface, callFunc func(), mappingMethods ...string) *beego.App {
	// 'isoft/isoft_iwork_web/controllers.(*WorkController).GetMetaInfo-fm'
	funcName := getFunctionName(callFunc)
	funcName = funcName[strings.LastIndex(funcName, ".")+1 : strings.LastIndex(funcName, "-")]
	if len(mappingMethods) > 0 {
		mappingMethods[0] = fmt.Sprintf(`%s:%s`, mappingMethods[0], funcName)
	} else {
		mappingMethods = append(mappingMethods, fmt.Sprintf(`%s:%s`, "get,post", funcName))
	}
	return beego.Router(rootpath, c, mappingMethods...)
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
