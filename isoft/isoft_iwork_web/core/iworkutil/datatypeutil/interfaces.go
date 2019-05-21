package datatypeutil

import "reflect"

func InterfaceConvertToSlice(data interface{}) []interface{} {
	sli := make([]interface{}, 0)
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			sli = append(sli, v.Index(i).Interface())
		}
	}
	return sli
}
