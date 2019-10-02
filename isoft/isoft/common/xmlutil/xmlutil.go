package xmlutil

import "encoding/xml"

func RenderToString(v interface{}) (s string) {
	bytes, err := xml.MarshalIndent(v, "", "\t")
	if err == nil {
		s = string(bytes)
	}
	return
}
