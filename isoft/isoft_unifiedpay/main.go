package main

import (
	"github.com/astaxie/beego"
	_ "isoft_unifiedpay/routers"
	_ "isoft_unifiedpay/startup/db"
	_ "isoft_unifiedpay/startup/globalSessions"
	_ "isoft_unifiedpay/startup/logger"
	_ "isoft_unifiedpay/startup/memory"
)

func main() {
	beego.Run()
}
