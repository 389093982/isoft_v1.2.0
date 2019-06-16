package main

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iwork_web/core/iworkpool"
	_ "isoft/isoft_iwork_web/routers"
	_ "isoft/isoft_iwork_web/startup/db"
	_ "isoft/isoft_iwork_web/startup/logger"
	"isoft/isoft_iwork_web/startup/task"
)

func main() {
	iworkpool.LoadAndCachePool()
	task.RegisterCronTask()
	task.InitialIWorkGlobalVar()		// 初始化全局变量
	beego.Run()
}
