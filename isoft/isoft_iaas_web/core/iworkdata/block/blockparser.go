package block

import (
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/models/iwork"
	"sort"
)

type BlockParser2 struct {
	ReferWork *iwork.Work
	Steps     []iwork.WorkStep
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func (this *BlockParser2) ParseToBlockStep() ([]*BlockStep, map[int64]*BlockStep) {
	// 获取顶级 blockStep
	return this.ParseToParentBlockSteps(-1, -1)
}

func (this *BlockParser2) filterStepsBetweenIndex(prefixIndex, suffixIndex int) []iwork.WorkStep {
	if prefixIndex < 0 || suffixIndex < 0 {
		return this.Steps
	}
	return this.Steps[prefixIndex+1 : suffixIndex]
}

func (this *BlockParser2) ParseToParentBlockSteps(prefixIndex, suffixIndex int) ([]*BlockStep, map[int64]*BlockStep) {
	blockSteps := make([]*BlockStep, 0)
	blockStepMapper := make(map[int64]*BlockStep)
	// 获取 prefixIndex 到 suffixIndex 之间所有 step 最小缩进值
	betweenSteps := this.filterStepsBetweenIndex(prefixIndex, suffixIndex)
	if len(betweenSteps) <= 0 {
		return blockSteps, blockStepMapper
	}
	minIndentIndexs := this.getMinIndentIndex(betweenSteps)
	for _, minIndentIndex := range minIndentIndexs {
		// 循环遍历每一个最小缩进的 BlockStep
		bStep := &BlockStep{
			ReferWork: this.ReferWork,
			Step:      &this.Steps[minIndentIndex],
		}
		blockSteps = append(blockSteps, bStep)
		blockStepMapper[bStep.Step.Id] = bStep
	}

	for index, blockStep := range blockSteps[:len(blockSteps)-1] {
		// 为顶级 blockStep 填充子级 blockStep
		prefixIndex, suffixIndex := this.getStepIndex(blockSteps[index].Step.Id), this.getStepIndex(blockSteps[index+1].Step.Id)
		childBlockSteps, childBlockStepMapper := this.ParseToParentBlockSteps(prefixIndex, suffixIndex)
		if len(childBlockSteps) > 0 {
			blockStep.HasChildren = true
			blockStep.ChildBlockSteps = childBlockSteps
			for _, childBlockStep := range childBlockSteps {
				// 设置 parent 属性
				childBlockStep.ParentBlockStep = blockStep
			}
			for key, value := range childBlockStepMapper {
				blockStepMapper[key] = value
			}
		}
	}
	return blockSteps, blockStepMapper
}

// 获取同批最小缩进值索引
func (this *BlockParser2) getMinIndentIndex(steps []iwork.WorkStep) []int {
	indentMap := make(map[int][]int, 0)
	for index, step := range steps {
		if _, ok := indentMap[step.WorkStepIndent]; !ok {
			indentMap[step.WorkStepIndent] = make([]int, 0)
		}
		indentMap[step.WorkStepIndent] = append(indentMap[step.WorkStepIndent], index)
	}
	indents := datatypeutil.GetMapKeySlice(indentMap, []int{}).([]int)
	sort.Ints(indents)
	return indentMap[indents[0]]
}

func (this *BlockParser2) getAllStepIds() []int64 {
	stepIds := make([]int64, 0)
	for _, step := range this.Steps {
		stepIds = append(stepIds, step.Id)
	}
	return stepIds
}

func (this *BlockParser2) getStepIndex(stepId int64) int {
	for index, _stepId := range this.getAllStepIds() {
		if _stepId == stepId {
			return index
		}
	}
	return -1
}

// 判断前置 step 在块范围内是否是可访问的
func CheckBlockAccessble(currentBlockStep *BlockStep, checkStepId int64) bool {
	for {
		// 获取父级别 blockStep
		parentBlockStep := currentBlockStep.ParentBlockStep
		if parentBlockStep == nil { // 顶层 blockStep
			if checkStepId < currentBlockStep.Step.Id {
				return true
			}
			return false
		}
		for _, cBlockStep := range parentBlockStep.ChildBlockSteps {
			if cBlockStep.Step.WorkStepId == checkStepId {
				return true
			}
		}
		currentBlockStep = parentBlockStep
	}
}
