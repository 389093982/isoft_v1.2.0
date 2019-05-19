package sso

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/sso"
)

type LoginRecordController struct {
	beego.Controller
}

func (this *LoginRecordController) LoginRecordList() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	loginrecords, count, err := sso.QueryLoginRecord(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "loginRecords": loginrecords,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
