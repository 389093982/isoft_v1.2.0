package controllers

import (
	"encoding/xml"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"io/ioutil"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"path"
	"time"
)

var fileServer = beego.AppConfig.String("work.cache.home")

func deleteHistory(workName string) {
	if workName != "" {
		filepath := path.Join(fileServer, workName+".work")
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
		filepath := path.Join(fileServer, work.WorkName+".work")
		fileutil.WriteFile(filepath, []byte(workHistory), false)
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

// @router /api/iwork/recover [get]
func (this *WorkController) Recover() {
	filepaths, _, _ := fileutil.GetAllSubFile(`D:\zhourui\program\go\goland_workspace\src\isoft\isoft_iwork_web\demo`)
	for _, filepath := range filepaths {
		recoverFile(filepath)
	}
	this.ServeJSON()
}

func recoverFile(filepath string) {
	workCache := iworkcache.WorkCache{}
	bytes, _ := ioutil.ReadFile(filepath)
	xml.Unmarshal(bytes, &workCache)
	work := workCache.Work
	work.CreatedTime = time.Now()
	work.LastUpdatedTime = time.Now()
	if _, err := models.QueryWorkByName(work.WorkName, orm.NewOrm()); err == nil {
		return
	}
	orm.NewOrm().Insert(&work)
	models.InsertOrUpdateWork(&work, orm.NewOrm())
	for _, step := range workCache.Steps {
		step.CreatedTime = time.Now()
		step.LastUpdatedTime = time.Now()
		orm.NewOrm().Insert(&step)
	}
}
