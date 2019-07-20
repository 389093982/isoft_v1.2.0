package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/core/iworkvalid"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/service"
	"strings"
	"sync"
	"time"
)

func (this *WorkController) LoadValidateResult() {
	serviceArgs := make(map[string]interface{}, 0)
	work_id, _ := this.GetInt64("work_id", -1)
	serviceArgs["work_id"] = work_id
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.LoadValidateResultService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "details": result["details"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) ValidateWork() {
	workId, _ := this.GetInt64("work_id", -1)
	validateWorks(workId) // 校验全部或者只校验单个 workId
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func validateWorks(workId int64) {
	trackingId := stringutil.RandomUUID()
	// 记录校验耗费时间
	defer recordCostTimeLog(trackingId, time.Now())
	// 记录日志
	recordValidateLogRecord(trackingId, workId)
	// 待校验的所有 work 信息
	workMap := prepareValiateWorks(workId)
	logCh := make(chan *models.ValidateLogDetail)
	workChan := make(chan int)
	for work, workSteps := range workMap {
		go func(work models.Work, workSteps []models.WorkStep) {
			validateWork(&work, workSteps, logCh, workChan)
		}(work, workSteps)
	}
	go func() {
		for i := 0; i < len(workMap); i++ {
			<-workChan
		}
		// 所有 work 执行完成后关闭 logCh
		close(logCh)
	}()
	recordValidateLogDetails(logCh, trackingId)
}

func prepareValiateWorks(workId int64) map[models.Work][]models.WorkStep {
	o := orm.NewOrm()
	dataMap := make(map[models.Work][]models.WorkStep, 0)
	if workId > 0 {
		work, _ := models.QueryWorkById(workId, o)
		steps, _ := models.QueryAllWorkStepInfo(work.Id, o)
		dataMap[work] = steps
	} else {
		works := models.QueryAllWorkInfo(orm.NewOrm())
		for _, work := range works {
			steps, _ := models.QueryAllWorkStepInfo(work.Id, o)
			dataMap[work] = steps
		}
	}
	return dataMap
}

// 校验单个 work
func validateWork(work *models.Work, steps []models.WorkStep, logCh chan *models.ValidateLogDetail, workChan chan int) {
	stepChan := make(chan int)
	// 验证流程必须以 work_start 开始,以 work_end 结束
	checkBeginAndEnd(steps, logCh, work)

	for _, step := range steps {
		go func(step models.WorkStep) {
			validateStep(&step, logCh, stepChan)
		}(step)
	}

	for i := 0; i < len(steps); i++ { // 阻塞直至所有 step 执行完
		<-stepChan
	}
	// 所有 step 执行完成后就往 workChan 里面发送完成通知
	workChan <- 1
}

func checkBeginAndEnd(steps []models.WorkStep, logCh chan *models.ValidateLogDetail, work *models.Work) {
	checkError := func(step *models.WorkStep, workStepName string, errorMsg string) {
		if step.WorkStepType != workStepName {
			logCh <- &models.ValidateLogDetail{
				WorkId:     work.Id,
				WorkStepId: step.WorkStepId,
				Detail:     errorMsg,
			}
		}
	}
	checkError(&steps[0], "work_start", "work must start with a work_start node!")
	checkError(&steps[len(steps)-1], "work_end", "work must end with a work_end node!")
	return
}

func parseToValidateLogDetail(step *models.WorkStep, err interface{}) *models.ValidateLogDetail {
	var detail string
	if _, ok := err.(error); ok {
		detail = string(errorutil.PanicTrace(4))
	} else if _err, ok := err.(string); ok {
		detail = _err
	} else if _err, ok := err.(models.ValidateLogDetail); ok {
		detail = _err.Detail
	}
	return &models.ValidateLogDetail{
		WorkId:     step.WorkId,
		WorkStepId: step.WorkStepId,
		Detail:     detail,
	}
}

// 校验单个 step,并将校验不通过的信息放入 logCh 中
// 校验单个 step,并将校验不通过的信息放入 logCh 中
func validateStep(step *models.WorkStep, logCh chan *models.ValidateLogDetail, stepChan chan int) {
	defer func() {
		if err := recover(); err != nil {
			logCh <- parseToValidateLogDetail(step, err)
		}
		// 每执行完一个 step 就往 stepChan 里面发送完成通知
		stepChan <- 1
	}()

	for _, checkResult := range getCheckResultsForStep(step) {
		logCh <- &models.ValidateLogDetail{
			WorkId:     step.WorkId,
			WorkStepId: step.WorkStepId,
			Detail:     checkResult,
		}
	}
}

func getCheckResultsForStep(step *models.WorkStep) (checkResult []string) {
	checkResultCh := make(chan []string, 10)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		// 校验 step 中的参数是否为空
		checkResults1 := iworkvalid.CheckEmpty(step, &node.WorkStepFactory{WorkStep: step})
		checkResultCh <- checkResults1
	}()
	go func() {
		defer wg.Done()
		checkResults2 := checkVariableRelationShip(step)
		checkResultCh <- checkResults2
	}()
	go func() {
		defer wg.Done()
		// 定制化校验
		checkResults3 := CheckCustom(step)
		checkResultCh <- checkResults3
	}()
	wg.Wait()
	close(checkResultCh)
	for _checkResult := range checkResultCh {
		checkResult = append(checkResult, _checkResult...)
	}
	return
}

func CheckCustom(step *models.WorkStep) (checkResult []string) {
	factory := &node.WorkStepFactory{WorkStep: step}
	return factory.ValidateCustom()
}

// 校验变量的引用关系
func checkVariableRelationShip(step *models.WorkStep) (checkResult []string) {
	parser := schema.WorkStepFactoryParamSchemaParser{WorkStep: step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: step}}
	inputSchema := parser.GetCacheParamInputSchema()
	for _, item := range inputSchema.ParamInputSchemaItems {
		result := checkVariableRelationShipDetail(item, step.WorkId, step.WorkStepId)
		checkResult = append(checkResult, result...)
	}
	return
}

