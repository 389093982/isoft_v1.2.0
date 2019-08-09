package memory

import (
	"isoft/isoft_iwork_web/models"
	"sync"
)

var MemoryGlobalVarMap sync.Map

func FlushAll() {
	FlushMemoryGlobalVar()
}

func FlushMemoryGlobalVar() {
	globalVars := models.QueryAllGlobalVar()
	for _, globalVar := range globalVars {
		MemoryGlobalVarMap.Store(globalVar.Name, globalVar)
	}
}
