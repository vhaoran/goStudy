package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Test_str(t *testing.T) {
	b := []byte{'a', 'b', 'c'}
	p2b := b
	fmt.Printf("%p\n", p2b)
	fmt.Println("-----------------")
	s := BytesToString(b)
	fmt.Printf("%p\n", &s)

}
