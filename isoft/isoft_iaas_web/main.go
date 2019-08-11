package main

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/imodules/misso"
	_ "isoft/isoft_iaas_web/routers"
	"isoft/isoft_iaas_web/startup/db"
	"isoft/isoft_iaas_web/startup/logger"
)

func init() {
	logger.InitLog()
	db.InitDb()

}

func main() {
	misso.RegisterISSOFilter()
	beego.Run()
}
