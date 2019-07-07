package share

import (
	"isoft/isoft/common/beegoutil"
	"isoft/isoft/common/httputil"
)

type ShowShareDetailResult struct {
	Status string `json:"status"`
	Cost   int64  `json:"cost_ms"`
	Result struct {
		ErrorMsg string      `json:"errorMsg"`
		Share    interface{} `json:"share"`
	}
}

func (this *ShareController) ShowShareDetail2() {
	paramMap := map[string]interface{}{
		"share_id": beegoutil.GetInt64(this, "share_id", -1),
	}
	url := "http://localhost:8086/api/iwork/httpservice/ShowShareDetail2"
	headerMap := map[string]interface{}{}
	result := new(ShowShareDetailResult)
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "share": result.Result.Share}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

type FilterShareListResult struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

func (this *ShareController) FilterShareList2() {
	paramMap := map[string]interface{}{
		"userName":     this.GetSession("UserName").(string),
		"current_page": beegoutil.GetInt64(this, "current_page", 1),
		"offset":       beegoutil.GetInt64(this, "offset", 10),
		"search_type":  this.GetString("search_type"),
	}
	url := "http://localhost:8086/api/iwork/httpservice/PostCatalogEdit2"
	headerMap := map[string]interface{}{}
	result := new(FilterShareListResult)
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

type AddNewShareResult struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

func (this *ShareController) AddNewShare2() {
	paramMap := map[string]interface{}{
		"share_type": this.GetString("share_type"),
		"share_desc": this.GetString("share_desc"),
		"link_href":  this.GetString("link_href"),
		"content":    this.GetString("content"),
		"userName":   this.GetSession("UserName").(string),
	}
	url := "http://localhost:8086/api/iwork/httpservice/AddNewShare2"
	headerMap := map[string]interface{}{}
	result := new(AddNewShareResult)
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
