package param

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type ParamVauleParser struct {
	ParamValue string
}

func (this *ParamVauleParser) CheckParamValueFormat() bool {
	if strings.HasPrefix(this.ParamValue, "$") && !strings.Contains(this.ParamValue, ".") {
		this.ParamValue = this.ParamValue + ".__output__" // 直接饮用某个节点的输出
	}
	this.removeUnsupportChars()
	if strings.HasPrefix(this.ParamValue, "$") && strings.Contains(this.ParamValue, ".") {
		return true
	}
	return false
}

func (this *ParamVauleParser) GetNodeNameFromParamValue() string {
	if this.CheckParamValueFormat() {
		return this.ParamValue[1:strings.Index(this.ParamValue, ".")]
	}
	return this.ParamValue
}

func (this *ParamVauleParser) GetParamNameFromParamValue() string {
	if this.CheckParamValueFormat() {
		return this.ParamValue[strings.Index(this.ParamValue, ".")+1:]
	}
	return this.ParamValue
}

// 去除不合理的字符
func (this *ParamVauleParser) removeUnsupportChars() {
	this.ParamValue = strings.TrimSpace(this.ParamValue)
	this.ParamValue = strings.Replace(this.ParamValue, "\n", "", -1)
}

func (this *ParamVauleParser) GetStaticParamValue() interface{} {
	this.removeUnsupportChars()
	if strings.HasPrefix(this.ParamValue, "$RESOURCE.") {
		resource_name := strings.TrimSpace(this.ParamValue)
		resource_name = strings.Replace(resource_name, "$RESOURCE.", "", -1)
		resource_name = strings.Replace(resource_name, ";", "", -1)
		resource_name = strings.TrimSpace(resource_name)
		resource, err := iwork.QueryResourceByName(resource_name)
		if err == nil {
			if resource.ResourceType == "db" {
				return resource.ResourceDsn
			} else if resource.ResourceType == "sftp" || resource.ResourceType == "ssh" {
				return resource
			}
		}
		return ""
	} else if strings.HasPrefix(this.ParamValue, "$WorkVars.") {
		return strings.Replace(strings.Replace(this.ParamValue, "$WorkVars.", "", -1), ";", "", -1)
	}
	return this.ParamValue
}

type ParamNameParser struct {
	ParamName string
	Step      *iwork.WorkStep
}

// 根据 ParamName 获取相对值,真值可能需要 ParamVauleParser 处理一下
func (this *ParamNameParser) ParseAndGetRelativeParamValue() string {
	var paramInputSchema iworkmodels.ParamInputSchema
	if err := json.Unmarshal([]byte(this.Step.WorkStepInput), &paramInputSchema); err != nil {
		return ""
	}
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == this.ParamName {
			// 非必须参数不得为空
			if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == "" {
				return ""
			}
			return item.ParamValue
		}
	}
	return ""
}

//根据步骤和参数名称获取静态参数值
func GetStaticParamValueWithStep(paramName string, step *iwork.WorkStep) interface{} {
	paramNameParser := &ParamNameParser{
		ParamName: paramName,
		Step:      step,
	}
	return GetStaticParamValue(paramNameParser.ParseAndGetRelativeParamValue())
}

func GetStaticParamValue(paramValue string) interface{} {
	paramValueParser := &ParamVauleParser{
		ParamValue: paramValue,
	}
	return paramValueParser.GetStaticParamValue()
}
