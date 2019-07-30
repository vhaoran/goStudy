package test

import (
	"fmt"
	"reflect"
	"testing"
)

func init() {

}

func Test_TypeOf_pointer(t *testing.T) {
	a := 3

	tp := reflect.TypeOf(a)

	fmt.Println(tp)
	fmt.Println(tp.Name())

	// is Ptr
	if tp.Kind() == reflect.Ptr {
		fmt.Println("is pointer")
	} else {
		fmt.Println("not a pointer")
	}

	fmt.Println("tp.Kind():", tp.Kind())
	if tp.Kind() == reflect.Int {
		fmt.Println("is int")
		v := reflect.ValueOf(a)
		//v_int := v.Interface().(int)
		v_int := v.Int()
		fmt.Println("v_int:", v_int)
	} else {
		fmt.Println("not a int")
	}

	//v := reflect.ValueOf(a)
}

func Test_(t *testing.T) {


}
