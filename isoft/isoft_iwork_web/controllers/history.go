package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

func saveHistory(work_id int64) (err error) {
	work, _ := models.QueryWorkById(work_id, orm.NewOrm())
	steps, _ := models.QueryAllWorkStepInfo(work_id, orm.NewOrm())
	historyMap := make(map[string]interface{})
	historyMap["work"] = work
	historyMap["steps"] = steps
	workHistory, err := json.MarshalIndent(historyMap, "", "\t")
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
