package iblog

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"isoft/isoft/common/beegoutil"
	"isoft/isoft/common/httputil"
	"isoft/isoft_iaas_web/controllers/utils"
	"net/http"
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
	var err error
	url := "http://localhost:8086/api/iwork/httpservice/BlogList2"
	headerMap := map[string]interface{}{}
	parseFunc := func(resp *http.Response) error {
		if resp.StatusCode == 200 {
			responsebody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				result := new(BlogListResult)
				err = json.Unmarshal(responsebody, &result)
				if err == nil {
					blogs := result.Result.Blogs
					paginator := result.Result.Paginator
					this.Data["json"] = &map[string]interface{}{
						"status":    "SUCCESS",
						"blogs":     &blogs,
						"paginator": &paginator,
					}
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return errors.New("服务调用失败!")
		}
		return nil
	}
	err = httputil.DoHttpRequestWithParserFunc(url, "post", paramMap, headerMap, parseFunc)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
