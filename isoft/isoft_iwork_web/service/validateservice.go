package service

import "isoft/isoft_iwork_web/models"

func LoadValidateResultService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	details, err := models.QueryLastValidateLogDetail(work_id)
	if err != nil {
		return nil, err
	}
	result["details"] = details
	return
}
