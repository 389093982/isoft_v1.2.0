package runtimecfg

import (
	"fmt"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"os"
	"path/filepath"
	"runtime"
)

// 静态文件服务器地址
var FileSavePath string

func init() {
	_, file, _, _ := runtime.Caller(0)
	FileSavePath = fmt.Sprintf("%s/upload", filepath.Dir(filepath.Dir(filepath.Dir(file))))
	os.MkdirAll(fileutil.ChangeToLinuxSeparator(FileSavePath), os.ModePerm)
}
