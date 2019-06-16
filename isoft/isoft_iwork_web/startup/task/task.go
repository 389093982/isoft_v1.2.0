package task

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/models"
	"time"
)

func RegisterCronTask() {

	startIWorkCronTask()
}

func InitialIWorkGlobalVar() {
	for _, name := range []string{"env_name", "env_address"} {
		if _, err := models.QueryGlobalVarByName(name); err != nil {
			gv := &models.GlobalVar{
				Name:            name,
				Value:           "",
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
