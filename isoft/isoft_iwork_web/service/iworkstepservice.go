package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkbuild"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkplugin/node/framework"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/core/iworkvalid"
	"isoft/isoft_iwork_web/models"
	"strings"
	"time"
)

func LoadResourceInfo() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{
		ParamOutputSchemaItems: []iworkmodels.ParamOutputSchemaItem{},
	}
	resources := models.QueryAllResource()
	for _, resource := range resources {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{
			ParamName: resource.ResourceName,
		})
	}
	return pos
}

func LoadWorkInfo() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{
		ParamOutputSchemaItems: []iworkmodels.ParamOutputSchemaItem{},
	}
	works := models.QueryAllWorkInfo(orm.NewOrm())
	for _, work := range works {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{
			ParamName: work.WorkName,
		})
	}
	return pos
}

func LoadGlobalVarInfo() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{
		ParamOutputSchemaItems: []iworkmodels.ParamOutputSchemaItem{},
	}
	globalVars := models.QueryAllGlobalVar()
	for _, globalVar := range globalVars {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{
			ParamName: globalVar.Name,
		})
	}
	return pos
}

func LoadErrorInfo() *iworkmodels.ParamOutputSchema {
	pos := &iworkmodels.ParamOutputSchema{
		ParamOutputSchemaItems: []iworkmodels.ParamOutputSchemaItem{},
	}
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{ParamName: "isError"})
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{ParamName: "isNoError"})
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{ParamName: "errorMsg"})
	pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, iworkmodels.ParamOutputSchemaItem{ParamName: "insensitiveErrorMsg"})
	return pos
}

// 加载前置节点输出参数
func LoadPreNodeOutputService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	prePosTreeNodeArr := make([]*iworkmodels.TreeNode, 0)
	// 加载 resource 参数
	pos := LoadResourceInfo()
	prePosTreeNodeArr = append(prePosTreeNodeArr, pos.RenderToTreeNodes("$RESOURCE"))
	// 加载 work 参
	pos = LoadWorkInfo()
	prePosTreeNodeArr = append(prePosTreeNodeArr, pos.RenderToTreeNodes("$WORK"))

	// 加载 error 参数
	pos = LoadErrorInfo()
	prePosTreeNodeArr = append(prePosTreeNodeArr, pos.RenderToTreeNodes("$Error"))

	// 加载 globalVar 参数
	pos = LoadGlobalVarInfo()
	prePosTreeNodeArr = append(prePosTreeNodeArr, pos.RenderToTreeNodes("$Global"))

	// 加载前置步骤输出
	if steps, err := models.QueryAllPreStepInfo(work_id, work_step_id, o); err == nil {
		// 当前步骤信息
		currentWorkStep, _ := models.QueryWorkStepInfo(work_id, work_step_id, orm.NewOrm())
		// 所有步骤信息
		allSteps, _ := models.QueryAllWorkStepInfo(work_id, orm.NewOrm())
		parser := &block.BlockParser{Steps: allSteps}
		_, blockStepMapper := parser.ParseToBlockSteps()
		currentBlockStep := blockStepMapper[currentWorkStep.WorkStepId]
		for _, step := range steps {
			// 判断前置 step 在块范围内是否是可访问的,且是否非 defer 步骤
			if block.CheckBlockAccessble(currentBlockStep, step.WorkStepId) && step.IsDefer != "true" {
				pos := node.GetCacheParamOutputSchema(&step)
				prePosTreeNodeArr = append(prePosTreeNodeArr, pos.RenderToTreeNodes("$"+step.WorkStepName))
			}
		}
	}
	// 返回结果
	result["prePosTreeNodeArr"] = prePosTreeNodeArr
	return
}

func GetAllWorkStepInfoService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	steps, err := models.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return nil, err
	}
	result["steps"] = steps
	return
}

func LoadWorkStepInfoService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	// 读取 work_step 信息
	step, err := models.QueryWorkStepInfo(work_id, work_step_id, o)
	if err != nil {
		return nil, err
	}
	var paramMappingsArr []iworkmodels.ParamMapping
	json.Unmarshal([]byte(step.WorkStepParamMapping), &paramMappingsArr)
	result["step"] = step
	result["paramInputSchema"] = node.GetCacheParamInputSchema(&step)
	result["paramOutputSchema"] = node.GetCacheParamOutputSchema(&step)
	result["paramOutputSchemaTreeNode"] = node.GetCacheParamOutputSchema(&step).RenderToTreeNodes("output")
	result["paramMappings"] = paramMappingsArr
	return
}

func CopyWorkStepByWorkStepIdService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	step, err := models.QueryOneWorkStep(work_id, work_step_id, o)
	if err != nil {
		return err
	}
	step.Id = 0
	step.WorkStepId = step.WorkStepId + 1
	step.WorkStepName = step.WorkStepName + "_copy"
	return insertWorkStepAfter(work_id, work_step_id, &step, o)
}

