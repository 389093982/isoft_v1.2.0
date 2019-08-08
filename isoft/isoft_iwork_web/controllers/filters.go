package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/models"
	"time"
)

func (this *WorkController) SaveFilters() {
	filter_id, _ := this.GetInt64("filter_id", -1)
	filter, _ := models.QueryWorkById(filter_id, orm.NewOrm())
	work_names := this.GetString("work_names")
	var workNames []string
	json.Unmarshal([]byte(work_names), &workNames)
	filters := make([]*models.Filters, 0)
	for _, work_name := range workNames {
		filters = append(filters, &models.Filters{
			FilterWorkId:    filter_id,
			FilterWorkName:  filter.WorkName,
			WorkName:        work_name,
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		})
	}
	_, err := models.InsertMultiFilters(filter_id, filters)
	if err == nil {
		iworkcache.ReloadAllWorkCache()
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