func checkVariableRelationShipDetail(item iworkmodels.ParamInputSchemaItem, work_id, work_step_id int64) (checkResult []string) {
	// 根据正则找到关联的节点名和字段名
	refers := iworkutil.GetRelativeValueWithReg(item.ParamValue)
	if len(refers) == 0 {
		return
	}
	preStepNodeNames := iworkutil.GetAllPreStepNodeName(work_id, work_step_id)
	skipNodeNames := []string{"RESOURCE", "WORK"}
	for _, refer := range refers {
		referNodeName := refer[1:strings.Index(refer, ".")]
		referFileName := refer[strings.Index(refer, ".")+1:]
		// 非节点类型直接跳过
		if stringutil.CheckContains(referNodeName, skipNodeNames) {
			break
		}
		// 判断节点名称是否有效
		if !stringutil.CheckContains(referNodeName, preStepNodeNames) {
			checkResult = append(checkResult, fmt.Sprintf("Invalid referNodeName relationship for %s was found!", referNodeName))
			continue
		}

		// 判断字段名称是否有效
		if step, err := models.QueryWorkStepByStepName(work_id, referNodeName, orm.NewOrm()); err == nil {
			parser := schema.WorkStepFactoryParamSchemaParser{WorkStep: &step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step}}
			outputSchema := parser.GetCacheParamOutputSchema()
			exist := false
			for _, item := range outputSchema.ParamOutputSchemaItems {
				if item.ParamName == referFileName {
					exist = true
					break
				}
			}
			if !exist {
				checkResult = append(checkResult, fmt.Sprintf("Invalid referFileName relationship for %s was found!", referFileName))
			}
		}
	}
	return
}

func recordValidateLogDetails(logCh chan *models.ValidateLogDetail, trackingId string) {
	workCaches := make(map[string]*models.Work, 0)
	workStepCaches := make(map[string]*models.WorkStep, 0)
	details := make([]*models.ValidateLogDetail, 0)
	// 从 logCh 中循环读取校验不通过的信息,并将其写入日志表中去
	for log := range logCh {
		_workCacheKey, _workStepCacheKey := string(log.WorkId), string(log.WorkId)+"_"+string(log.WorkStepId)
		if _, ok := workCaches[_workCacheKey]; !ok {
			_work, _ := models.QueryWorkById(log.WorkId, orm.NewOrm())
			workCaches[_workCacheKey] = &_work
		}
		if _, ok := workCaches[_workStepCacheKey]; !ok {
			_step, _ := models.QueryOneWorkStep(log.WorkId, log.WorkStepId, orm.NewOrm())
			workStepCaches[_workStepCacheKey] = &_step
		}
		work, _ := workCaches[_workCacheKey]
		step, _ := workStepCaches[_workStepCacheKey]
		details = append(details, fillWorkValidateLogDetail(log, trackingId, work, step))
	}
	models.InsertMultiValidateLogDetail(details)
}

func fillWorkValidateLogDetail(log *models.ValidateLogDetail, trackingId string, work *models.Work, step *models.WorkStep) *models.ValidateLogDetail {
	log.TrackingId = trackingId
	log.WorkName = work.WorkName
	log.WorkStepName = step.WorkStepName
	log.CreatedBy = "SYSTEM"
	log.LastUpdatedBy = "SYSTEM"
	log.CreatedTime = time.Now()
	log.LastUpdatedTime = time.Now()
	return log
}

func recordValidateLogRecord(trackingId string, workId int64) (int64, error) {
	return models.InsertValidateLogRecord(&models.ValidateLogRecord{
		TrackingId:      trackingId,
		WorkId:          workId,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})
}

// 统计操作所花费的时间方法
func recordCostTimeLog(trackingId string, start time.Time) {
	detail := fmt.Sprintf("validate complete! total cost %d ms!", time.Now().Sub(start).Nanoseconds()/1e6)
	log := newValidateLogDetail(trackingId, detail)
	models.InsertMultiValidateLogDetail([]*models.ValidateLogDetail{log})
}

func newValidateLogDetail(trackingId, detail string) *models.ValidateLogDetail {
	return &models.ValidateLogDetail{
		TrackingId:      trackingId,
		Detail:          detail,
		CreatedBy:       "SYSTEM",
		LastUpdatedBy:   "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedTime: time.Now(),
	}
}
