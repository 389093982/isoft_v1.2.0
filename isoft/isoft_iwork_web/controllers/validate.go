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
	"time"
)

func (this *WorkController) LoadValidateResult() {
	if result, err := service.ExecuteResultServiceWithTx(map[string]interface{}{}, service.LoadValidateResultService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "details": result["details"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) ValidateWork() {
	// 传入 work_id 则只校验单个 work, 否则校验全部
	work_id, _ := this.GetInt64("work_id", -1)
	validateAll(work_id)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

// 统计操作所花费的时间方法
func recordCostTimeLog(trackingId string, start time.Time) {
	models.InsertValidateLogDetail(getValidateLogDetail(trackingId,
		fmt.Sprintf("validate complete! total cost %d ms!", time.Now().Sub(start).Nanoseconds()/1e6)))
}

func validateAll(work_id int64) {
	trackingId := stringutil.RandomUUID()
	// 记录校验耗费时间
	defer recordCostTimeLog(trackingId, time.Now())
	// 记录日志
	models.InsertValidateLogRecord(&models.ValidateLogRecord{
		TrackingId:      trackingId,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})
	logCh := make(chan *models.ValidateLogDetail)
	workChan := make(chan int)
	// 待校验的所有 work
	works := make([]models.Work, 0)
	if work_id > 0 {
		work, _ := models.QueryWorkById(work_id, orm.NewOrm())
		works = append(works, work)
	} else {
		works = models.QueryAllWorkInfo(orm.NewOrm())
	}

	for _, work := range works {
		go func(work models.Work) {
			validateWork(&work, logCh, workChan)
		}(work)
	}

	go func() {
		for i := 0; i < len(works); i++ {
			<-workChan
		}
		// 所有 work 执行完成后关闭 logCh
		close(logCh)
	}()

	// 从 logCh 中循环读取校验不通过的信息,并将其写入日志表中去
	for log := range logCh {
		work, _ := models.QueryWorkById(log.WorkId, orm.NewOrm())
		step, _ := models.QueryOneWorkStep(work.Id, log.WorkStepId, orm.NewOrm())
		log.TrackingId = trackingId
		log.WorkName = work.WorkName
		log.WorkStepName = step.WorkStepName
		log.CreatedBy = "SYSTEM"
		log.LastUpdatedBy = "SYSTEM"
		log.CreatedTime = time.Now()
		log.LastUpdatedTime = time.Now()
		models.InsertValidateLogDetail(log)
	}
}

// 校验单个 work
func validateWork(work *models.Work, logCh chan *models.ValidateLogDetail, workChan chan int) {
	stepChan := make(chan int)
	steps, _ := models.QueryAllWorkStepInfo(work.Id, orm.NewOrm())
	// 验证流程必须以 work_start 开始,以 work_end 结束
	checkBeginAndEnd(steps, logCh, work)

	for _, step := range steps {
		go func(step models.WorkStep) {
			validateStep(&step, logCh, stepChan)
		}(step)
	}

	for i := 0; i < len(steps); i++ {
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

// 校验单个 step,并将校验不通过的信息放入 logCh 中
func validateStep(step *models.WorkStep, logCh chan *models.ValidateLogDetail, stepChan chan int) {
	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(error); ok {
				logCh <- &models.ValidateLogDetail{
					WorkId:     step.WorkId,
					WorkStepId: step.WorkStepId,
					Detail:     string(errorutil.PanicTrace(4)),
				}
			} else if _err, ok := err.(string); ok {
				logCh <- &models.ValidateLogDetail{
					WorkId:     step.WorkId,
					WorkStepId: step.WorkStepId,
					Detail:     _err,
				}
			} else if _err, ok := err.(models.ValidateLogDetail); ok {
				logCh <- &_err
			}
		}
		// 每执行完一个 step 就往 stepChan 里面发送完成通知
		stepChan <- 1
	}()

	checkResults := CheckStep(step)
	for _, checkResult := range checkResults {
		logCh <- &models.ValidateLogDetail{
			WorkId:     step.WorkId,
			WorkStepId: step.WorkStepId,
			Detail:     checkResult,
		}
	}
}

func CheckStep(step *models.WorkStep) (checkResult []string) {
	// 校验 step 中的参数是否为空
	checkResults1 := iworkvalid.CheckEmpty(step, &node.WorkStepFactory{WorkStep: step})
	checkResults2 := checkVariableRelationShip(step)
	// 定制化校验
	checkResults3 := CheckCustom(step)
	checkResult = append(checkResult, checkResults1...)
	checkResult = append(checkResult, checkResults2...)
	checkResult = append(checkResult, checkResults3...)
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

func getValidateLogDetail(trackingId, detail string) *models.ValidateLogDetail {
	return &models.ValidateLogDetail{
		TrackingId:      trackingId,
		Detail:          detail,
		CreatedBy:       "SYSTEM",
		LastUpdatedBy:   "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedTime: time.Now(),
	}
}
