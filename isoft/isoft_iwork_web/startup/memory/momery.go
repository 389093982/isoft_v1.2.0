package memory

import (
	"isoft/isoft_iwork_web/models"
	"sync"
)

var GlobalVarMap sync.Map
var ResourceMap sync.Map
var FilterMap sync.Map

func FlushAll() {
	FlushMemoryGlobalVar()
	FlushMemoryResource()
	FlushMemoryFilter()
}

func FlushMemoryFilter() {
	filters, _ := models.QueryAllFilters()
	for _, filter := range filters {
		// 不存在则初始化并添加
		// 存在则获取后添加
		if fs, ok := FilterMap.LoadOrStore(filter.FilterWorkName, []models.Filters{filter}); ok {
			FilterMap.Store(filter.FilterWorkName, append(fs.([]models.Filters), filter))
		}
	}
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
