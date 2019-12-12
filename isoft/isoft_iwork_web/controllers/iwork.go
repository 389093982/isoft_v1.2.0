package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkfunc"
	"isoft/isoft_iwork_web/core/iworkplugin/node/regist"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/service"
	"isoft/isoft_iwork_web/startup/runtimecfg"
	"path"
	"strconv"
	"sync"
	"time"
)

type WorkController struct {
	beego.Controller
	interfaces.IFileUpload
}

func (this *WorkController) BuildIWorkDL() {
	//dls := make([]*IWorkDL,0)
	//works := iwork.GetAllWorkInfo()
	//for _, work := range works{
	//	dl := &IWorkDL{}
	//	steps, _ := iwork.GetAllWorkStepInfo(work.Id)
	//	for _, step := range steps{
	//		if step.WorkStepType == "work_start"{
	//			dl.RequestInfo = step.WorkStepInput
	//		}
	//		if step.WorkStepType == "work_end"{
	//			dl.ResponseInfo = step.WorkStepOutput
	//		}
	//	}
	//	dls = append(dls, dl)
	//}
}

func (this *WorkController) GetRelativeWork() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.GetRelativeWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":      "SUCCESS",
			"parentWorks": result["parentWorks"],
			"subworks":    result["subworks"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetLastRunLogDetail() {
	tracking_id := this.GetString("tracking_id")
	runLogRecord, _ := models.QueryRunLogRecordWithTracking(tracking_id)
	runLogDetails, err := models.QueryLastRunLogDetail(tracking_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "runLogDetails": runLogDetails, "runLogRecord": runLogRecord}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageLogRecord() {
	work_id, _ := this.GetInt64("work_id")
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	logLevel := this.GetString("logLevel")
	serviceArgs := map[string]interface{}{"work_id": work_id, "logLevel": logLevel, "offset": offset, "current_page": current_page, "ctx": this.Ctx}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.FilterPageLogRecord); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":        "SUCCESS",
			"runLogRecords": result["runLogRecords"],
			"paginator":     result["paginator"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) RunWork() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if err := service.ExecuteWithTx(serviceArgs, service.RunWork); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWork() {
	// 将请求参数封装成 work
	var work models.Work
	work_id, err := this.GetInt64("work_id", -1)
	if err == nil && work_id > 0 {
		work.Id = work_id
	}
	work.WorkName = this.GetString("work_name")
	work.WorkDesc = this.GetString("work_desc")
	work.WorkType = this.GetString("work_type")
	work.ModuleName = this.GetString("module_name")
	work.CacheResult, _ = this.GetBool("cache_result", false)
	work.CreatedBy = "SYSTEM"
	work.CreatedTime = time.Now()
	work.LastUpdatedBy = "SYSTEM"
	work.LastUpdatedTime = time.Now()
	if stringutil.CheckIgnoreCaseContains(work.WorkName, iworkconst.FORBIDDEN_WORK_NAMES) {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "非法名称!"}
	} else {
		serviceArgs := map[string]interface{}{"work": work}
		if err := service.ExecuteWithTx(serviceArgs, service.EditWorkService); err == nil {
			work, _ := models.QueryWorkByName(work.WorkName, orm.NewOrm())
			flushCache(work.Id)
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageWorks() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr["search"] = this.GetString("search")
	condArr["search_work_type"] = this.GetString("search_work_type")
	condArr["search_module"] = this.GetString("search_module")
	serviceArgs := map[string]interface{}{"condArr": condArr, "offset": offset, "current_page": current_page, "ctx": this.Ctx}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.FilterPageWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":            "SUCCESS",
			"works":             result["works"],
			"paginator":         result["paginator"],
			"runLogRecordCount": GetRunLogRecordCount(result["works"].([]models.Work)),
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteOrCopyWorkById() {
	defer this.ServeJSON()
	var err error
	id, _ := this.GetInt64("id")
	operate := this.GetString("operate")
	work, _ := models.QueryWorkById(id, orm.NewOrm())
	if operate == "copy" {
		o := orm.NewOrm()
		steps, _ := models.QueryAllWorkStepByWorkName(work.WorkName, o)
		work.Id = 0
		work.WorkName = work.WorkName + "_copy"
		id, err = models.InsertOrUpdateWork(&work, o)
		for index, _ := range steps {
			steps[index].Id = 0
			steps[index].WorkId = id
		}
		_, err = o.InsertMulti(len(steps), steps)
	} else {
		serviceArgs := map[string]interface{}{"id": id}
		if err = service.ExecuteWithTx(serviceArgs, service.DeleteWorkByIdService); err == nil {
			flushOneWorkCache(id, work.WorkName)
		}
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
}

func flushOneWorkCache(work_id int64, work_name string) {
	works := make([]models.Work, 0)
	work, err := models.QueryWorkById(work_id, orm.NewOrm())
	if err != nil && errors.As(err, &orm.ErrNoRows) {
		iworkcache.DeleteWorkCache(work_id) // 不存在则删除
	} else {
		serviceArgs := map[string]interface{}{"work_id": work_id}
		if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.GetRelativeWorkService); err == nil {
			works = append(works, result["parentWorks"].([]models.Work)...)
			works = append(works, result["subworks"].([]models.Work)...)
		}
		works = append(works, work)
	}
	for _, work := range works {
		if err = iworkcache.UpdateWorkCache(work.Id); err != nil {
			break
		}
		if workCache, err := iworkcache.LoadWorkCache(work.Id); err == nil {
			go saveHistory(workCache)
		}
	}
}

func flushAllWorkCache() {
	works := models.QueryAllWorkInfo(orm.NewOrm())
	for _, work := range works {
		if err := iworkcache.UpdateWorkCache(work.Id); err != nil {
			break
		}
		if workCache, err := iworkcache.LoadWorkCache(work.Id); err == nil {
			go saveHistory(workCache)
		}
	}
}

func flushCache(workId ...int64) (err error) {
	if len(workId) > 0 {
		flushOneWorkCache(workId[0], "")
	} else {
		flushAllWorkCache()
	}
	return
}

func (this *WorkController) Download() {
	work_id := this.Ctx.Input.Param(":work_id")
	workId, _ := strconv.ParseInt(work_id, 10, 64)
	workCache, _ := iworkcache.GetWorkCache(workId)
	tofile := path.Join(runtimecfg.FileSavePath, stringutil.RandomUUID())
	fileutil.WriteFile(tofile, []byte(workCache.RenderToString()), false)
	defer fileutil.RemoveFileOrDirectory(tofile)
	this.Ctx.Output.Download(tofile, fmt.Sprintf(`%s.txt`, workCache.Work.WorkName))
}

func (this *WorkController) GetAllFiltersAndWorks() {
	filterWorks, _ := models.GetAllFilterWorks()
	modules, _ := models.QueryAllModules()
	works, _ := models.QueryAllWorks()
	filters, _ := models.QueryAllFilters()
	this.Data["json"] = &map[string]interface{}{
		"status":      "SUCCESS",
		"filterWorks": filterWorks,
		"filters":     filters,
		"modules":     modules,
		"works":       works,
	}
	this.ServeJSON()
}

func (this *WorkController) QueryWorkDetail() {
	workId, _ := this.GetInt64("work_id", -1)
	workName := this.GetString("work_name", "")
	var (
		work models.Work
		err  error
	)
	if workId > 0 {
		work, err = models.QueryWorkById(workId, orm.NewOrm())
	} else if workName != "" {
		work, err = models.QueryWorkByName(workName, orm.NewOrm())
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "work": work}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) GetMetaInfo() {
	meta := this.GetString("meta")
	resultMap := map[string]interface{}{"status": "SUCCESS"}
	if meta == "funcCallers" {
		resultMap["funcCallers"] = (&iworkfunc.IWorkFuncProxy{}).GetFuncCallers()
	} else if meta == "nodeMetas" {
		resultMap["nodeMetas"] = regist.GetNodeMeta()
	}
	this.Data["json"] = &resultMap
	this.ServeJSON()
}

func GetRunLogRecordCount(works []models.Work) interface{} {
	m := make(map[int64]map[string]int64)
	var sm = new(sync.Map)
	var wg sync.WaitGroup
	wg.Add(len(works))
	for _, work := range works {
		go func(work models.Work) {
			errorCount, _ := models.QueryRunLogRecordCount(work.Id, "ERROR")
			allCount, _ := models.QueryRunLogRecordCount(work.Id, "")
			// 此处使用 map 会有并发问题 fatal error: concurrent map writes,故使用 sync.Map
			sm.Store(work.Id, map[string]int64{"errorCount": errorCount, "allCount": allCount})
			wg.Done()
		}(work)
	}
	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		m[key.(int64)] = value.(map[string]int64)
		return true
	})

	return m
}

func (this *WorkController) SaveProject() {
	persistentToFile()
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *WorkController) ImportProject() {
	importProject()
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}
