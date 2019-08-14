package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"] = append(beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"],
		beego.ControllerComments{
			Method:           "DeleteModuleById",
			Router:           `/api/iwork/deleteModuleById`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"] = append(beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"],
		beego.ControllerComments{
			Method:           "GetAllModules",
			Router:           `/api/iwork/getAllModules`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
