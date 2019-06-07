package controllers

import "encoding/json"

func (this *WorkController) ProxyCall() {
	methodName := this.GetString("methodName")
	this.Ctx.Redirect(301, "/api/iwork/"+methodName)
}

func (this *WorkController) GetMethodArgMap() map[string]interface{} {
	methodArgs := this.GetString("methodArgs")
	methodArgMap := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(methodArgs), &methodArgMap)
	return methodArgMap
}

func (this *WorkController) Formate() {
	var err error
	var formate_result string
	methodArgMap := this.GetMethodArgMap()
	if methodArgMap["formate_type"] == "sql" {
		// TODO
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "formate_result": formate_result}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