func insertWorkStepAfter(work_id, work_step_id int64, step *models.WorkStep, o orm.Ormer) error {
	// 将 work_step_id 之后的所有节点后移一位
	if err := models.BatchChangeWorkStepIdOrder(work_id, work_step_id, "+", o); err != nil {
		return err
	}
	_, err := models.InsertOrUpdateWorkStep(step, o)
	return err
}

func DeleteWorkStepByWorkStepIdService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	if step, err := models.QueryWorkStepInfo(work_id, work_step_id, o); err == nil {
		if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
			return errors.New("start 节点和 end 节点不能被删除!")
		}
	}
	if err := models.DeleteWorkStepByWorkStepId(work_id, work_step_id, o); err != nil {
		return err
	}
	return nil
}

func WorkStepListService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	condArr := make(map[string]interface{})
	condArr["work_id"] = serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	worksteps, err := models.QueryWorkStep(condArr, o)
	if err != nil {
		return nil, err
	}
	result["worksteps"] = worksteps
	return
}

func AddWorkStepService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	work_step_meta := serviceArgs["work_step_meta"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	// end 节点之后不能添加节点
	if step, err := models.QueryWorkStepInfo(work_id, work_step_id, o); err == nil {
		if step.WorkStepType == "work_end" {
			return errors.New("不能再 end 节点后面添加节点!")
		}
	}

	if strings.HasPrefix(work_step_meta, "work_type__") {
		work_step_type := strings.TrimPrefix(work_step_meta, "work_type__")
		step := models.WorkStep{
			WorkId:          work_id,
			WorkStepName:    work_step_type + "_" + fmt.Sprintf("%v", time.Now().Unix()),
			WorkStepType:    work_step_type,
			WorkStepDesc:    "",
			IsDefer:         "false", // 默认不延迟执行
			WorkStepIndent:  0,       // 默认缩进级别为 0
			WorkStepId:      work_step_id + 1,
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		return insertWorkStepAfter(work_id, work_step_id, &step, o)
	} else if strings.HasPrefix(work_step_meta, "work_name__") {
		subWorkName := strings.TrimPrefix(work_step_meta, "work_name__")
		workStepName := strings.Join([]string{iworkconst.NODE_TYPE_WORK_SUB, subWorkName,
			fmt.Sprintf("%v", time.Now().Unix())}, "_")
		step := models.WorkStep{
			WorkId:          work_id,
			WorkStepName:    workStepName,
			WorkStepType:    iworkconst.NODE_TYPE_WORK_SUB,
			WorkStepInput:   framework.PrepareEmptyInputForWorkSub(subWorkName).RenderToJson(),
			WorkStepDesc:    "",
			IsDefer:         "false", // 默认不延迟执行
			WorkStepIndent:  0,       // 默认缩进级别为 0
			WorkStepId:      work_step_id + 1,
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if err := insertWorkStepAfter(work_id, work_step_id, &step, o); err == nil {
			// 动态构建输入输出
			BuildDynamic(work_id, step.WorkStepId, step, o)
		} else {
			return err
		}
	}
	return nil
}

// 更改邻近两个节点的顺序
func changeNearWorkStepOrder(work_id, work_step_id int64, o orm.Ormer, nearStepLength int64) error {
	// 获取当前步骤
	step, err := models.QueryOneWorkStep(work_id, work_step_id, o)
	if err != nil {
		return err
	}
	// 获取邻近步骤
	nearStep, err := models.QueryOneWorkStep(work_id, work_step_id-nearStepLength, o)
	if err != nil {
		return err
	}
	if nearStep.WorkStepType == "work_start" || nearStep.WorkStepType == "work_end" {
		return errors.New("start 节点和 end 节点不能移动位置!")
	}
	nearStep.WorkStepId = nearStep.WorkStepId + nearStepLength
	step.WorkStepId = step.WorkStepId - nearStepLength
	// 更新邻近步骤
	if _, err := models.InsertOrUpdateWorkStep(&nearStep, o); err != nil {
		return err
	}
	// 更新当前步骤
	if _, err := models.InsertOrUpdateWorkStep(&step, o); err != nil {
		return err
	}
	return nil
}

func ChangeWorkStepOrderService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	_type := serviceArgs["_type"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	if _type == "up" {
		return changeNearWorkStepOrder(work_id, work_step_id, o, 1)
	} else {
		return changeNearWorkStepOrder(work_id, work_step_id, o, -1)
	}
	return nil
}

func EditWorkStepBaseInfoService(serviceArgs map[string]interface{}) error {
	o := serviceArgs["o"].(orm.Ormer)
	step, err := models.QueryOneWorkStep(serviceArgs["work_id"].(int64), serviceArgs["work_step_id"].(int64), o)
	if err != nil {
		return err
	}
	old_work_step_name := step.WorkStepName
	old_work_step_type := step.WorkStepType
	// 替换成新值
	step.WorkStepName = serviceArgs["work_step_name"].(string)
	step.WorkStepType = serviceArgs["work_step_type"].(string)
	step.WorkStepDesc = serviceArgs["work_step_desc"].(string)
	step.IsDefer = serviceArgs["is_defer"].(string)
	// 变更类型需要置空 input 和 output 参数
	if step.WorkStepType != old_work_step_type {
		step.WorkStepInput = ""
		step.WorkStepOutput = ""
	}
	if _, err := models.InsertOrUpdateWorkStep(&step, o); err != nil {
		return err
	}
	// 级联更改相关联的步骤名称
	if err := ChangeReferencesWorkStepName(step.WorkId, old_work_step_name, step.WorkStepName, o); err != nil {
		return err
	}
	return nil
}

func ChangeReferencesWorkStepName(work_id int64, oldWorkStepName, workStepName string, o orm.Ormer) error {
	if oldWorkStepName == workStepName {
		return nil
	}
	steps, err := models.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return err
	}
	for _, step := range steps {
		step.WorkStepInput = strings.Replace(step.WorkStepInput, "$"+oldWorkStepName, "$"+workStepName, -1)
		_, err := models.InsertOrUpdateWorkStep(&step, o)
		if err != nil {
			return err
		}
	}
	return nil
}

func createSubWork(refactor_worksub_name string, o orm.Ormer) (int64, error) {
	// 创建子流程
	subWork := &models.Work{
		WorkName:        refactor_worksub_name,
		WorkDesc:        "refactor worksub",
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	if _, err := models.InsertOrUpdateWork(subWork, o); err != nil {
		return -1, err
	}
	// 为子流程添加开始和结束节点
	if err := InsertStartEndWorkStepNode(subWork.Id, o); err != nil {
		return -1, err
	}
	return subWork.Id, nil
}

func getRefactorWorkStep(work_id, work_step_id int64, o orm.Ormer) (step models.WorkStep, err error) {
	step, err = models.QueryWorkStepInfo(work_id, int64(work_step_id), o)
	if err != nil {
		return
	}
	if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
		return step, errors.New("start 和 end 节点不能重构！")
	}
	return step, nil
}

func refactorSubWork(refactorStep models.WorkStep, subWorkId int64, o orm.Ormer) error {
	// 将子流程 start 节点后面所有的步骤 id 顺序 + 1
	err := models.BatchChangeWorkStepIdOrder(subWorkId, 1, "+", o)
	if err != nil {
		return err
	}
	newStep := models.CopyWorkStepInfo(refactorStep)
	newStep.WorkId = subWorkId
	newStep.WorkStepId = 2
	// 在 2 号步骤位置插入当前步骤
	if _, err := models.InsertOrUpdateWorkStep(newStep, o); err != nil {
		return err
	}
	return nil
}

func refactorCurrentWorkByDelete(refactorStep models.WorkStep, o orm.Ormer) error {
	_serviceArgs := map[string]interface{}{"work_id": refactorStep.WorkId, "work_step_id": refactorStep.WorkStepId, "o": o}
	return DeleteWorkStepByWorkStepIdService(_serviceArgs)
}

func refactorCurrentWorkByChangeToWorkSub(subWorkId int64, refactor_worksub_name string, refactorStep models.WorkStep, o orm.Ormer) error {
	// 修改 refactorStep 的类型
	refactorStep.WorkStepType = "work_sub"
	// 修改 refactorStep 的 subWorkId
	refactorStep.WorkSubId = subWorkId
	// 修改 refactorStep 的 WorkStepInput
	factory := node.WorkStepFactory{WorkStep: &refactorStep}
	inputSchema := factory.GetDefaultParamInputSchema()
	for index, item := range inputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.STRING_PREFIX+"work_sub" {
			item.ParamValue = fmt.Sprintf("$WORK.%s;", refactor_worksub_name)
			inputSchema.ParamInputSchemaItems[index] = item
			break
		}
	}
	refactorStep.WorkStepOutput = (&iworkmodels.ParamOutputSchema{}).RenderToJson() // 输出置空
	refactorStep.WorkStepInput = inputSchema.RenderToJson()
	if _, err := models.InsertOrUpdateWorkStep(&refactorStep, o); err != nil {
		return err
	}
	return nil
}

