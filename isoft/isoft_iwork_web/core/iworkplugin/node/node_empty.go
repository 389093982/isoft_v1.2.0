package node

import (
	"isoft/isoft_iwork_web/models"
)

type EmptyNode struct {
	BaseNode
	WorkStep *models.WorkStep
}

func (this *EmptyNode) Execute(trackingId string) {

}
