package misso

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/models/sso"
)

func RegisterModel() {
	orm.RegisterModel(new(sso.User))
	orm.RegisterModel(new(sso.AppRegister))
	orm.RegisterModel(new(sso.LoginRecord))
	orm.RegisterModel(new(sso.UserToken))
}
