package framework

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
)

type MapNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}
