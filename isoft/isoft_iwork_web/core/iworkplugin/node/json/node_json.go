package json

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type JsonRenderNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *JsonRenderNode) Execute(trackingId string) {
	json_object := this.TmpDataMap[iworkconst.COMPLEX_PREFIX+"json_data"].([]map[string]interface{})
	bytes, err := json.Marshal(json_object)
	if err == nil {
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "json_data": string(bytes)})
	}
}

func (this *JsonRenderNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.COMPLEX_PREFIX + "json_data", "需要传入json对象"},
	}
	return this.BPIS1(paramMap)
}

func (this *JsonRenderNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "json_data"})
}

type JsonParserNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *JsonParserNode) Execute(trackingId string) {
	paramMap := make(map[string]interface{}, 0)
	json_str := this.TmpDataMap[iworkconst.STRING_PREFIX+"json_data"].(string)
	json_objects := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(json_str), &json_objects)
	if err == nil {
		paramMap["rows"] = json_objects
		for index, json_object := range json_objects {
			for paramName, paramValue := range json_object {
				paramMap[fmt.Sprintf("rows[%d].%s", index, paramName)] = paramValue
				if index == 0 {
					paramMap[fmt.Sprintf("rows.%s", paramName)] = paramValue
				}
			}
		}
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *JsonParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "json_data", "需要转换成json对象的字符串"},
		2: {"json_fields", "json对象的字段列表"},
	}
	return this.BPIS1(paramMap)
}

func (this *JsonParserNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	if json_fields := param.GetStaticParamValueWithStep("json_fields", this.WorkStep).(string); strings.TrimSpace(json_fields) != "" {
		jsonArr := strings.Split(json_fields, ",")
		for _, paramName := range jsonArr {
			if _paramName := strings.TrimSpace(paramName); _paramName != "" {
				items = append(items, iworkmodels.ParamOutputSchemaItem{
					ParentPath: "rows",
					ParamName:  _paramName,
				})
			}
		}
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
