package block

import (
	"isoft/isoft_iwork_web/models"
)

type BlockStep struct {
	ReferWork           *models.Work     // 关联的 work
	Step                *models.WorkStep // 步骤
	HasChildren         bool             // 是否有子步骤
	ChildBlockSteps     []*BlockStep     // 子步骤列表
	ParentBlockStep     *BlockStep       // 父级 BlockStep
	PreviousBlockStep   *BlockStep       // 前置 BlockStep
	AfterBlockStep      *BlockStep       // 后置 BlockStep
	SiblingBlockSteps   []*BlockStep     // 兄弟步骤列表,包括自己
	AfterJudgeInterrupt bool             // 随后判断 blockStep 停止
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func ParseToBlockStep(steps []models.WorkStep) []*BlockStep {
	parser := &BlockParser{Steps: steps}
	blockSteps, _ := parser.ParseToBlockSteps()
	return blockSteps
}

func GetTopLevelBlockSteps(blockSteps []*BlockStep) (topLevelblockSteps []*BlockStep) {
	for _, blockStep := range blockSteps {
		if blockStep.ParentBlockStep == nil {
			topLevelblockSteps = append(topLevelblockSteps, blockStep)
		}
	}
	return
}
