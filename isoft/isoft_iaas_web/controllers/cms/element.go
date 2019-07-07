package cms

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/cms"
	"time"
)

func (this *CMSController) FilterElementByPlacement() {
	placement := this.GetString("placement")
	elements, err := cms.FilterElementByPlacement(placement)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "elements": &elements}
	}
	this.ServeJSON()
}

func (this *CMSController) FilterElements() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search := this.GetString("search")
	if search != "" {
		condArr["search"] = search
	}
	elements, count, err := cms.FilterElements(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "elements": &elements, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}

// 增加通用连接地址
func (this *CMSController) AddElement() {
	placement := this.GetString("placement")
	title := this.GetString("title")
	content := this.GetString("content")
	imgpath := this.GetString("imgpath")
	linked_refer := this.GetString("linked_refer")
	element := &cms.Element{
		Placement:       placement,
		Title:           title,
		Content:         content,
		ImgPath:         imgpath,
		LinkedRefer:     linked_refer,
		Status:          0, // 默认停用
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	_, err := cms.AddElement(element)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *CMSController) UpdateElementStatus() {
	id, _ := this.GetInt64("id", -1)
	status, _ := this.GetInt("status", -1)
	var err error
	if status == 2 {
		err = cms.DeleteElement(id)
	} else {
		err = cms.UpdateElementStatus(id, status)
	}
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}
