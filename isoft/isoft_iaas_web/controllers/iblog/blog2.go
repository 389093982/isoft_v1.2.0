package iblog

import (
	"isoft/isoft/common/beegoutil"
	"isoft/isoft/common/httputil"
)

type BlogListResult struct {
	Status string `json:"status"`
	Cost   int64  `json:"cost_ms"`
	Result struct {
		Blogs     interface{} `json:"blogs"`
		Paginator interface{} `json:"paginator"`
	}
}

func (this *BlogController) BlogList2() {
	paramMap := map[string]interface{}{
		"offset":       beegoutil.GetInt64(this, "offset", 10),      // 每页记录数
		"current_page": beegoutil.GetInt64(this, "current_page", 1), // 当前页
		"catalog_id":   beegoutil.GetInt64(this, "catalog_id", -1),
		"search_text":  this.GetString("search_text"),
	}
	url := "http://localhost:8086/api/iwork/httpservice/BlogList2"
	headerMap := map[string]interface{}{}
	result := new(BlogListResult)
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		blogs := result.Result.Blogs
		paginator := result.Result.Paginator
		this.Data["json"] = &map[string]interface{}{
			"status":    "SUCCESS",
			"blogs":     &blogs,
			"paginator": &paginator,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
