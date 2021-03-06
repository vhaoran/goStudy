package g

import (
	"errors"
	"reflect"
	"strings"
)

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmptyAll(l ...interface{}) bool {
	for _, v := range l {
		if !IsEmpty(v) {
			return false
		}
	}
	return true
}

func IsEmptyAllInfo(info string, l ...interface{}) error {
	for _, v := range l {
		if !IsEmpty(v) {
			return nil
		}
	}
	return errors.New(info)
}

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmptyOr(l ...interface{}) bool {
	for _, v := range l {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmptyOrInfo(info string, l ...interface{}) (err error) {
	all := strings.Split(info, "/")
	s := ""

	for i, v := range l {
		if IsEmpty(v) {
			if i >= 0 && i < len(all) {
				s += "/" + all[i]
			} else {
				if s == "" {
					s += "/不合法数据項"
				}
			}
		}
	}
	if len(s) > 0 {
		return errors.New(s)
	}
	return nil
}

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmpty(dst interface{}) bool {
	val := reflect.Indirect(reflect.ValueOf(dst))

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			return IsZero(val.Int())
		}
	case reflect.Float32, reflect.Float64:
		{
			return IsZero(val.Float())
		}
	case reflect.String:
		{
			return val.Len() == 0
		}
	case reflect.Map, reflect.Slice, reflect.Array:
		{ //len == 0
			if val.IsNil() {
				return true
			}
			return val.Len() == 0
		}
	case reflect.Chan:
		{ //len == 0
			if val.IsNil() {
				return true
			}
			return val.Len() == 0
		}
	case reflect.Struct:
		{ //len == 0
			if val.IsNil() {
				return true
			}
			return false
		}

	}
	return false
}
