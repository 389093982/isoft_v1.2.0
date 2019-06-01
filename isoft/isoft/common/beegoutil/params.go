package beegoutil

import "reflect"

func GetInt64(proxy interface{}, key string, def int64) int64 {
	ref := reflect.ValueOf(proxy)
	method := ref.MethodByName("GetInt64")
	if !method.IsValid() {
		return def
	}
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(key))
	args = append(args, reflect.ValueOf(def))
	rs := method.Call(args)
	if rs[1].Interface() != nil {
		return def
	} else {
		return rs[0].Interface().(int64)
	}
}
