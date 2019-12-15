package framework

import (
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/reflectutil"
	"isoft/isoft_iwork_web/models"
	"text/template"
)

/**
SELECT * FROM article where
	{{if eq .search "_all"}}
		1=1
	{{else if eq .search "_hot"}}
		1=1
	{{else if eq .search "_personal"}}
		created_by = :search
	{{else}}
		catalog_name = :search
	{{end}}
and book_id = -1 order by last_updated_time desc
*/
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

func (this *TemplateNode) getTemplateDataMap() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	if dataMap, ok := this.TmpDataMap[iworkconst.COMPLEX_PREFIX+"template_dataMap"].(interface{}); ok {
		keys, values := reflectutil.InterfaceToMap(dataMap)
		for index, key := range keys {
			m[key.Interface().(string)] = values[index].Interface()
		}
	} else {
		panic("template_dataMap is not a map type")
	}
	return m
}

func (this *TemplateNode) Execute(trackingId string) {
	template_text := this.TmpDataMap[iworkconst.STRING_PREFIX+"template_text"].(string)
	dataMap := this.getTemplateDataMap()
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
		iworkconst.STRING_PREFIX + "template_text":     w.s,
		iworkconst.COMPLEX_PREFIX + "template_dataMap": dataMap,
	})
}

func (this *TemplateNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "template_text", "模板文字"},
		2: {iworkconst.COMPLEX_PREFIX + "template_dataMap", "模板变量绑定数据"},
	}
	return this.BPIS1(paramMap)
}

func (this *TemplateNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "template_text", iworkconst.COMPLEX_PREFIX + "template_dataMap"})
}
