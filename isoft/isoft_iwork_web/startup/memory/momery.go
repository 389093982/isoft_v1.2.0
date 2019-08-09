package memory

import (
	"isoft/isoft_iwork_web/models"
	"sync"
)

var MemoryGlobalVarMap sync.Map
var MemoryResourceMap sync.Map

func FlushAll() {
	FlushMemoryGlobalVar()
	FlushMemoryResource()
}

func FlushMemoryGlobalVar() {
	globalVars := models.QueryAllGlobalVar()
	for _, globalVar := range globalVars {
		MemoryGlobalVarMap.Store(globalVar.Name, globalVar)
	}
}

func FlushMemoryResource() {
	resources := models.QueryAllResource()
	for _, resource := range resources {
		MemoryResourceMap.Store(resource.ResourceName, resource)
	}
}
