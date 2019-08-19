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
			Method:           "EditSqlMigrate",
			Router:           `/api/iwork/editSqlMigrate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"] = append(beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"],
		beego.ControllerComments{
			Method:           "ExecuteMigrate",
			Router:           `/api/iwork/executeMigrate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"] = append(beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"],
		beego.ControllerComments{
			Method:           "FilterPageSqlMigrate",
			Router:           `/api/iwork/filterPageSqlMigrate`,
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

	beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"] = append(beego.GlobalControllerRouter["isoft/isoft_iwork_web/controllers:WorkController"],
		beego.ControllerComments{
			Method:           "GetSqlMigrateInfo",
			Router:           `/api/iwork/getSqlMigrateInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
