package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"path"
	"time"
)

func saveHistory(wc *iworkcache.WorkCache) (err error) {
	work := wc.Work
	workHistory := wc.RenderToString()
	if err == nil {
		history := &models.WorkHistory{
			WorkId:          work.Id,
			WorkName:        work.WorkName,
			WorkDesc:        work.WorkDesc,
			WorkHistory:     string(workHistory),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		_, err = models.InsertOrUpdateWorkHistory(history)
	}
	fileServer := beego.AppConfig.String("work.cache.home")
	filename := path.Join(fileServer, work.WorkName+".work")
	fileutil.WriteFile(filename, []byte(workHistory), false)
	return
}

func (this *WorkController) FilterPageWorkHistory() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr := map[string]string{}
	histories, count, err := models.QueryWorkHistory(condArr, current_page, offset, orm.NewOrm())
	if err == nil {
		paginator := pagination.SetPaginator(this.Ctx, offset, count)
		this.Data["json"] = &map[string]interface{}{
			"status":        "SUCCESS",
			"workHistories": histories,
			"paginator":     pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums()),
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
