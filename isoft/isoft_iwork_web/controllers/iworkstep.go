package controllers

import (
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkfunc"
	"isoft/isoft_iwork_web/service"
)

func (this *WorkController) AddWorkStep() {
	serviceArgs := make(map[string]interface{}, 0)
	work_id, _ := this.GetInt64("work_id")
	serviceArgs["work_id"] = work_id
	serviceArgs["work_step_id"], _ = this.GetInt64("work_step_id")
	serviceArgs["work_step_type"] = this.GetString("work_step_type") // 需要创建的节点类型
	if err := service.ExecuteWithTx(serviceArgs, service.AddWorkStepService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) EditWorkStepBaseInfo() {
	work_id, _ := this.GetInt64("work_id", -1)
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	serviceArgs := map[string]interface{}{
		"work_id":        work_id,
		"work_step_id":   work_step_id,
		"work_step_name": this.GetString("work_step_name"),
		"work_step_type": this.GetString("work_step_type"),
		"work_step_desc": this.GetString("work_step_desc"),
		"is_defer":       this.GetString("is_defer"),
	}
	if err := service.ExecuteWithTx(serviceArgs, service.EditWorkStepBaseInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) WorkStepList() {
	jsonMap := make(map[string]interface{})
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if workCache, err := iworkcache.GetWorkCache(work_id); err == nil {
		jsonMap["usedMap"] = workCache.Usage.UsedMap
	}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.WorkStepListService); err == nil {
		jsonMap["status"] = "SUCCESS"
		jsonMap["worksteps"] = result["worksteps"]
	} else {
		jsonMap["status"] = "ERROR"
		jsonMap["errorMsg"] = err.Error()
	}
	this.Data["json"] = jsonMap
	this.ServeJSON()
}

func (this *WorkController) CopyWorkStepByWorkStepId() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if err := service.ExecuteWithTx(serviceArgs, service.CopyWorkStepByWorkStepIdService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkStepByWorkStepId() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if err := service.ExecuteWithTx(serviceArgs, service.DeleteWorkStepByWorkStepIdService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) LoadWorkStepInfo() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.LoadWorkStepInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":                    "SUCCESS",
			"step":                      result["step"],
			"paramInputSchema":          result["paramInputSchema"],
			"paramOutputSchema":         result["paramOutputSchema"],
			"paramOutputSchemaTreeNode": result["paramOutputSchemaTreeNode"],
			"paramMappings":             result["paramMappings"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetAllWorkStepInfo() {
	var jsonMap = make(map[string]interface{})
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.GetAllWorkStepInfoService); err == nil {
		jsonMap["status"] = "SUCCESS"
		jsonMap["steps"] = result["steps"]
	} else {
		jsonMap["status"] = "ERROR"
		jsonMap["errorMsg"] = err.Error()
	}
	this.Data["json"] = jsonMap
	this.ServeJSON()
}

func (this *WorkController) ChangeWorkStepOrder() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	_type := this.GetString("type")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id, "_type": _type}
	if err := service.ExecuteWithTx(serviceArgs, service.ChangeWorkStepOrderService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) LoadPreNodeOutput() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, service.LoadPreNodeOutputService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "prePosTreeNodeArr": result["prePosTreeNodeArr"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) RefactorWorkStepInfo() {
	serviceArgs := make(map[string]interface{}, 0)
	work_id, _ := this.GetInt64("work_id")
	serviceArgs["work_id"] = work_id
	serviceArgs["refactor_worksub_name"] = this.GetString("refactor_worksub_name")
	serviceArgs["refactor_work_step_ids"] = this.GetString("selections")
	if err := service.ExecuteWithTx(serviceArgs, service.RefactorWorkStepInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) BatchChangeIndent() {
	serviceArgs := make(map[string]interface{}, 0)
	work_id, _ := this.GetInt64("work_id")
	serviceArgs["work_id"] = work_id
	serviceArgs["mod"] = this.GetString("mod")
	serviceArgs["indent_work_step_ids"] = this.GetString("selections")
	if err := service.ExecuteWithTx(serviceArgs, service.BatchChangeIndentService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) EditWorkStepParamInfo() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	paramInputSchemaStr := this.GetString("paramInputSchemaStr")
	paramMappingsStr := this.GetString("paramMappingsStr")
	serviceArgs := map[string]interface{}{
		"work_id":             work_id,
		"work_step_id":        work_step_id,
		"paramInputSchemaStr": paramInputSchemaStr,
		"paramMappingsStr":    paramMappingsStr,
	}
	if err := service.ExecuteWithTx(serviceArgs, service.EditWorkStepParamInfo); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	flushCache(work_id)
	this.ServeJSON()
}

func (this *WorkController) ParseToMultiValue() {
	pureText, _ := this.GetBool("pureText", false)
	value := this.GetString("value")
	if !pureText {
		multiVals, err := iworkfunc.SplitWithLexerAnalysis(value)
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "multiVals": multiVals}
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "multiVals": []string{value}}
	}

	this.ServeJSON()
}
