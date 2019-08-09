package main

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iwork_web/core/iworkplugin/node/regist"
	"isoft/isoft_iwork_web/core/iworkpool"
	_ "isoft/isoft_iwork_web/routers"
	_ "isoft/isoft_iwork_web/startup/db"
	_ "isoft/isoft_iwork_web/startup/dipool"
	"isoft/isoft_iwork_web/startup/filter"
	_ "isoft/isoft_iwork_web/startup/logger"
	_ "isoft/isoft_iwork_web/startup/sysconfig"
	"isoft/isoft_iwork_web/startup/task"
	"log"
	"net/http"
	_ "net/http/pprof"
)



func main() {
	go func() {
		// http://127.0.0.1:10000/debug/pprof/
		log.Println(http.ListenAndServe("localhost:10000", nil))
	}()

	beego.InsertFilter("/api/iwork/httpservice/*", beego.BeforeExec, filter.FilterFunc)

	iworkpool.LoadAndCachePool()
	regist.RegistNodes()
	task.RegisterCronTask()
	task.InitialIWorkGlobalVar()		// 初始化全局变量

	fileServerPath := beego.AppConfig.String("file.server")
	beego.SetStaticPath("/api/files", fileServerPath)

	beego.Run()
}
