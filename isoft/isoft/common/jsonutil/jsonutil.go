package jsonutil

import "encoding/json"

func RenderToJson(v interface{}) string {
	if bytes, err := json.MarshalIndent(v, "", "\t"); err == nil {
		return string(bytes)
	}
	return ""
}

func RenderToNoIndentJson(v interface{}) string {
	if bytes, err := json.MarshalIndent(v, "", ""); err == nil {
		return string(bytes)
	}
	return ""
}
