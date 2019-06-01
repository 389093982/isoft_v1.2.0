package iblog

import (
	"isoft/isoft/common/beegoutil"
	"isoft/isoft/common/httputil"
)

type PostCatalogEditResult struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

func (this *CatalogController) PostCatalogEdit2() {
	paramMap := map[string]interface{}{
		"user_name":    this.Ctx.Input.Session("UserName").(string),
		"catalog_id":   beegoutil.GetInt64(this, "catalog_id", -1),
		"catalog_name": this.GetString("catalog_name"),
		"catalog_desc": this.GetString("catalog_desc"),
	}
	url := "http://localhost:8086/api/iwork/httpservice/PostCatalogEdit2"
	headerMap := map[string]interface{}{}
	result := new(PostCatalogEditResult)
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
