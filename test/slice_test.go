package test

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := l[0:3]
	fmt.Println("-----------------")
	fmt.Println(s)
	fmt.Println("-----------------")
	fmt.Println(l)
}