func BatchChangeIndentService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	mod := serviceArgs["mod"].(string)
	indent_work_step_ids := serviceArgs["indent_work_step_ids"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	var indent_work_step_id_arr []int
	json.Unmarshal([]byte(indent_work_step_ids), &indent_work_step_id_arr)
	for _, work_step_id := range indent_work_step_id_arr {
		if step, err := models.QueryWorkStepInfo(work_id, int64(work_step_id), o); err == nil {
			if mod == "left" && step.WorkStepIndent > 0 {
				step.WorkStepIndent -= 1
			} else if mod == "right" {
				step.WorkStepIndent += 1
			}
			if _, err := models.InsertOrUpdateWorkStep(&step, o); err != nil {
				return err
			}
		}
	}
	return nil
}

func RefactorWorkStepInfoService(serviceArgs map[string]interface{}) error {
	// 获取参数
	work_id := serviceArgs["work_id"].(int64)
	refactor_worksub_name := serviceArgs["refactor_worksub_name"].(string)
	refactor_work_step_ids := serviceArgs["refactor_work_step_ids"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	var refactor_work_step_id_arr []int
	json.Unmarshal([]byte(refactor_work_step_ids), &refactor_work_step_id_arr)
	// 校验 refactor_work_step_id_arr 是否连续
	if refactor_work_step_id_arr[len(refactor_work_step_id_arr)-1]-refactor_work_step_id_arr[0] != len(refactor_work_step_id_arr)-1 {
		return errors.New("refactor workStepId 必须是连续的!")
	}
	// 创建子流程
	subWorkId, err := createSubWork(refactor_worksub_name, o)
	if err != nil {
		return err
	}
	// 循环移动子步骤,移动一个删除一个,反转slice,从 id 大的开始执行
	for index, work_step_id := range datatypeutil.ReverseSlice(refactor_work_step_id_arr).([]int) {
		refactorStep, err := getRefactorWorkStep(work_id, int64(work_step_id), o)
		if err != nil {
			return err
		}
		if err := refactorSubWork(refactorStep, subWorkId, o); err != nil {
			return err
		}
		if index == len(refactor_work_step_id_arr)-1 {
			// 最后一次操作不再是删除,而是替换成子节点
			if err := refactorCurrentWorkByChangeToWorkSub(subWorkId, refactor_worksub_name, refactorStep, o); err != nil {
				return err
			}
		} else {
			// 删除节点
			if err := refactorCurrentWorkByDelete(refactorStep, o); err != nil {
				return err
			}
		}
	}
	return nil
}

func formatChecker(paramInputSchema *iworkmodels.ParamInputSchema) error {
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		formatChecker := iworkvalid.ParamValueFormatChecker{
			ParamName:  item.ParamName,
			PureText:   item.PureText,
			ParamValue: item.ParamValue,
		}
		if ok, err := formatChecker.Check(); !ok && err != nil {
			return err
		}
	}
	return nil
}

