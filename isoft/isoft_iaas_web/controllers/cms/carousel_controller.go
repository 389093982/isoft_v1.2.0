package cms

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/cms"
	"time"
)

func (this *CMSController) FilterCarouselByPlacement() {
	placement := this.GetString("placement")
	carousels, err := cms.FilterCarouselByPlacement(placement)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "carousels": &carousels}
	}
	this.ServeJSON()
}

func (this *CMSController) FilterCarousels() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search := this.GetString("search")
	if search != "" {
		condArr["search"] = search
	}
	carousels, count, err := cms.FilterCarousels(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "carousels": &carousels, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}

// 增加通用连接地址
func (this *CMSController) AddCarousel() {
	placement := this.GetString("placement")
	title := this.GetString("title")
	content := this.GetString("content")
	imgpath := this.GetString("imgpath")
	linked_refer := this.GetString("linked_refer")
	carousel := &cms.Carousel{
		Placement:       placement,
		Title:           title,
		Content:         content,
		ImgPath:         imgpath,
		LinkedRefer:     linked_refer,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	_, err := cms.AddCarousel(carousel)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}
