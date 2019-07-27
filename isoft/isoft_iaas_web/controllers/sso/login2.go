package sso

import (
	"isoft/isoft/common/httputil"
	"time"
)

func (this *LoginController) PostRegist2() {
	paramMap := map[string]interface{}{
		"username":        this.Input().Get("username"),
		"passwd":          this.Input().Get("passwd"),
		"CreatedBy":       "SYSTEM",
		"CreatedTime":     time.Now(),
		"LastUpdatedBy":   "SYSTEM",
		"LastUpdatedTime": time.Now(),
	}
	url := "http://localhost:8086/api/iwork/httpservice/PostRegist2"
	headerMap := map[string]interface{}{}
	result := make(map[string]interface{})
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		this.Data["json"] = &result
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
