package helper

import "reflect"

// 如果src为空，则取默认值des
func SetDefault[T StringInteger](src, des T) T {
	v := reflect.ValueOf(src)
	switch v.Kind() {
	case reflect.String:
		if v.Interface().(string) == "" {
			return des
		}
	case reflect.Uint8:
		if v.Interface().(uint8) == 0 {
			return des
		}
	case reflect.Uint16:
		if v.Interface().(uint16) == 0 {
			return des
		}
	case reflect.Uint:
		if v.Interface().(uint) == 0 {
			return des
		}
	case reflect.Uint32:
		if v.Interface().(uint32) == 0 {
			return des
		}
	case reflect.Uint64:
		if v.Interface().(uint64) == 0 {
			return des
		}
	case reflect.Int8:
		if v.Interface().(int8) == 0 {
			return des
		}
	case reflect.Int16:
		if v.Interface().(int16) == 0 {
			return des
		}
	case reflect.Int:
		if v.Interface().(int) == 0 {
			return des
		}
	case reflect.Int32:
		if v.Interface().(int32) == 0 {
			return des
		}
	case reflect.Int64:
		if v.Interface().(int64) == 0 {
			return des
		}
	}

	return src
}
