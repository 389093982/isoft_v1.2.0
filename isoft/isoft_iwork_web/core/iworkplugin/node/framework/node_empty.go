package framework

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type EmptyNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *EmptyNode) Execute(trackingId string) {

}
