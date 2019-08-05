package datatypeutil

import (
	"reflect"
	"strconv"
)

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

func InterfaceConvertToInt(v interface{}, defaultVal int) int {
	if value, ok := v.(int); ok {
		return value
	} else if value, ok := v.(string); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultVal
}
