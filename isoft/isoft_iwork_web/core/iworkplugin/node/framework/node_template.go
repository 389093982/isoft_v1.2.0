package framework

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type TemplateNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *TemplateNode) Execute(trackingId string) {

}
