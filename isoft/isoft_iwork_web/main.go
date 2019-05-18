package main

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iwork_web/imodules/misso"
	_ "isoft/isoft_iwork_web/routers"
	"isoft/isoft_iwork_web/startup/db"
	"isoft/isoft_iwork_web/startup/logger"
	"isoft/isoft_iwork_web/startup/task"
)

func init() {
	logger.ConfigureLogInfo()
	db.ConfigureDBInfo()

}

func main() {
	misso.RegisterISSOFilter()
	task.RegisterCronTask()
	task.InitialIWorkGlobalVar()		// 初始化全局变量
	beego.Run()
}
