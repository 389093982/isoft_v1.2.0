package regist

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkplugin/node/file"
	"reflect"
	"strings"
)

var typeMap = make(map[string]reflect.Type, 0)

func RegistNodes() {
	vs := []interface{}{
		node.WorkStartNode{},
		node.WorkEndNode{},
		node.WorkSubNode{},
		node.SQLExecuteNode{},
		node.SQLQueryNode{},
		node.SQLQueryPageNode{},
		node.JsonRenderNode{},
		node.JsonParserNode{},
		node.HttpRequestNode{},
		node.MapperNode{},
		file.FileReadNode{},
		file.FileWriteNode{},
		file.FileSyncNode{},
		file.FileDeleteNode{},
		node.HrefParserNode{},
		node.EntityParserNode{},
		node.CalHashNode{},
		node.SetEnvNode{},
		node.GetEnvNode{},
		node.RunCmdNode{},
		node.SftpUploadNode{},
		node.SSHShellNode{},
		node.TarGzUnCompressNode{},
		node.TarGzCompressNode{},
		node.IniReadNode{},
		node.IniWriteNode{},
		node.IFNode{},
		node.ElIfNode{},
		node.ElseNode{},
		node.EmptyNode{},
		node.Base64EncodeNode{},
		node.Base64DecodeNode{},
		node.DefineVarNode{},
		node.AssignVarNode{},
		node.MapNode{},
		node.ForeachNode{},
		node.PanicErrorNode{},
	}
	for _, v := range vs {
		t := reflect.ValueOf(v).Type()
		typeMap[strings.ToUpper(t.Name())] = t
	}

	node.Regist(typeMap)
}
