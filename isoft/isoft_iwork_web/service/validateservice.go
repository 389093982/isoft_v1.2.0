package service

import "isoft/isoft_iwork_web/models"

func LoadValidateResultService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	details, err := models.QueryLastValidateLogDetail()
	if err != nil {
		return nil, err
	}
	result["details"] = details
	return
}
