package reflectutil

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

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
