package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"isoft/isoft_iwork_web/controllers"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworkplugin/node/regist"
	"isoft/isoft_iwork_web/core/iworkpool"
	"isoft/isoft_iwork_web/core/iworkrun"
	_ "isoft/isoft_iwork_web/routers"
	_ "isoft/isoft_iwork_web/startup/db"
	_ "isoft/isoft_iwork_web/startup/logger"
	_ "isoft/isoft_iwork_web/startup/sysconfig"
	"isoft/isoft_iwork_web/startup/task"
)

func filterFunc(ctx *context.Context) {
	work_name := ctx.Input.Param(":work_name")
	parser := schema.WorkStepFactoryParamSchemaParser{}
	workCache, err := iworkcache.GetWorkCacheWithName(work_name, &parser)
	if err != nil{
		panic(err)
	}
	for _, filterName := range workCache.FilterNames{
		if workCache, err := iworkcache.GetWorkCacheWithName(filterName, &parser); err == nil{
			mapData := controllers.ParseParam(ctx, workCache.Steps)
			mapData[iworkconst.HTTP_REQUEST_OBJECT] = ctx.Request // 传递 request 对象
			receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: mapData})
			if receiver != nil {
				tempDataMap := receiver.TmpDataMap
				if doErrorFilter, ok := tempDataMap["doErrorFilter"]; ok && doErrorFilter.(string) == "ERROR" {
					ctx.ResponseWriter.WriteHeader(401)
				}
			}
		}
	}
	//
	//filter := new(ssofilter.LoginFilter)
	//filter.LoginWhiteList = &[]string{"/api/sso/user/login", "/api/sso/user/regist", "/api/sso/user/checkOrInValidateTokenString"}
	//filter.PrefixWhiteList = &[]string{"/api/iwork"}
	//filter.LoginUrl = ctx.Input.URL()
	//filter.Ctx = ctx
	//filter.SsoAddress = beego.AppConfig.String("isoft.sso.web.addr")
	//filter.ErrorFunc = func() {
	//	filter.Ctx.ResponseWriter.WriteHeader(401)
	//}
	//filter.Filter()
}

func main() {
	beego.InsertFilter("/api/iwork/httpservice/*", beego.BeforeExec, filterFunc)

	iworkpool.LoadAndCachePool()
	regist.RegistNodes()
	task.RegisterCronTask()
	task.InitialIWorkGlobalVar()		// 初始化全局变量

	fileServerPath := beego.AppConfig.String("file.server")
	beego.SetStaticPath("/api/files", fileServerPath)

	beego.Run()
}
