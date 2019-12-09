package regist

import (
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkplugin/node/bash"
	"isoft/isoft_iwork_web/core/iworkplugin/node/chiper"
	"isoft/isoft_iwork_web/core/iworkplugin/node/file"
	"isoft/isoft_iwork_web/core/iworkplugin/node/framework"
	"isoft/isoft_iwork_web/core/iworkplugin/node/http"
	"isoft/isoft_iwork_web/core/iworkplugin/node/json"
	"isoft/isoft_iwork_web/core/iworkplugin/node/mail"
	"isoft/isoft_iwork_web/core/iworkplugin/node/os"
	"isoft/isoft_iwork_web/core/iworkplugin/node/sql"
	"isoft/isoft_iwork_web/core/iworkplugin/node/zip"
	"reflect"
	"strings"
)

var typeMap = make(map[string]reflect.Type, 0)

func GetNodeMeta() []map[string]string {
	return []map[string]string{
		{"name": "work_start", "icon": "ios-arrow-dropright"},
		{"name": "work_end", "icon": "ios-arrow-dropleft"},
		{"name": "empty", "icon": "ios-mail-open-outline"},
		{"name": "work_sub", "icon": "logo-buffer"},
		{"name": "sql_query", "icon": "ios-cube-outline"},
		{"name": "sql_execute", "icon": "ios-crop-outline"},
		{"name": "db_parser", "icon": "ios-crop-outline"},
		{"name": "json_render", "icon": "ios-git-branch"},
		{"name": "json_parser", "icon": "ios-git-compare"},
		{"name": "http_request", "icon": "ios-globe-outline"},
		{"name": "http_request_parser", "icon": "ios-globe-outline"},
		{"name": "mapper", "icon": "ios-infinite"},
		{"name": "do_receive_file", "icon": "ios-book-outline"},
		{"name": "do_response_receive_file", "icon": "ios-book-outline"},
		{"name": "file_read", "icon": "ios-book-outline"},
		{"name": "file_write", "icon": "ios-create-outline"},
		{"name": "file_delete", "icon": "ios-log-out"},
		{"name": "file_sync", "icon": "md-paper"},
		{"name": "href_parser", "icon": "ios-ionitron-outline"},
		{"name": "entity_parser", "icon": "ios-refresh-circle-outline"},
		{"name": "set_env", "icon": "ios-nuclear-outline"},
		{"name": "get_env", "icon": "ios-nuclear"},
		{"name": "cal_hash", "icon": "ios-flower-outline"},
		{"name": "run_cmd", "icon": "md-bonfire"},
		{"name": "sftp_upload", "icon": "md-arrow-up"},
		{"name": "ssh_shell", "icon": "ios-cloud-upload-outline"},
		{"name": "targz_uncompress", "icon": "ios-aperture"},
		{"name": "targz_compress", "icon": "ios-aperture-outline"},
		{"name": "ini_read", "icon": "ios-fastforward"},
		{"name": "ini_write", "icon": "ios-aperture-outline"},
		{"name": "base64_encode", "icon": "ios-magnet"},
		{"name": "base64_decode", "icon": "ios-magnet-outline"},
		{"name": "if", "icon": "md-code-working"},
		{"name": "elif", "icon": "md-code-working"},
		{"name": "else", "icon": "md-code-working"},
		{"name": "define_var", "icon": "md-hammer"},
		{"name": "assign_var", "icon": "md-hammer"},
		{"name": "map", "icon": "md-hammer"},
		{"name": "foreach", "icon": "md-hammer"},
		{"name": "panic_error", "icon": "md-hammer"},
		{"name": "template", "icon": "md-hammer"},
		{"name": "create_jwt", "icon": "md-hammer"},
		{"name": "parse_jwt", "icon": "md-hammer"},
		{"name": "do_error_filter", "icon": "md-hammer"},
		{"name": "send_mail", "icon": "md-hammer"},
	}
}

func RegistNodes() {
	vs := []interface{}{
		sql.SQLExecuteNode{},
		sql.SQLQueryNode{},
		sql.DBParserNode{},
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
		mail.SendMailNode{},
	}
	for _, v := range vs {
		t := reflect.ValueOf(v).Type()
		typeMap[strings.ToUpper(t.Name())] = t
	}
	node.Regist(typeMap)
}
