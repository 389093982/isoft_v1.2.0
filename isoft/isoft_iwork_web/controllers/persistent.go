package controllers

import (
	"fmt"
	"isoft/isoft/common/fileutils"
	"isoft/isoft/common/xmlutil"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"path"
	"path/filepath"
	"runtime"
)

var persistentDirPath string

func init() {
	// 获取 persistent 目录
	_, file, _, _ := runtime.Caller(0)
	persistentDirPath = fileutils.ChangeToLinuxSeparator(fmt.Sprintf("%s/persistent2", filepath.Dir(filepath.Dir(file))))
}

func persistentToFile() {
	persistentFiltersToFile()
	persistentModulesToFile()
	persistentGlobalVarsToFile()
	persistentQuartzsToFile()
	persistentResourceToFile()
}

func persistentModulesToFile() {
	modules, _ := models.QueryAllModules()
	for _, module := range modules {
		filepath := path.Join(persistentDirPath, "modules", fmt.Sprintf(`%s.module`, module.ModuleName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(module)), false)
	}
}

func persistentFiltersToFile() {
	filters, _ := models.QueryAllFilters()
	for _, filter := range filters {
		filepath := path.Join(persistentDirPath, "filters", fmt.Sprintf(`%s.filter`, filter.FilterWorkName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(filter)), false)
	}
}

func persistentGlobalVarsToFile() {
	globalVars := models.QueryAllGlobalVar()
	for _, globalVar := range globalVars {
		filepath := path.Join(persistentDirPath, "globalVars", fmt.Sprintf(`%s.globalVar`, globalVar.Name))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(globalVar)), false)
	}
}

func persistentQuartzsToFile() {
	metas, _ := models.QueryAllCronMeta()
	for _, meta := range metas {
		filepath := path.Join(persistentDirPath, "quartzs", fmt.Sprintf(`%s.quartz`, meta.TaskName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(meta)), false)
	}
}

func persistentResourceToFile() {
	resources := models.QueryAllResource()
	for _, resource := range resources {
		filepath := path.Join(persistentDirPath, "resources", fmt.Sprintf(`%s.resource`, resource.ResourceName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(resource)), false)
	}
}
