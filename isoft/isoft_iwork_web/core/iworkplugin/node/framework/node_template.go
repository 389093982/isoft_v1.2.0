package framework

import (
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"strings"
	"text/template"
)

type TemplateNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

type stringWriter struct {
	s string
}

func (this *stringWriter) Write(p []byte) (n int, err error) {
	this.s += string(p)
	return len(p), nil
}

func parseToDataMap(template_varName string, template_varValue []interface{}) map[string]interface{} {
	dataMap := make(map[string]interface{}, 0)
	names := strings.Split(template_varName, ",")
	if len(names) != len(template_varValue) {
		panic("miss match length error for template_varName and template_varValue")
	}
	for index, name := range names {
		dataMap[strings.TrimSpace(name)] = template_varValue[index]
	}
	return dataMap
}

func (this *TemplateNode) convertTemplateVarValue() []interface{} {
	if values, ok := this.TmpDataMap[iworkconst.COMPLEX_PREFIX+"template_varValue"].([]interface{}); ok {
		return values
	} else if value, ok := this.TmpDataMap[iworkconst.COMPLEX_PREFIX+"template_varValue"].(interface{}); ok {
		return []interface{}{value}
	}
	panic("template_varValue is invalid parameter")
}

func (this *TemplateNode) Execute(trackingId string) {
	template_text := this.TmpDataMap[iworkconst.STRING_PREFIX+"template_text"].(string)
	template_varName := this.TmpDataMap[iworkconst.STRING_PREFIX+"template_varName"].(string)
	template_varValue := this.convertTemplateVarValue()
	dataMap := parseToDataMap(template_varName, template_varValue)
	tmpl, err := template.New("template").Parse(template_text)
	if err != nil {
		panic(errors.Wrap(err, "模板语法错误!"))
	}
	w := &stringWriter{}
	err = tmpl.ExecuteTemplate(w, "template", dataMap)
	if err != nil {
		panic(err)
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{
		iworkconst.STRING_PREFIX + "template_text": w.s,
	})
}

func (this *TemplateNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "template_text", "模板文字"},
		2: {iworkconst.STRING_PREFIX + "template_varName", "模板变量名称"},
		3: {iworkconst.COMPLEX_PREFIX + "template_varValue", "模板变量值"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *TemplateNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "template_text"})
}
