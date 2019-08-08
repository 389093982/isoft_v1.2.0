package dipool

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/startup/dipool/pool"
)

func init() {
	parser := node.ParamSchemaParser{}
	pool.Container.SetSingleton("parser", &parser)
}
