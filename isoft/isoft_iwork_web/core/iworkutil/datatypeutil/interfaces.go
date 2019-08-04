package datatypeutil

import "reflect"

func InterfaceConvertToSlice(data interface{}) []interface{} {
	sli := make([]interface{}, 0)
	if data == nil {
		return sli
	}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			sli = append(sli, v.Index(i).Interface())
		}
	} else {
		sli = append(sli, data)
	}
	return sli
}
