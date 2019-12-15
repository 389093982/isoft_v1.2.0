package framework

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/core/iworkutil/reflectutil"
	"isoft/isoft_iwork_web/models"
	"strings"
)

type WorkSubNode struct {
	node.BaseNode
	WorkStep       *models.WorkStep
	WorkSubRunFunc func(work_id int64, dispatcher *entry.Dispatcher) (trackingId string, receiver *entry.Receiver)
}

func (this *WorkSubNode) Execute(trackingId string) {
	// 获取子流程流程名称
	workSubName := this.WorkCache.SubWorkNameMap[this.WorkStep.WorkStepId]
	// 运行子流程
	if workSubCache, err := iworkcache.GetWorkCacheWithName(workSubName); err == nil {
		this.RunOnceSubWork(workSubCache.WorkId, trackingId, this.TmpDataMap, this.DataStore)
	} else {
		panic(errors.Wrap(err, fmt.Sprintf("Load subWork failed for %s", workSubName)))
	}
}

func (this *WorkSubNode) RunOnceSubWork(work_id int64, trackingId string,
	tmpDataMap map[string]interface{}, dataStore *datastore.DataStore) {
	// 继续传递 request 对象
	tmpDataMap[iworkconst.HTTP_REQUEST_OBJECT] = this.Dispatcher.TmpDataMap[iworkconst.HTTP_REQUEST_OBJECT]
	tmpDataMap["logwriter"] = this.LogWriter // 共享 LoggerWriter
	trackingId, receiver := this.WorkSubRunFunc(work_id, &entry.Dispatcher{TrackingId: trackingId, TmpDataMap: tmpDataMap, TxManger: this.DataStore.TxManger})
	// 接收子流程数据存入 dataStore
	for paramName, paramValue := range receiver.TmpDataMap {
		dataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{paramName: paramValue})
	}
}

func (this *WorkSubNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "work_sub", "子流程信息"},
	}
	return this.BPIS1(paramMap)
}

func (this *WorkSubNode) getOrmer() orm.Ormer {
	if this.O == nil {
		this.O = orm.NewOrm()
	}
	return this.O
}

// 获取动态输入值
func (this *WorkSubNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	// 获取子流程信息
	workSubName := this.getWorkSubName()
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := models.QueryAllWorkStepByWorkName(workSubName, this.getOrmer())
		if err != nil {
			panic(err)
		}
		for _, subStep := range subSteps {
			// 找到子流程起始节点
			if strings.ToUpper(subStep.WorkStepType) == "WORK_START" {
				// 子流程起始节点输入参数
				subItems := this.ParamSchemaCacheParser.GetCacheParamInputSchema(&subStep)
				for _, subItem := range subItems.ParamInputSchemaItems {
					items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: subItem.ParamName})
				}
			}
		}
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkSubNode) getWorkSubName() string {
	// 读取历史输入值
	paramInputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema()
	// 从历史输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	return workSubName
}

func (this *WorkSubNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	// 读取静态输入值
	paramInputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema()
	// 从静态输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := models.QueryAllWorkStepByWorkName(workSubName, this.O)
		if err != nil {
			panic(err)
		}
		for _, subStep := range subSteps {
			// 找到子流程结束节点
			if strings.ToUpper(subStep.WorkStepType) == "WORK_END" {
				// 将当前 paramSchemaCacheParser 拷贝副本
				paramSchemaCacheParser := this.ParamSchemaCacheParser
				// 填充新参数至 ParamSchemaParser
				reflectutil.FillFieldValueToStruct(paramSchemaCacheParser, map[string]interface{}{"WorkStep": &subStep})
				// 子流程结束节点输出参数
				subItems := paramSchemaCacheParser.GetCacheParamOutputSchema()
				for _, subItem := range subItems.ParamOutputSchemaItems {
					items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: subItem.ParamName})
				}
			}
		}
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *WorkSubNode) ValidateCustom() (checkResult []string) {
	workSubName := this.getWorkSubName()
	if workSubName == "" {
		checkResult = append(checkResult, fmt.Sprintf("Empty workSubName was found!"))
		return
	}
	work, err := models.QueryWorkByName(workSubName, orm.NewOrm())
	if err != nil {
		checkResult = append(checkResult, fmt.Sprintf("WorkSubName for %s was not found!", workSubName))
		return
	}

	if startStep, err := models.QueryWorkStepByStepName(work.Id, "start", orm.NewOrm()); err == nil {
		workSubInputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema()
		inputSchema := this.ParamSchemaCacheParser.GetCacheParamInputSchema(&startStep)
		for _, item := range inputSchema.ParamInputSchemaItems {
			exist := false
			for _, workSubItem := range workSubInputSchema.ParamInputSchemaItems {
				if item.ParamName == workSubItem.ParamName {
					exist = true
					break
				}
			}
			if !exist {
				checkResult = append(checkResult, fmt.Sprintf("Miss paramName for %s was found!", item.ParamName))
			}
		}
	} else {
		checkResult = append(checkResult, fmt.Sprintf("Miss start node for worksub %s was found!", work.WorkName))
	}
	return
}

func PrepareEmptyInputForWorkSub(subWorkName string) *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{
		ParamInputSchemaItems: []iworkmodels.ParamInputSchemaItem{
			iworkmodels.ParamInputSchemaItem{
				ParamName:  iworkconst.STRING_PREFIX + "work_sub",
				ParamValue: fmt.Sprintf("$WORK.%s;", subWorkName),
			},
		},
	}
}
