package reflectDemo

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reflect_3333(t *testing.T) {
	s := make([]int, 0, 4)
	s = append(s, 1, 2, 3)
	AppendSlice(&s, 4)

	fmt.Println(s) //  [1 2 3]  如何修改AppendSlice 函数让s 返回 [1,2,3,4]
}

func AppendSlice(i, e interface{}) {
	rv := reflect.ValueOf(i)
	rv = rv.Elem()
	rv = reflect.Append(rv, reflect.ValueOf(e))
	rt := reflect.ValueOf(i)
	rt.Elem().Set(rv)
}
