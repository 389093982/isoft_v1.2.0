package iblog

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/iblog"
)

func (this *BlogController) BlogList2() {
	condArr := make(map[string]interface{})
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	catalog_id, _ := this.GetInt64("catalog_id", -1)
	search_text := this.GetString("search_text")
	// personal="personal"表示查询自己的博文,否则就查询热门博文
	personal := this.GetString("personal")
	if personal == "personal" {
		condArr["Author"] = this.Ctx.Input.Session("UserName").(string)
	} else {
		// 满足热门博文的条件,默认按照浏览次数排行
		condArr["querysOrder"] = "-Views"
		// 默认查询已发布的博文
		condArr["BlogStatus"] = 1
	}
	if catalog_id > 0 {
		condArr["catalog_id"] = catalog_id
	}
	if search_text != "" {
		condArr["search_text"] = search_text
	}
	blogs, count, err := iblog.QueryBlog(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "blogs": &blogs, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}
