package http

import (
	"isoft/isoft/common/httputil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"net/http"
	"strings"
)

type HttpRequestParserNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *HttpRequestParserNode) Execute(trackingId string) {
	request := this.Dispatcher.TmpDataMap[iworkconst.HTTP_REQUEST_OBJECT].(*http.Request)
	headers := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"headers", this.WorkStep).(string)
	headerSlice := strings.Split(headers, ",")
	paramMap := make(map[string]interface{}, 0)
	for _, header := range headerSlice {
		headerVal := request.Header.Get(header)
		paramMap[header] = headerVal
	}
	paramMap["ip"] = httputil.GetClientIP(request)
	// 将数据数据存储到数据中心
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *HttpRequestParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "headers", "解析的请求头参数,多个参数使用逗号分隔"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *HttpRequestParserNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{"ip"})
}

func (this *HttpRequestParserNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{}
	headers := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"headers", this.WorkStep).(string)
	headerSlice := strings.Split(headers, ",")
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, header := range headerSlice {
		items = append(items, iworkmodels.ParamOutputSchemaItem{
			ParamName: header,
		})
	}
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, items...)
	return pos
}
