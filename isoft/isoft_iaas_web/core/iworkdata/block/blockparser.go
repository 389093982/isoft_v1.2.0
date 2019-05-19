package block

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/models/iwork"
	"sort"
)

type BlockParser struct {
	ReferWork *iwork.Work
	Steps     []iwork.WorkStep
}

func (this *BlockParser) CheckInvalidWorkSteps(blockStepMapper map[int64]*BlockStep) {
	for _, step := range this.Steps {
		if _, ok := blockStepMapper[step.WorkStepId]; !ok {
			panic(errors.New(fmt.Sprintf(`invalid step for %s, never be called!`, step.WorkStepName)))
		}
	}
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func (this *BlockParser) ParseToBlockSteps() ([]*BlockStep, map[int64]*BlockStep) {
	// 获取顶级 blockStep
	blockSteps, blockStepMapper := this.ParseToParentBlockSteps(-1, -1, -1)
	this.CheckInvalidWorkSteps(blockStepMapper)
	return blockSteps, blockStepMapper
}

// 获取 prefixIndex 和 suffixIndex 之间的所有 workStep 信息
func (this *BlockParser) filterStepsBetweenIndex(prefixIndex, suffixIndex int) []iwork.WorkStep {
	if prefixIndex < 0 || suffixIndex < 0 {
		return this.Steps
	}
	return this.Steps[prefixIndex+1 : suffixIndex]
}

// 获取当前层对应的 blockSteps
func (this *BlockParser) ParseToParentBlockSteps(prefixIndex, suffixIndex int, parentMinWorkStepIndent int) ([]*BlockStep, map[int64]*BlockStep) {
	blockSteps := make([]*BlockStep, 0)           // 存放当前层所有的 blockSteps
	blockStepMapper := make(map[int64]*BlockStep) // 存放所有 stepId 和 step 对应 map, 包括 parent 和 children 级别
	// 获取 prefixIndex 到 suffixIndex 之间所有 step 最小缩进值
	betweenSteps := this.filterStepsBetweenIndex(prefixIndex, suffixIndex)
	if len(betweenSteps) <= 0 {
		return blockSteps, blockStepMapper
	}
	currentMinWorkStepIndent, minIndentIndexs := this.getMinIndentIndex(betweenSteps)
	var previousBlockStep *BlockStep
	for _, minIndentIndex := range minIndentIndexs {
		// 循环遍历每一个最小缩进的 BlockStep
		bStep := &BlockStep{
			ReferWork: this.ReferWork,
			Step:      &this.Steps[minIndentIndex],
		}
		if previousBlockStep != nil {
			bStep.PreviousBlockStep = previousBlockStep // 设置前置
			previousBlockStep.AfterBlockStep = bStep
		}
		previousBlockStep = bStep // 存储为前置
		blockSteps = append(blockSteps, bStep)
		blockStepMapper[bStep.Step.WorkStepId] = bStep
		bStep.SiblingBlockSteps = blockSteps // 将当前层所有的 step 存储为兄弟 steps

		if parentMinWorkStepIndent+1 != currentMinWorkStepIndent {
			panic(errors.New(fmt.Sprintf("invalid indent was found for %s!", bStep.Step.WorkStepName)))
		}
	}
	// 获取 children 层对应的 blockSteps
	this.ParseToChildrenBlockSteps(blockSteps, blockStepMapper, currentMinWorkStepIndent)
	return blockSteps, blockStepMapper
}

// 获取 children 层对应的 blockSteps
func (this *BlockParser) ParseToChildrenBlockSteps(blockSteps []*BlockStep, blockStepMapper map[int64]*BlockStep, parentMinWorkStepIndent int) {
	// 循环得到每一个 blockStep
	for index, blockStep := range blockSteps[:len(blockSteps)-1] {
		// 为顶级 blockStep 填充子级 blockStep
		prefixIndex, suffixIndex := this.getStepIndex(blockSteps[index].Step.Id), this.getStepIndex(blockSteps[index+1].Step.Id)
		childBlockSteps, childBlockStepMapper := this.ParseToParentBlockSteps(prefixIndex, suffixIndex, parentMinWorkStepIndent)
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
}

// 获取同批最小缩进值索引
func (this *BlockParser) getMinIndentIndex(steps []iwork.WorkStep) (int, []int) {
	indentMap := make(map[int][]int, 0)
	for _, step := range steps {
		if _, ok := indentMap[step.WorkStepIndent]; !ok {
			indentMap[step.WorkStepIndent] = make([]int, 0)
		}
		indentMap[step.WorkStepIndent] = append(indentMap[step.WorkStepIndent], this.getStepIndex(step.Id))
	}
	indents := datatypeutil.GetMapKeySlice(indentMap, []int{}).([]int)
	sort.Ints(indents)
	return indents[0], indentMap[indents[0]]
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
