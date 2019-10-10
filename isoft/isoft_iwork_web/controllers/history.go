package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"path"
	"time"
)

var iwork_persistent_path = beego.AppConfig.String("iwork_persistent_path")

func deleteHistory(workName string) {
	if workName != "" {
		filepath := path.Join(iwork_persistent_path, "works", workName+".work")
		fileutil.RemoveFileOrDirectory(filepath)
	}
}

func saveHistory(wc *iworkcache.WorkCache) (err error) {
	work := wc.Work
	workHistory := wc.RenderToString()
	hash := hashutil.CalculateHashWithString(work.WorkName + workHistory)
	if _, err := models.QueryWorkHistoryByHash(hash); err != nil && errors.As(err, &orm.ErrNoRows) {
		history := &models.WorkHistory{
			WorkId:          work.Id,
			WorkName:        work.WorkName,
			WorkDesc:        work.WorkDesc,
			WorkHistory:     workHistory,
			Hash:            hash,
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		_, err = models.InsertOrUpdateWorkHistory(history)
		filepath := path.Join(iwork_persistent_path, "works", work.WorkName+".work")
		fileutil.WriteFile(filepath, []byte(workHistory), false)
	}
	return
}

func (this *WorkController) FilterPageWorkHistory() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search := this.GetString("search")
	condArr := map[string]string{"search": search}
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
