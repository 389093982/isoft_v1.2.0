package memory

import (
	"isoft/isoft_iwork_web/models"
	"sync"
)

var GlobalVarMap sync.Map
var ResourceMap sync.Map

func FlushAll() {
	FlushMemoryGlobalVar()
	FlushMemoryResource()
}

func FlushMemoryGlobalVar() {
	globalVars := models.QueryAllGlobalVar()
	for _, globalVar := range globalVars {
		GlobalVarMap.Store(globalVar.Name, globalVar)
	}
}

func FlushMemoryResource() {
	resources := models.QueryAllResource()
	for _, resource := range resources {
		ResourceMap.Store(resource.ResourceName, resource)
	}
}
