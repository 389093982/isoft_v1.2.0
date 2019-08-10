package stringutil

import (
	"github.com/satori/go.uuid"
	"reflect"
	"regexp"
	"strings"
)

func RandomUUID() string {
	return uuid.NewV4().String()
}

func GetTypeOfInterface(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func CheckIndexContains(s string, slice []string) (int, bool) {
	return CheckIgnoreCaseIndexContains(s, slice, false)
}

func CheckIgnoreCaseIndexContains(s string, slice []string, ignoreCase bool) (int, bool) {
	if len(slice) == 0 {
		return -1, false
	}
	for index, ss := range slice {
		if (!ignoreCase && ss == s) || (ignoreCase && strings.ToUpper(ss) == strings.ToUpper(s)) {
			return index, true
		}
	}
	return -1, false
}

func CheckIgnoreCaseContains(s string, slice []string) bool {
	_, b := CheckIgnoreCaseIndexContains(s, slice, true)
	return b
}

func CheckContains(s string, slice []string) bool {
	_, b := CheckIgnoreCaseIndexContains(s, slice, false)
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
