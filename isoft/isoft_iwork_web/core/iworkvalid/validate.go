package iworkvalid

import (
	"fmt"
	"isoft/isoft_iwork_web/core/interfaces"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/models"
	"strings"
)

// 对必须参数进行非空校验
func CheckEmpty(step *models.WorkStep, paramSchemaParser interfaces.IParamSchemaCacheParser) (checkResult []string) {
	defer func() {
		if err := recover(); err != nil {
			checkResult = append(checkResult, errorutil.ToError(err).Error())
		}
	}()
	if strings.TrimSpace(step.WorkStepName) == "" || strings.TrimSpace(step.WorkStepType) == "" {
		checkResult = append(checkResult, fmt.Sprintf("Empty workStepName or empty workStepType was found!"))
		return
	}
	paramInputSchema := paramSchemaParser.GetCacheParamInputSchema(step)
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// work_start 节点参数由调度者提供,不做非空校验
		if step.WorkStepType != "work_start" {
			_, result := CheckEmptyForItem(item)
			checkResult = append(checkResult, result...)
		}
	}
	return
}

// 对输入参数做非空校验
func CheckEmptyForItem(item iworkmodels.ParamInputSchemaItem) (ok bool, checkResult []string) {
	if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == "" {
		checkResult = append(checkResult, fmt.Sprintf("Empty paramValue for %s was found!", item.ParamName))
	}
	return len(checkResult) == 0, checkResult
}
