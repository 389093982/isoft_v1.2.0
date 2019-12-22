package test

import (
	"fmt"
	"reflect"
)

func TestReflect() {

	var str string
	str = "123456"
	fmt.Printf("type:%v and value:%v\n", reflect.TypeOf(str), reflect.ValueOf(str))

	v := DuanYan(false)
	fmt.Printf("v:%v\n", v)

	reflectTest(32)
}

func reflectTest(a interface{}) {

	//aType := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)
	aV := aValue.Interface()
	fmt.Printf("aV type:%T,value:%v\n", aV, aV)

	aV2 := aValue.Int()
	fmt.Printf("aV2 type:%T,value:%v\n", aV2, aV2)

}

func DuanYan(value interface{}) string {
	result := ""
	if v, ok := value.(string); ok {
		result = fmt.Sprintf("%v", v)
	}
	return result
}
