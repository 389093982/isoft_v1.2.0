package block

import (
	"isoft/isoft_iwork_web/models/iwork"
)

type BlockStep struct {
	ReferWork           *iwork.Work     // 关联的 work
	Step                *iwork.WorkStep // 步骤
	HasChildren         bool            // 是否有子步骤
	ChildBlockSteps     []*BlockStep    // 子步骤列表
	ParentBlockStep     *BlockStep      // 父级 BlockStep
	PreviousBlockStep   *BlockStep      // 前置 BlockStep
	AfterBlockStep      *BlockStep      // 后置 BlockStep
	SiblingBlockSteps   []*BlockStep    // 兄弟步骤列表,包括自己
	AfterJudgeInterrupt bool            // 随后判断 blockStep 停止
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func ParseToBlockStep(steps []iwork.WorkStep) []*BlockStep {
	parser := &BlockParser{Steps: steps}
	blockSteps, _ := parser.ParseToBlockSteps()
	return blockSteps
}
