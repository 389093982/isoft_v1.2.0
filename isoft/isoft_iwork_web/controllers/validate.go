package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
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
	validateWorks(work_id) // 触发校验
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
	completeFlag := make(chan int) // 校验并记录日志完成标识
	trackingId := stringutil.RandomUUID()
	// 记录校验耗费时间
	defer recordCostTimeLog(trackingId, time.Now())
	// 待校验的所有 work 信息
	workMap := prepareValiateWorks(workId)
	// 记录日志
	recordValidateLogRecord(trackingId, workId)
	logCh := make(chan *models.ValidateLogDetail, 10) // 指定容量
	go func() {
		recordValidateLogDetails(logCh, trackingId, workMap) // 开协程保证读和写同时进行
		completeFlag <- 1
	}()
	var wg sync.WaitGroup
	for work, workSteps := range workMap {
		wg.Add(1)
		validateWork(&work, workSteps, logCh, &wg)
	}
	wg.Wait()
	// 所有 work 执行完成后关闭 logCh
	close(logCh)
	<-completeFlag

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
func validateWork(work *models.Work, steps []models.WorkStep, logCh chan *models.ValidateLogDetail, workWg *sync.WaitGroup) {
	defer workWg.Done()
	defer func(start time.Time) {
		logCh <- &models.ValidateLogDetail{
			WorkId: work.Id,
			Detail: fmt.Sprintf(`validate %s work cost %d ms!`, work.WorkName, time.Now().Sub(start).Nanoseconds()/1e6),
		}
	}(time.Now())

	// 验证流程必须以 work_start 开始,以 work_end 结束
	checkBeginAndEnd(steps, logCh, work)
	var wg sync.WaitGroup
	wg.Add(len(steps))
	for _, step := range steps {
		validateStep(&step, logCh, &wg)
	}
	wg.Wait()
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
func validateStep(step *models.WorkStep, logCh chan *models.ValidateLogDetail, stepWg *sync.WaitGroup) {
	defer stepWg.Done()
	defer func() {
		if err := recover(); err != nil {
			logCh <- parseToValidateLogDetail(step, err)
		}
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
	checkResult = make([]string, 0)
	checkResultCh := make(chan string, 10)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		// 校验 step 中的参数是否为空
		for _, s := range iworkvalid.CheckEmpty(step, &node.ParamSchemaParser{WorkStep: step, ParamSchemaParser: &node.WorkStepFactory{WorkStep: step}}) {
			checkResultCh <- s
		}
	}()
	go func() {
		defer wg.Done()
		checkVariableRelationShip(step, checkResultCh)
	}()
	go func() {
		defer wg.Done()
		// 定制化校验
		for _, s := range CheckCustom(step) {
			checkResultCh <- s
		}
	}()
	wg.Wait()
	close(checkResultCh)
	for s := range checkResultCh {
		checkResult = append(checkResult, s)
	}
	return
}

func CheckCustom(step *models.WorkStep) (checkResult []string) {
	defer func() {
		if err := recover(); err != nil {
			checkResult = append(checkResult, errorutil.ToError(err).Error())
		}
	}()
	factory := &node.WorkStepFactory{WorkStep: step}
	return factory.ValidateCustom()
}

// 校验变量的引用关系
func checkVariableRelationShip(step *models.WorkStep, checkResultCh chan string) {
	defer func() {
		if err := recover(); err != nil {
			checkResultCh <- errorutil.ToError(err).Error()
		}
	}()
	inputSchema := node.GetCacheParamInputSchema(step)
	for _, item := range inputSchema.ParamInputSchemaItems {
		checkVariableRelationShipDetail(item, step.WorkId, step.WorkStepId, checkResultCh)
	}
	return
}

func checkVariableRelationShipDetail(item iworkmodels.ParamInputSchemaItem, work_id, work_step_id int64, checkResultCh chan string) {
	// 根据正则找到关联的节点名和字段名
	refers := iworkutil.GetRelativeValueWithReg(item.ParamValue)
	if len(refers) > 0 {
		preStepNodeNames := iworkutil.GetAllPreStepNodeName(work_id, work_step_id)
		skipNodeNames := []string{"RESOURCE", "WORK", "Error"}
		for _, refer := range refers {
			referNodeName, referFiledName := parseReferNodeAndFiledName(refer)
			if stringutil.CheckContains(referNodeName, skipNodeNames) {
				continue
			}
			if referNodeName == "Global" {
				checkVariableRelationShipForGlobal(referFiledName, checkResultCh)
			} else {
				checkVariableRelationShipForNode(referNodeName, referFiledName, preStepNodeNames, checkResultCh, work_id)
			}
		}
	}
	return
}

func checkVariableRelationShipForGlobal(referFiledName string, checkResultCh chan string) {
	_, err := models.QueryGlobalVarByName(referFiledName)
	if err != nil && errors.As(err, &orm.ErrNoRows) {
		checkResultCh <- fmt.Sprintf("Invalid referFiledName relationship for %s was found!", referFiledName)
	}
}

func checkVariableRelationShipForNode(referNodeName, referFiledName string, preStepNodeNames []string, checkResultCh chan string, work_id int64) {
	// 判断节点名称是否有效
	if !stringutil.CheckContains(referNodeName, preStepNodeNames) {
		checkResultCh <- fmt.Sprintf("Invalid referNodeName relationship for %s was found!", referNodeName)
		return
	}
	// $NodeName;
	// 只引用节点名的空字段直接跳过
	if referFiledName == "" {
		return
	}
	checkVariableRelationShipForNodeDetail(work_id, referNodeName, referFiledName, checkResultCh)
}

func checkVariableRelationShipForNodeDetail(work_id int64, referNodeName, referFiledName string, checkResultCh chan string) {
	// 判断字段名称是否有效
	if step, err := models.QueryWorkStepByStepName(work_id, referNodeName, orm.NewOrm()); err == nil {
		outputSchema := node.GetCacheParamOutputSchema(&step)
		exist := false
		for _, item := range outputSchema.ParamOutputSchemaItems {
			if item.ParamName == referFiledName || item.ParentPath+"."+item.ParamName == referFiledName {
				exist = true
				break
			}
		}
		if !exist {
			checkResultCh <- fmt.Sprintf("Invalid referFiledName relationship for %s was found!", referFiledName)
		}
	}
}

func parseReferNodeAndFiledName(refer string) (referNodeName, referFiledName string) {
	if strings.Contains(refer, ".") {
		referNodeName = refer[1:strings.Index(refer, ".")]
		referFiledName = refer[strings.Index(refer, ".")+1:]
	} else {
		referNodeName = refer
		referFiledName = ""
	}
	return referNodeName, referFiledName
}

func recordValidateLogDetails(logCh chan *models.ValidateLogDetail, trackingId string, workMap map[models.Work][]models.WorkStep) {
	workCaches := make(map[string]*models.Work, 0)
	workStepCaches := make(map[string]*models.WorkStep, 0)
	for work, steps := range workMap {
		workCaches[string(work.Id)] = &work
		for _, step := range steps {
			workStepCaches[string(work.Id)+"_"+string(step.WorkStepId)] = &step
		}
	}
	details := make([]*models.ValidateLogDetail, 0)
	// 从 logCh 中循环读取校验不通过的信息,并将其写入日志表中去
	for log := range logCh {
		workCacheKey, workStepCacheKey := string(log.WorkId), string(log.WorkId)+"_"+string(log.WorkStepId)
		work, _ := workCaches[workCacheKey]
		step, _ := workStepCaches[workStepCacheKey]
		details = append(details, fillWorkValidateLogDetail(log, trackingId, work, step))
	}
	models.InsertMultiValidateLogDetail(details)
}

func fillWorkValidateLogDetail(log *models.ValidateLogDetail, trackingId string, work *models.Work, step *models.WorkStep) *models.ValidateLogDetail {
	log.TrackingId = trackingId
	log.WorkName = work.WorkName
	log.CreatedBy = "SYSTEM"
	log.LastUpdatedBy = "SYSTEM"
	log.CreatedTime = time.Now()
	log.LastUpdatedTime = time.Now()
	if step != nil { // 不一定含有 step
		log.WorkStepName = step.WorkStepName
	}
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
