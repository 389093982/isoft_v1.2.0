package regist

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkplugin/node/bash"
	"isoft/isoft_iwork_web/core/iworkplugin/node/chiper"
	"isoft/isoft_iwork_web/core/iworkplugin/node/file"
	"isoft/isoft_iwork_web/core/iworkplugin/node/framework"
	"isoft/isoft_iwork_web/core/iworkplugin/node/html"
	"isoft/isoft_iwork_web/core/iworkplugin/node/json"
	"isoft/isoft_iwork_web/core/iworkplugin/node/os"
	"isoft/isoft_iwork_web/core/iworkplugin/node/sql"
	"isoft/isoft_iwork_web/core/iworkplugin/node/zip"
	"reflect"
	"strings"
)

var typeMap = make(map[string]reflect.Type, 0)

func RegistNodes() {
	vs := []interface{}{
		framework.WorkStartNode{},
		framework.WorkEndNode{},
		framework.WorkSubNode{},
		sql.SQLExecuteNode{},
		sql.SQLQueryNode{},
		json.JsonRenderNode{},
		json.JsonParserNode{},
		html.HttpRequestNode{},
		framework.MapperNode{},
		file.FileReadNode{},
		file.FileWriteNode{},
		file.FileSyncNode{},
		file.FileDeleteNode{},
		html.HrefParserNode{},
		framework.EntityParserNode{},
		chiper.CalHashNode{},
		os.SetEnvNode{},
		os.GetEnvNode{},
		bash.RunCmdNode{},
		os.SftpUploadNode{},
		bash.SSHShellNode{},
		zip.TarGzUnCompressNode{},
		zip.TarGzCompressNode{},
		file.IniReadNode{},
		file.IniWriteNode{},
		framework.IFNode{},
		framework.ElIfNode{},
		framework.ElseNode{},
		framework.EmptyNode{},
		chiper.Base64EncodeNode{},
		chiper.Base64DecodeNode{},
		framework.DefineVarNode{},
		framework.AssignVarNode{},
		framework.MapNode{},
		framework.ForeachNode{},
		framework.PanicErrorNode{},
	}
	for _, v := range vs {
		t := reflect.ValueOf(v).Type()
		typeMap[strings.ToUpper(t.Name())] = t
	}

	node.Regist(typeMap)
}
