package httputil

import (
	"fmt"
	"net/http"
	"strings"
)

func DoHttpRequestWithParserFunc(url string, method string, paramMap map[string]interface{},
	headerMap map[string]interface{}, parseFunc func(resp *http.Response) error) (err error) {
	client := &http.Client{}
	req, err := http.NewRequest(checkMethod(method), url, GetParamReader(paramMap))
	if err != nil {
		return err
	}
	setHeaderParameter(req, headerMap)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = parseFunc(resp)
	return
}

func setHeaderParameter(req *http.Request, headerMap map[string]interface{}) {
	// 设置默认请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// 添加或者覆盖默认请求头
	for paramName, paramValue := range headerMap {
		req.Header.Set(getStandardName(paramName), paramValue.(string))
	}
}

func getStandardName(paramName string) string {
	standardNames := []string{"Content-Type"}
	for _, sn := range standardNames {
		_sn := strings.Replace(sn, "-", "", -1)
		_sn = strings.ToUpper(_sn)
		_paramName := strings.Replace(paramName, "-", "", -1)
		_paramName = strings.ToUpper(_paramName)
		if _sn == _paramName {
			return sn
		}
	}
	return paramName
}

func GetParamReader(paramMap map[string]interface{}) *strings.Reader {
	if paramMap == nil || len(paramMap) == 0 {
		return strings.NewReader("")
	}
	s := make([]string, 0)
	for k, v := range paramMap {
		s = append(s, fmt.Sprintf("%s=%v", k, v))
	}
	paramStr := strings.Join(s, "&")
	return strings.NewReader(paramStr)
}

func checkMethod(method string) string {
	method = strings.TrimSpace(method)
	if method == "" {
		return "GET"
	}
	defaultMethods := []string{"GET", "POST", "PUT", "DELETE"}
	for _, dm := range defaultMethods {
		if strings.ToUpper(method) == dm {
			return dm
		}
	}
	panic(fmt.Sprintf("unsupport method : %s", method))
}
