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
	"isoft/isoft_iwork_web/startup/memory"
	"isoft/isoft_iwork_web/startup/runtimecfg"
	_ "isoft/isoft_iwork_web/startup/sysconfig"
	"isoft/isoft_iwork_web/startup/task"
)

func main() {
	beego.InsertFilter("/api/iwork/httpservice/*", beego.BeforeExec, filter.FilterFunc)
	beego.InsertFilter("/*", beego.BeforeExec, filter.IPFilterFunc)

	memory.FlushAll()
	iworkpool.LoadAndCachePool()
	regist.RegistNodes()
	task.RegisterCronTask()

	task.InitialData() // 初始化数据

	beego.SetStaticPath("/api/files", runtimecfg.FileSavePath)

	beego.Run()
}
