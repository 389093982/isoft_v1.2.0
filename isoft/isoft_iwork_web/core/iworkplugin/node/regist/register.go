package regist

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkplugin/node/bash"
	"isoft/isoft_iwork_web/core/iworkplugin/node/chiper"
	"isoft/isoft_iwork_web/core/iworkplugin/node/file"
	"isoft/isoft_iwork_web/core/iworkplugin/node/framework"
	"isoft/isoft_iwork_web/core/iworkplugin/node/http"
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
		sql.SQLExecuteNode{},
		sql.SQLQueryNode{},
		json.JsonRenderNode{},
		json.JsonParserNode{},
		http.HttpRequestNode{},
		file.FileReadNode{},
		file.FileWriteNode{},
		file.FileSyncNode{},
		file.FileDeleteNode{},
		file.DoReceiveFileNode{},
		file.DoResponseReceiveFileNode{},
		http.HrefParserNode{},
		http.HttpRequestParserNode{},
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
		framework.EntityParserNode{},
		framework.WorkStartNode{},
		framework.WorkEndNode{},
		framework.WorkSubNode{},
		framework.MapperNode{},
		framework.IFNode{},
		framework.ElIfNode{},
		framework.ElseNode{},
		framework.EmptyNode{},
		framework.DoErrorFilterNode{},
		framework.DefineVarNode{},
		framework.AssignVarNode{},
		framework.MapNode{},
		framework.ForeachNode{},
		framework.PanicErrorNode{},
		framework.TemplateNode{},
		chiper.Base64EncodeNode{},
		chiper.Base64DecodeNode{},
		chiper.CreateJWTNode{},
		chiper.ParseJWTNode{},
	}
	for _, v := range vs {
		t := reflect.ValueOf(v).Type()
		typeMap[strings.ToUpper(t.Name())] = t
	}

	node.Regist(typeMap)
}
