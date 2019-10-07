package reflectutil

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func InterfaceToSlice(v interface{}) (values []reflect.Value) {
	_v := reflect.ValueOf(v)
	for i := 0; i < _v.Len(); i++ {
		values = append(values, _v.Index(i))
	}
	return
}

func InterfaceToMap(v interface{}) (keys []reflect.Value, values []reflect.Value) {
	_v := reflect.ValueOf(v)
	keys = _v.MapKeys()
	for _, key := range keys {
		values = append(values, _v.MapIndex(key))
	}
	return
}

// 将结构体里的成员按照字段名字来赋值
func FillFieldValueToStruct(ptr interface{}, fields map[string]interface{}) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		if value, ok := fields[fieldInfo.Name]; ok {
			//给结构体赋值
			//保证赋值时数据类型一致
			if value == nil || reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			} else {
				panic(errors.New(fmt.Sprintf("mismatch fill field for struct, required %s, but get %s!",
					v.FieldByName(fieldInfo.Name).Type().String(), reflect.ValueOf(value).Type().String())))
			}
		}
	}
	return
}
