package main

import (
	"github.com/astaxie/beego"
	_ "isoft/isoft_iwork_web/routers"
	"isoft/isoft_iwork_web/startup/db"
	"isoft/isoft_iwork_web/startup/logger"
	"isoft/isoft_iwork_web/startup/task"
)

func init() {
	logger.InitLog()
	db.InitDb()

}

func main() {
	task.RegisterCronTask()
	task.InitialIWorkGlobalVar()		// 初始化全局变量
	beego.Run()
}
