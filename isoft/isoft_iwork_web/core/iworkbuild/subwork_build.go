package iworkbuild

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/models"
	"strings"
	"time"
)

func BuildAutoCreateSubWork(step models.WorkStep, o orm.Ormer, insertStartEndWorkStepNodeFunc func(work_id int64, o orm.Ormer) error) {
	if step.WorkStepType != iworkconst.NODE_TYPE_WORK_SUB {
		return
	}
	paramInputSchema := node.GetCacheParamInputSchema(&step)
	for index, item := range paramInputSchema.ParamInputSchemaItems {
		// 参数名称代表 work_sub
		if item.ParamName == iworkconst.STRING_PREFIX+iworkconst.NODE_TYPE_WORK_SUB {
			// work_sub 名称支持纯文本和 $WORK.xxx 两种格式,统一转换成 $WORK.xxx 格式
			workSubNameRef := strings.TrimSpace(item.ParamValue)
			if !strings.HasPrefix(workSubNameRef, "$WORK.") {
				// 修改值并同步到数据库
				paramInputSchema.ParamInputSchemaItems[index] = iworkmodels.ParamInputSchemaItem{
					ParamName:  item.ParamName,
					ParamValue: strings.Join([]string{"$WORK.", workSubNameRef}, ""),
				}
				step.WorkStepInput = paramInputSchema.RenderToJson()
			} else {
				workSubNameRef = strings.TrimPrefix(workSubNameRef, "$WORK.")
				workSubNameRef = strings.TrimSpace(workSubNameRef)
				workSubNameRef = strings.TrimSuffix(workSubNameRef, ";")
			}
			// 自动创建子流程
			createOrUpdateSubWork(workSubNameRef, o, insertStartEndWorkStepNodeFunc)
			workSubNameRef = iworkutil.GetSingleRelativeValueWithReg(workSubNameRef) // 去除多余的 ; 等字符
			workSubName := strings.Replace(workSubNameRef, "$WORK.", "", -1)         // 去除前缀和多余的其它字符
			// 维护 work 的 WorkSubId 属性
			subWork, _ := models.QueryWorkByName(workSubName, o)
			step.WorkSubId = subWork.Id
			break
		}
	}
	models.InsertOrUpdateWorkStep(&step, o)
}

func createOrUpdateSubWork(work_name string, o orm.Ormer, insertStartEndWorkStepNodeFunc func(work_id int64, o orm.Ormer) error) error {
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
			return insertStartEndWorkStepNodeFunc(work.Id, o)
		}
	}
	return nil
}
