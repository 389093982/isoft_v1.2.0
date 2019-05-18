package iworknode

import (
	"isoft/isoft_iwork_web/models/iwork"
)

type EmptyNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *EmptyNode) Execute(trackingId string) {

}
