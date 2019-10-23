package controllers

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup/memory"
	"time"
)

func (this *WorkController) SaveFilters() {
	filter_id, _ := this.GetInt64("filter_id", -1)
	filterWork, _ := models.QueryWorkById(filter_id, orm.NewOrm())
	filters := make([]*models.Filters, 0)
	filters = this.appendSimpleWorkName(filters, filterWork)
	filters = this.appendComplexWorkName(filters, filterWork)
	_, err := models.InsertMultiFilters(filter_id, filters)
	if err == nil {
		flushMemoryFilter()
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) appendSimpleWorkName(filters []*models.Filters, filterWork models.Work) []*models.Filters {
	// workNames 以逗号分隔
	workNames := this.GetString("workNames")
	filters = append(filters, &models.Filters{
		FilterWorkId:    filterWork.Id,
		FilterWorkName:  filterWork.WorkName,
		WorkName:        workNames,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})
	return filters
}

func (this *WorkController) appendComplexWorkName(filters []*models.Filters, filterWork models.Work) []*models.Filters {
	complex_work_name := this.GetString("complexWorkName")
	filters = append(filters, &models.Filters{
		FilterWorkId:    filterWork.Id,
		FilterWorkName:  filterWork.WorkName,
		ComplexWorkName: complex_work_name,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})
	return filters
}

func flushMemoryFilter() {
	memory.FlushMemoryFilter()
}
