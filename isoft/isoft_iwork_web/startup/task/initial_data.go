package task

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/models"
	"time"
)

var (
	initialGlobalVarMap map[string]string
)

func init() {
	initialGlobalVarMap = make(map[string]string)
	initialGlobalVarMap["SUCCESS"] = "SUCCESS"
	initialGlobalVarMap["ERROR"] = "ERROR"
}

func InitialData() {
	for key, value := range initialGlobalVarMap {
		if _, err := models.QueryGlobalVarByName(key); err != nil {
			gv := &models.GlobalVar{
				Name:            key,
				Value:           value,
				Type:            0,
				CreatedBy:       "SYSTEM",
				CreatedTime:     time.Now(),
				LastUpdatedBy:   "SYSTEM",
				LastUpdatedTime: time.Now(),
			}
			models.InsertOrUpdateGlobalVar(gv, orm.NewOrm())
		}
	}
}
