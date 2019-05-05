package miwork

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/models/iwork"
)

func RegisterModel() {
	if imodules.CheckModule("iwork") {
		orm.RegisterModel(new(iwork.CronMeta))
		orm.RegisterModel(new(iwork.Resource))
		orm.RegisterModel(new(iwork.Work))
		orm.RegisterModel(new(iwork.WorkStep))
		orm.RegisterModel(new(iwork.RunLogRecord))
		orm.RegisterModel(new(iwork.RunLogDetail))
		orm.RegisterModel(new(iwork.Entity))
		orm.RegisterModel(new(iwork.ValidateLogRecord))
		orm.RegisterModel(new(iwork.ValidateLogDetail))
		orm.RegisterModel(new(iwork.WorkHistory))
		orm.RegisterModel(new(iwork.TableMigrate))
	}
}
