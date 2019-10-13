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

func InterfaceConvertToInt(v interface{}, defaultVal ...int64) int {
	return int(InterfaceConvertToInt64(v, defaultVal...))
}

func InterfaceConvertToInt64(v interface{}, defaultVal ...int64) int64 {
	if value, ok := v.(int); ok {
		return int64(value)
	} else if value, ok := v.(int64); ok {
		return value
	} else if value, ok := v.(string); ok {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultVal[0]
}
