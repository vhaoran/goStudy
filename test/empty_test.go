package test

import (
	"fmt"
	"reflect"
	"testing"
)

func init() {

}

func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	fmt.Println("kind:", v.Kind())
	//fmt.Println("kindOf:",reflect.KindOf(a))
	switch v.Kind() {
	case reflect.Invalid:
		{
			return true
		}
	case reflect.Ptr:
		{
			return v.IsNil()
		}
	case reflect.Int:
		{
			return v.Int() == 0
		}

	case reflect.String:
		{
			fmt.Println("v.string", v.String())
			return len(v.String()) <= 0
		}
	}

	return false
}

func Test_empty(t *testing.T) {
	var p *string
	p = nil

	fmt.Println("1:", IsEmpty(1))
	fmt.Println("0:", IsEmpty(0))
	fmt.Println("-1:", IsEmpty(1))
	fmt.Println("-----------------")
	fmt.Println("nil:", IsEmpty(nil))
	fmt.Println("-----------------")
	fmt.Println("point:", IsEmpty(p))
	fmt.Println("-----------------")
	fmt.Println("str :", IsEmpty(""))

}
