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
	headers := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"headers?", this.WorkStep).(string)
	headerSlice := strings.Split(headers, ",")
	paramMap := make(map[string]interface{}, 0)
	for _, header := range headerSlice {
		headerVal := request.Header.Get(header)
		paramMap["header_"+header] = headerVal
	}
	cookies := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"cookies?", this.WorkStep).(string)
	cookieSlice := strings.Split(cookies, ",")
	for _, cookie := range cookieSlice {
		if cookieVal, err := request.Cookie(cookie); err == nil {
			paramMap["cookie_"+cookie] = cookieVal.Value
		}
	}
	paramMap["ip"] = httputil.GetClientIP(request)
	// 将数据数据存储到数据中心
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *HttpRequestParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "headers?", "解析的请求头参数,多个参数使用逗号分隔"},
		2: {iworkconst.STRING_PREFIX + "cookies?", "解析的 cookies 参数,多个参数使用逗号分隔"},
	}
	return this.BPIS1(paramMap)
}

func (this *HttpRequestParserNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{"ip"})
}

func (this *HttpRequestParserNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{}
	headers := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"headers?", this.WorkStep).(string)
	cookies := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"cookies?", this.WorkStep).(string)
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, parseToItems(headers, "header_")...)
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, parseToItems(cookies, "cookie_")...)
	return pos
}

func parseToItems(paramStr, paramPrefix string) []iworkmodels.ParamOutputSchemaItem {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	if paramStr != "" {
		paramSlice := strings.Split(paramStr, ",")
		for _, param := range paramSlice {
			items = append(items, iworkmodels.ParamOutputSchemaItem{
				ParamName: paramPrefix + param,
			})
		}
	}
	return items
}
