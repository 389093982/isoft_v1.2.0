package controllers

import (
	"fmt"
	"isoft/isoft/common/fileutils"
	"isoft/isoft/common/xmlutil"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var persistentDirPath string

func init() {
	// 获取 persistent 目录
	_, file, _, _ := runtime.Caller(0)
	persistentDirPath = fileutils.ChangeToLinuxSeparator(fmt.Sprintf("%s/persistent2", filepath.Dir(filepath.Dir(file))))
	os.MkdirAll(path.Join(persistentDirPath, "modules"), os.ModePerm)
}

func persistentToFile() {
	persistentModulesToFile()
}

func persistentModulesToFile() {
	modules, _ := models.QueryAllModules()
	for _, module := range modules {
		filepath := path.Join(persistentDirPath, "modules", fmt.Sprintf(`%s.module`, module.ModuleName))
		fileutil.WriteFile(filepath, []byte(xmlutil.RenderToString(module)), false)
	}
}
