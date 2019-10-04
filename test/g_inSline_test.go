package test

import (
	"fmt"
	"testing"

	"goStudy/lib/g"
)

func Test_slice_in(t *testing.T) {
	a := g.InSlice(1, []int{11, 2, 3, 1})
	fmt.Println(a)
	fmt.Println("-----------------")
	a = g.InSlice("a", []string{"b", "a", "d"})
	fmt.Println(a)
	fmt.Println("-----------------")
	i1, i2, i3, i4 := 5, 2, 3, 4
	fmt.Println(g.InSlice(int(2), []*int{&i1, &i2, &i3, &i4}))
}

func Test_a(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["a"] = 1
	m["a"] = 1
	m["a"] = 1

}
