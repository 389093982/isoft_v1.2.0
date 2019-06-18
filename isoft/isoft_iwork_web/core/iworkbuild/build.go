package iworkbuild

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/models"
	"strings"
	"time"
)

func getPisItems(step models.WorkStep) *iworkmodels.ParamInputSchema {
	parser := schema.WorkStepFactorySchemaParser{
		WorkStep:          &step,
		ParamSchemaParser: &node.WorkStepFactory{WorkStep: &step},
	}
	return parser.GetCacheParamInputSchema()
}

func BuildAutoCreateSubWork(step models.WorkStep, o orm.Ormer, insertStartEndWorkStepNodeFunc func(work_id int64, o orm.Ormer) error) {
	if step.WorkStepType != iworkconst.NODE_TYPE_WORK_SUB {
		return
	}
	paramInputSchema := getPisItems(step)
	for index, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.STRING_PREFIX+iworkconst.NODE_TYPE_WORK_SUB {
			paramValue := strings.TrimSpace(item.ParamValue)
			if !strings.HasPrefix(paramValue, "$WORK.") {
				// 修改值并同步到数据库
				paramInputSchema.ParamInputSchemaItems[index] = iworkmodels.ParamInputSchemaItem{
					ParamName:  item.ParamName,
					ParamValue: strings.Join([]string{"$WORK.", paramValue}, ""),
				}
				step.WorkStepInput = paramInputSchema.RenderToJson()
				// 自动创建子流程
				checkAndCreateSubWork(paramValue, o, insertStartEndWorkStepNodeFunc)
			}
			// 维护 work 的 WorkSubId 属性
			paramValue = iworkutil.GetSingleRelativeValueWithReg(paramValue) // 去除多余的 ; 等字符
			subWork, _ := models.QueryWorkByName(strings.Replace(paramValue, "$WORK.", "", -1), orm.NewOrm())
			step.WorkSubId = subWork.Id
			break
		}
	}
	models.InsertOrUpdateWorkStep(&step, o)
}

func checkAndCreateSubWork(work_name string, o orm.Ormer, insertStartEndWorkStepNodeFunc func(work_id int64, o orm.Ormer) error) {
	if _, err := models.QueryWorkByName(work_name, orm.NewOrm()); err != nil {
		// 不存在 work 则直接创建
		work := &models.Work{
			WorkName:        work_name,
			WorkDesc:        fmt.Sprintf("自动创建子流程:%s", work_name),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _, err := models.InsertOrUpdateWork(work, o); err == nil {
			// 写入 DB 并自动创建开始和结束节点
			insertStartEndWorkStepNodeFunc(work.Id, o)
		}
	}
}
