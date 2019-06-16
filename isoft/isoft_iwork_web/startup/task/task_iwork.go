package task

import (
	"fmt"
	"github.com/robfig/cron"
	"isoft/isoft/common/httputil"
	"isoft/isoft_iwork_web/models"
)

func startIWorkCronTask() {
	if metas, err := models.QueryAllCronMeta(true); err == nil {
		c := cron.New()
		for _, meta := range metas {
			c.AddJob(meta.CronStr, &iworkJob{meta: &meta})
		}
		c.Start()
	}
}

type iworkJob struct {
	meta *models.CronMeta
}

func (this *iworkJob) Run() {
	paramMap := make(map[string]interface{}, 0)
	headerMap := make(map[string]interface{}, 0)
	url := fmt.Sprintf("http://localhost:8086/api/iwork/httpservice/%s", this.meta.TaskName)
	httputil.DoHttpRequest(url, "post", paramMap, headerMap)
}
