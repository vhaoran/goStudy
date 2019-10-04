package test

import (
	"fmt"
	"testing"
)

func Test_aa(t *testing.T) {
	k := int32(8008888)
	fmt.Println(k)
}
func Test_map_r(t *testing.T) {
	fmt.Println(ret_map())
}

func ret_map() (m map[string]string) {
	m = make(map[string]string)
	m["a"] = "c"
	m["b"] = "b_v"
	m["c"] = "c_v"
	for k := range m {
		fmt.Println(k)
	}

	return
}
