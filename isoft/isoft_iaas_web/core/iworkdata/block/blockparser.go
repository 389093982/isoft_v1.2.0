package block

import (
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/models/iwork"
	"sort"
)

type BlockParser struct {
	ReferWork *iwork.Work
	Steps     []iwork.WorkStep
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func (this *BlockParser) ParseToBlockSteps() ([]*BlockStep, map[int64]*BlockStep) {
	// 获取顶级 blockStep
	return this.ParseToParentBlockSteps(-1, -1)
}

func (this *BlockParser) filterStepsBetweenIndex(prefixIndex, suffixIndex int) []iwork.WorkStep {
	if prefixIndex < 0 || suffixIndex < 0 {
		return this.Steps
	}
	return this.Steps[prefixIndex+1 : suffixIndex]
}

// 获取当前层对应的 blockSteps
func (this *BlockParser) ParseToParentBlockSteps(prefixIndex, suffixIndex int) ([]*BlockStep, map[int64]*BlockStep) {
	blockSteps := make([]*BlockStep, 0)           // 存放当前层所有的 blockSteps
	blockStepMapper := make(map[int64]*BlockStep) // stepId 和 step 对应 map
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
		bStep.SiblingBlockSteps = blockSteps // 将当前层所有的 step 存储为兄弟 steps
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
func (this *BlockParser) getMinIndentIndex(steps []iwork.WorkStep) []int {
	indentMap := make(map[int][]int, 0)
	for _, step := range steps {
		if _, ok := indentMap[step.WorkStepIndent]; !ok {
			indentMap[step.WorkStepIndent] = make([]int, 0)
		}
		indentMap[step.WorkStepIndent] = append(indentMap[step.WorkStepIndent], this.getStepIndex(step.Id))
	}
	indents := datatypeutil.GetMapKeySlice(indentMap, []int{}).([]int)
	sort.Ints(indents)
	return indentMap[indents[0]]
}

func (this *BlockParser) getAllStepIds() []int64 {
	stepIds := make([]int64, 0)
	for _, step := range this.Steps {
		stepIds = append(stepIds, step.Id)
	}
	return stepIds
}

func (this *BlockParser) getStepIndex(stepId int64) int {
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
		// 从兄弟节点中查找
		for _, siblingBlockStep := range currentBlockStep.SiblingBlockSteps {
			if siblingBlockStep.Step.WorkStepId == checkStepId {
				return checkStepId < currentBlockStep.Step.WorkStepId
			}
		}
		// 从父节点中查找
		parentBlockStep := currentBlockStep.ParentBlockStep
		if parentBlockStep != nil {
			if parentBlockStep.Step.WorkStepId == checkStepId {
				return true
			}
			currentBlockStep = parentBlockStep
		} else {
			return false
		}
	}
}
