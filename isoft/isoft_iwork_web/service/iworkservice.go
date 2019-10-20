package service

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/models"
	"strings"
	"time"
)

func GetRelativeWorkService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)

	_, parentWorks, _, err := models.QueryParentWorks(work_id, o)
	if err != nil {
		return nil, err
	}
	subWorks := make([]models.Work, 0)
	steps, err := models.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return nil, err
	}
	for _, step := range steps {
		if step.WorkSubId > 0 {
			subwork, err := models.QueryWorkById(step.WorkSubId, o)
			if err != nil {
				return nil, err
			}
			subWorks = append(subWorks, subwork)
		}
	}
	result["parentWorks"] = parentWorks
	result["subworks"] = subWorks
	return
}

func RunWork(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	go iworkrun.RunOneWork(work_id, nil)
	return nil
}

func FilterPageLogRecord(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	offset := serviceArgs["offset"].(int)
	current_page := serviceArgs["current_page"].(int)
	logLevel := serviceArgs["logLevel"].(string)
	ctx := serviceArgs["ctx"].(*context.Context)
	runLogRecords, count, err := models.QueryRunLogRecord(work_id, logLevel, current_page, offset)
	if err != nil {
		return nil, err
	}
	paginator := pagination.SetPaginator(ctx, offset, count)
	if err != nil {
		return nil, err
	}
	result["runLogRecords"] = runLogRecords
	result["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	return
}

func FilterPageWorkService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	condArr := serviceArgs["condArr"].(map[string]string)
	offset := serviceArgs["offset"].(int)
	current_page := serviceArgs["current_page"].(int)
	ctx := serviceArgs["ctx"].(*context.Context)
	o := serviceArgs["o"].(orm.Ormer)
	works, count, err := models.QueryWork(condArr, current_page, offset, o)
	if err != nil {
		return nil, err
	}
	paginator := pagination.SetPaginator(ctx, offset, count)
	if err != nil {
		return nil, err
	}
	result["works"] = works
	result["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	return
}

// 根据 id 信息查找旧的 work 信息
func GetOldWorkInfoById(id int64, o orm.Ormer) (oldWorkName string, oldWorkId int64) {
	if id <= 0 {
		return
	}
	if work, err := models.QueryWorkById(id, o); err == nil {
		oldWorkName = work.WorkName
		oldWorkId = work.Id
	}
	return
}

func EditWorkService(serviceArgs map[string]interface{}) error {
	work := serviceArgs["work"].(models.Work)
	o := serviceArgs["o"].(orm.Ormer)
	oldWorkName, oldWorkId := GetOldWorkInfoById(work.Id, o)
	// 插入或者更新 work 信息
	if _, err := models.InsertOrUpdateWork(&work, o); err != nil {
		return err
	}
	if oldWorkName == "" {
		// 新增 work 场景,自动添加开始和结束节点
		if err := InsertStartEndWorkStepNode(work.Id, o); err != nil {
			return err
		}
	} else {
		// 修改 work 场景
		if err := ChangeReferencesWorkName(oldWorkId, oldWorkName, work.WorkName, o); err != nil {
			return err
		}
	}
	return nil
}

func DeleteWorkByIdService(serviceArgs map[string]interface{}) error {
	id := serviceArgs["id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	if work, err := models.QueryWorkById(id, o); err == nil {
		// cronMeta 可以提前删除,删除失败不影响功能
		models.DeleteCronMetaByTaskName(work.WorkName, o)
	}
	return models.DeleteWorkById(id, o)
}

func ChangeReferencesWorkName(work_id int64, oldWorkName, workName string, o orm.Ormer) error {
	if oldWorkName == workName {
		return nil
	}
	_, parentWorks, _, err := models.QueryParentWorks(work_id, o)
	if err != nil {
		return nil
	}
	for _, parentWork := range parentWorks {
		steps, _ := models.QueryAllWorkStepInfo(parentWork.Id, o)
		for _, step := range steps {
			if step.WorkStepType != "work_sub" {
				continue
			}
			inputSchema := node.GetCacheParamInputSchema(&step)
			for index, item := range inputSchema.ParamInputSchemaItems {
				if item.ParamName == iworkconst.STRING_PREFIX+"work_sub" && strings.Contains(item.ParamValue, oldWorkName) {
					inputSchema.ParamInputSchemaItems[index].ParamValue = strings.Replace(item.ParamValue, oldWorkName, workName, -1)
				}
			}
			step.WorkStepInput = inputSchema.RenderToJson()
			models.InsertOrUpdateWorkStep(&step, o)
		}
	}
	return nil
}

func InsertStartEndWorkStepNode(work_id int64, o orm.Ormer) error {
	insertDefaultWorkStepNodeFunc := func(nodeName string, work_step_id int64) error {
		step := &models.WorkStep{
			WorkId:          work_id,
			WorkStepId:      work_step_id,
			WorkStepName:    nodeName,
			WorkStepDesc:    fmt.Sprintf("%s节点", nodeName),
			WorkStepType:    fmt.Sprintf("work_%s", nodeName),
			IsDefer:         "false",
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _, err := models.InsertOrUpdateWorkStep(step, o); err != nil {
			return err
		}
		return nil
	}
	if err := insertDefaultWorkStepNodeFunc("start", 1); err != nil {
		return err
	}
	if err := insertDefaultWorkStepNodeFunc("end", 2); err != nil {
		return err
	}
	return nil
}