func EditWorkStepParamInfo(serviceArgs map[string]interface{}) error {
	var (
		step models.WorkStep
		err  error
	)
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	paramInputSchemaStr := serviceArgs["paramInputSchemaStr"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	if step, err = models.QueryOneWorkStep(work_id, work_step_id, o); err != nil {
		return err
	}
	paramInputSchema, _ := iworkmodels.ParseToParamInputSchema(paramInputSchemaStr)
	if err = formatChecker(paramInputSchema); err != nil {
		return err
	}

	// 保存完静态参数后自动构建获动态参数并保存
	BuildDynamic(work_id, work_step_id, step, o)

	// 编辑开始或结束节点时需要通知调度流程重新 BuildDynamic
	if step.WorkStepType == iworkconst.NODE_TYPE_WORK_START || step.WorkStepType == iworkconst.NODE_TYPE_WORK_END {
		BuildParentWork(work_id, step, o)
	}
	return nil
}

func BuildParentWork(work_id int64, step models.WorkStep, o orm.Ormer) {
	if workSteps, _, _, err := models.QueryParentWorks(work_id, o); err == nil {
		for _, workStep := range workSteps {
			BuildDynamic(workStep.WorkId, workStep.WorkStepId, step, o)
		}
	}
}

// 构建动态值,每次 build 之前需要重读 step 信息
func BuildDynamic(work_id int64, work_step_id int64, step models.WorkStep, o orm.Ormer) {
	step, _ = models.QueryWorkStepInfo(work_id, work_step_id, o)
	// 自动创建子流程
	iworkbuild.BuildAutoCreateSubWork(step, o, InsertStartEndWorkStepNode)
	step, _ = models.QueryWorkStepInfo(work_id, work_step_id, o)
	// 构建动态输入值
	iworkbuild.BuildDynamicInput(step, o)

	step, _ = models.QueryWorkStepInfo(work_id, work_step_id, o)
	// 构建动态输出值
	iworkbuild.BuildDynamicOutput(step, o)
}
