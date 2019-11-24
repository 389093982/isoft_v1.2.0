package controllers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/models"
	"time"
)

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

func (this *WorkController) RestoreFromWorkHistory() {
	defer this.ServeJSON()
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": errorutil.ToError(err).Error()}
		}
	}()
	id, _ := this.GetInt64("id", -1)
	workHistory, _ := models.QueryWorkHistoryById(id)
	restoreFromWorkHistoryToDB(workHistory.WorkHistory)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
}

func restoreFromWorkHistoryToDB(workHistory string) {
	workCache := iworkcache.WorkCache{}
	err := xml.Unmarshal([]byte(workHistory), &workCache)
	errorutil.CheckError(err)
	work := workCache.Work
	// id 置空, workName 重命名
	work.Id = 0
	work.WorkName = fmt.Sprintf(`%s_%s`, work.WorkName, time.Now().Format("20060102150405"))
	work.CreatedTime = time.Now()
	work.LastUpdatedTime = time.Now()
	id, err := models.InsertOrUpdateWork(&work, orm.NewOrm())
	errorutil.CheckError(err)
	for _, step := range workCache.Steps {
		// 重置 id, workId
		step.Id = 0
		step.WorkId = id
		step.CreatedTime = time.Now()
		step.LastUpdatedTime = time.Now()
		_, err := models.InsertOrUpdateWorkStep(&step, orm.NewOrm())
		errorutil.CheckError(err)
	}
}
