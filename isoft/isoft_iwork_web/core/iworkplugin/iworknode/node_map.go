package iworknode

import (
	"isoft/isoft_iwork_web/models/iwork"
)

type MapNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}
