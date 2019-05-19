package task

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func RegisterCronTask() {
	startILearningCronTask()

	startIWorkCronTask()
}

func InitialIWorkGlobalVar() {
	for _, name := range []string{"env_name", "env_address"} {
		if _, err := iwork.QueryGlobalVarByName(name); err != nil {
			gv := &iwork.GlobalVar{
				Name:            name,
				Value:           "",
				Type:            0,
				CreatedBy:       "SYSTEM",
				CreatedTime:     time.Now(),
				LastUpdatedBy:   "SYSTEM",
				LastUpdatedTime: time.Now(),
			}
			iwork.InsertOrUpdateGlobalVar(gv, orm.NewOrm())
		}
	}
}
