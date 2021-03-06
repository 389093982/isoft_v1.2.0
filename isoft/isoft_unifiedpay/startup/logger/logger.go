package logger

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	var logDir string
	runmode := beego.AppConfig.String("runmode")
	if runmode != "prod" {
		logDir = "../../isoft_unifiedpay_log"
	} else {
		logDir = "../../isoft_unifiedpay_log" //项目部署的同级目录下放置log
	}

	logs.EnableFuncCallDepth(true)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"`+logDir+`/isoft_unifiedpay.log","maxdays":90,"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.Async(1e3)
}
