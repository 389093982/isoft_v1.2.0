package block

import (
	"isoft/isoft_iwork_web/models"
)

type BlockParser struct {
	ReferWork *models.Work
	Steps     []models.WorkStep
}

func (this *BlockParser) ParseToBlockSteps() ([]*BlockStep, map[int64]*BlockStep) {
	blockStepMapper := make(map[int64]*BlockStep)
	// 循环遍历将每一个 step 转换成 BlockStep
	blockSteps := this.ParseToNoRelationBlockSteps()
	for _, blockStep := range blockSteps {
		// 相互之间建立关系
		this.BuildRelation(blockSteps, blockStep)
		blockStepMapper[blockStep.Step.WorkStepId] = blockStep
	}
	return blockSteps, blockStepMapper
}

func (this *BlockParser) ParseToNoRelationBlockSteps() []*BlockStep {
	// 循环遍历将每一个 step 转换成 BlockStep
	blockSteps := make([]*BlockStep, 0)
	for index, _ := range this.Steps {
		bStep := &BlockStep{
			ReferWork: this.ReferWork,
			Step:      &this.Steps[index],
		}
		blockSteps = append(blockSteps, bStep)
	}
	return blockSteps
}

func (this *BlockParser) BuildRelation(blockSteps []*BlockStep, blockStep *BlockStep) {
	if blockStep.Step.WorkStepIndent == 0 {
		this.BuildSiblingRelation(blockSteps, blockStep)
	} else {
		this.BuildParentRelation(blockSteps, blockStep)
		this.BuildSiblingRelation(blockSteps, blockStep)
	}
}

func (this *BlockParser) BuildSiblingRelation(blockSteps []*BlockStep, blockStep *BlockStep) {
	// 找最近同级别缩进作为兄弟节点
	// 1、找同级别 blockSteps
	sameLevelBlockSteps := make([]*BlockStep, 0)
	for _, _blockStep := range blockSteps {
		if _blockStep.Step.WorkStepIndent == blockStep.Step.WorkStepIndent {
			sameLevelBlockSteps = append(sameLevelBlockSteps, _blockStep)
		}
	}
	if len(sameLevelBlockSteps) == 0 {
		return
	}
	// 找同级别 blockSteps 中比当前 blockStep 索引小的
	lowSameLevelBlockSteps := make([]*BlockStep, 0)
	for _, sameLevelBlockStep := range sameLevelBlockSteps {
		if sameLevelBlockStep.Step.WorkStepId < blockStep.Step.WorkStepId {
			lowSameLevelBlockSteps = append(lowSameLevelBlockSteps, sameLevelBlockStep)
		}
	}
	if len(lowSameLevelBlockSteps) == 0 {
		return
	}
	previousSiblingBlockStep := lowSameLevelBlockSteps[len(lowSameLevelBlockSteps)-1]
	previousSiblingBlockStep.AfterBlockStep = blockStep
	previousSiblingBlockStep.SiblingBlockSteps = append(previousSiblingBlockStep.SiblingBlockSteps, blockStep)
	blockStep.PreviousBlockStep = previousSiblingBlockStep
	blockStep.SiblingBlockSteps = append(blockStep.SiblingBlockSteps, previousSiblingBlockStep)

	// 找前置兄弟的前置兄弟结拜
	previousSiblingBlockStep = previousSiblingBlockStep.PreviousBlockStep
	for {
		if previousSiblingBlockStep == nil {
			break
		}
		previousSiblingBlockStep.SiblingBlockSteps = append(previousSiblingBlockStep.SiblingBlockSteps, blockStep)
		blockStep.SiblingBlockSteps = append(blockStep.SiblingBlockSteps, previousSiblingBlockStep)

		previousSiblingBlockStep = previousSiblingBlockStep.PreviousBlockStep
	}
}

func (this *BlockParser) BuildParentRelation(blockSteps []*BlockStep, blockStep *BlockStep) {
	// 找最近前置缩进作为父节点
	// 1、找上一级 blockSteps
	preLevelBlockSteps := make([]*BlockStep, 0)
	for _, _blockStep := range blockSteps {
		if _blockStep.Step.WorkStepIndent+1 == blockStep.Step.WorkStepIndent {
			preLevelBlockSteps = append(preLevelBlockSteps, _blockStep)
		}
	}
	if len(preLevelBlockSteps) == 0 {
		panic("无有效的父级节点")
	}
	// 找上一级 blockSteps 中比当前 blockStep 索引小的
	lowPreLevelBlockSteps := make([]*BlockStep, 0)
	for _, preLevelBlockStep := range preLevelBlockSteps {
		if preLevelBlockStep.Step.WorkStepId < blockStep.Step.WorkStepId {
			lowPreLevelBlockSteps = append(lowPreLevelBlockSteps, preLevelBlockStep)
		}
	}
	if len(preLevelBlockSteps) == 0 {
		panic("无有效的父级节点")
	}
	parentBlockStep := lowPreLevelBlockSteps[len(lowPreLevelBlockSteps)-1]
	// 父节点之后发现缩进大于 1 的更适合父节点
	for i := parentBlockStep.Step.WorkStepId; i < blockStep.Step.WorkStepId; i++ {
		for _, step := range this.Steps {
			if i == step.WorkStepId && step.WorkStepIndent < parentBlockStep.Step.WorkStepIndent {
				panic("父节点之后发现缩进大于 1 的更适合父节点")
			}
		}
	}
	parentBlockStep.ChildBlockSteps = append(parentBlockStep.ChildBlockSteps, blockStep)
	parentBlockStep.HasChildren = true
	blockStep.ParentBlockStep = parentBlockStep
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
