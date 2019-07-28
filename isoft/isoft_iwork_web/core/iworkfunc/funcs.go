package iworkfunc

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
	"path/filepath"
	"strconv"
	"strings"
)

type IWorkFuncProxy struct {
}

func (this *IWorkFuncProxy) StringsEq(args []interface{}) interface{} {
	return args[0].(string) == args[1].(string)
}

func (this *IWorkFuncProxy) StringsContains(args []interface{}) interface{} {
	return strings.Contains(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) StringsHasSuffix(args []interface{}) interface{} {
	return strings.HasSuffix(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) StringsHasPrefix(args []interface{}) interface{} {
	return strings.HasPrefix(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) StringsToLower(args []interface{}) interface{} {
	return strings.ToLower(args[0].(string))
}

func (this *IWorkFuncProxy) StringsToUpper(args []interface{}) interface{} {
	return strings.ToUpper(args[0].(string))
}

func (this *IWorkFuncProxy) StringsJoin(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs, "")
}

func (this *IWorkFuncProxy) Int64Add(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount("Int64Add", sargs, 2)
	return sargs[0] + sargs[1]
}

func (this *IWorkFuncProxy) Int64Sub(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount("Int64Sub", sargs, 2)
	return sargs[0] - sargs[1]
}

func (this *IWorkFuncProxy) Int64Gt(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount("Int64Gt", sargs, 2)
	return sargs[0] > sargs[1]
}

func (this *IWorkFuncProxy) Int64Lt(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount("Int64Lt", sargs, 2)
	return sargs[0] < sargs[1]
}

func (this *IWorkFuncProxy) Int64Eq(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount("Int64Eq", sargs, 2)
	return sargs[0] == sargs[1]
}

func (this *IWorkFuncProxy) Int64Multi(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount("Int64Multi", sargs, 2)
	return sargs[0] * sargs[1]
}

func (this *IWorkFuncProxy) StringsJoinWithSep(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs[:len(args)-1], sargs[len(args)-1])
}

func (this *IWorkFuncProxy) Or(args []interface{}) interface{} {
	sargs := make([]bool, 0)
	for _, arg := range args {
		if arg == nil {
			sargs = append(sargs, false)
		} else {
			sargs = append(sargs, arg.(bool))
		}
	}
	return sargs[0] || sargs[1]
}

func (this *IWorkFuncProxy) And(args []interface{}) interface{} {
	sargs := make([]bool, 0)
	for _, arg := range args {
		if arg == nil {
			sargs = append(sargs, false)
		} else {
			sargs = append(sargs, arg.(bool))
		}
	}
	return sargs[0] && sargs[1]
}

func (this *IWorkFuncProxy) Not(args []interface{}) interface{} {
	sargs := make([]bool, 0)
	for _, arg := range args {
		if arg == nil {
			sargs = append(sargs, false)
		} else {
			sargs = append(sargs, arg.(bool))
		}
	}
	return !sargs[0]
}

func (this *IWorkFuncProxy) Uuid(args []interface{}) interface{} {
	return stringutil.RandomUUID()
}

func (this *IWorkFuncProxy) Isempty(args []interface{}) interface{} {
	if val, ok := args[0].(string); ok {
		return val == ""
	}
	return args[0] == nil
}

func (this *IWorkFuncProxy) GetDirPath(args []interface{}) string {
	return filepath.Dir(args[0].(string))
}

func (this *IWorkFuncProxy) IfThenElse(args []interface{}) interface{} {
	if args[0] == nil { // 参数为空条件为假
		return args[2]
	} else {
		if bol, ok := args[0].(bool); ok {
			if bol {
				return args[1] // bool 值且 true
			} else {
				return args[2] // bool 值且 false
			}
		} else { // 非空且不是 bool 值为真
			return args[1]
		}
	}
}

func checkArgsAmount(funcName string, sargs []int64, amount int) {
	if len(sargs) < amount {
		panic(errors.New(fmt.Sprintf(`%s 函数参数个数不足或者参数类型有误！`, funcName)))
	}
}

func parseArgsToInt64Arr(args []interface{}) []int64 {
	sargs := make([]int64, 0)
	for _, arg := range args {
		if _arg, ok := arg.(int64); ok {
			sargs = append(sargs, _arg)
		} else if _arg, ok := arg.(int); ok {
			sargs = append(sargs, int64(_arg))
		} else if _arg, ok := arg.(string); ok {
			if _arg, err := strconv.ParseInt(_arg, 10, 64); err == nil {
				sargs = append(sargs, _arg)
			}
		}
	}
	return sargs
}
