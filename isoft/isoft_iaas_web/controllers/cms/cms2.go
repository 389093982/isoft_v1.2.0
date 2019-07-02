package cms

import (
	"isoft/isoft/common/httputil"
)

type AddPlacementResult struct {
	Status   string `json:"status"`
	Cost     int64  `json:"cost_ms"`
	ErrorMsg string `json:"errorMsg"`
}

func (this *CMSController) AddPlacement2() {
	placement_name := this.GetString("placement_name")
	placement_desc := this.GetString("placement_desc")

	paramMap := map[string]interface{}{
		"placement_name": placement_name,
		"placement_desc": placement_desc,
	}
	url := "http://localhost:8086/api/iwork/httpservice/AddPlacement2"
	headerMap := map[string]interface{}{}
	result := new(AddPlacementResult)
	err := httputil.DoHttpRequestAndParseToObj(url, "post", paramMap, headerMap, &result)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":   result.Status,
			"errorMsg": result.ErrorMsg,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
