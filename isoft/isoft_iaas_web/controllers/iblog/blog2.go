package iblog

import (
	"encoding/json"
	"io/ioutil"
	"isoft/isoft/common/httputil"
	"net/http"
)

func (this *BlogController) BlogList2() {
	var err error
	url := "http://localhost:8086/api/iwork/httpservice/BlogList2"
	paramMap := map[string]interface{}{}
	headerMap := map[string]interface{}{}
	parseFunc := func(resp *http.Response) {
		if resp.StatusCode == 200 {
			responsebody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				responseMap := map[string]interface{}{}
				err = json.Unmarshal(responsebody, &responseMap)
				if err == nil {
					blogs := responseMap["result"].(map[string]interface{})["blogs"]
					paginator := responseMap["result"].(map[string]interface{})["paginator"]
					this.Data["json"] = &map[string]interface{}{
						"status":    "SUCCESS",
						"blogs":     &blogs,
						"paginator": &paginator,
					}
				}
			}
		}
	}
	httputil.DoHttpRequestWithParserFunc(url, "post", paramMap, headerMap, parseFunc)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
