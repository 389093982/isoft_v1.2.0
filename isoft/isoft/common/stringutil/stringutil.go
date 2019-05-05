package stringutil

import (
	"github.com/satori/go.uuid"
	"reflect"
	"regexp"
)

func RandomUUID() string {
	return uuid.NewV4().String()
}

func GetTypeOfInterface(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func CheckIndexContains(s string, slice []string) (bool, int) {
	if len(slice) == 0 {
		return false, -1
	}
	for index, ss := range slice {
		if ss == s {
			return true, index
		}
	}
	return false, -1
}

func CheckContains(s string, slice []string) bool {
	b, _ := CheckIndexContains(s, slice)
	return b
}

func ChangeStringsToInterfaces(ss []string) []interface{} {
	result := make([]interface{}, 0)
	for _, s := range ss {
		result = append(result, s)
	}
	return result
}

func GetSubStringWithRegexp(s, regex string) []string {
	reg := regexp.MustCompile(regex)
	return reg.FindAllString(s, -1)
}

func GetNoRepeatSubStringWithRegexp(s, regex string) []string {
	ss := GetSubStringWithRegexp(s, regex)
	return RemoveRepeatForSlice(ss)
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepeatForSlice(slc []string) []string {
	result := make([]string, 0)
	// 存放不重复主键
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			// 加入map后,map长度变化,则元素不重复
			result = append(result, e)
		}
	}
	return result
}

func RemoveItemFromSlice(slc []string, item string) []string {
	result := make([]string, 0)
	for _, s := range slc {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}
