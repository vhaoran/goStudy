package skiplistutil

import (
	"fmt"
	"github.com/sean-public/fast-skiplist"
	"testing"
	"time"
)

func Test_skip_list_1(t *testing.T) {
	list := skiplist.New()
	h := 1000
	c := h * 10000

	for i := 0; i < c; i++ {
		list.Set(float64(i), "This string data is stored at key 123!")
		//fmt.Println(list.Get(123).Value())
		//fmt.Println(list.Length) // prints 1
		//list.Remove(123)
		//fmt.Println(list.Length) // prints 0
	}

	start := time.Now()
	z := ""
	for i := 0; i < c; i++ {
		s := list.Get(float64(i)).Value().(string)
		z = s
	}
	fmt.Println("---", z, "----")
	offset := time.Now().Sub(start)
	fmt.Println("couht: ", c, "---ms:", offset)
}

func Test_map_1(t *testing.T) {
	list := make(map[int]string)
	h := 100
	c := h * 10000

	for i := 0; i < c; i++ {
		list[i] = "This string data is stored at key 123!"
	}

	start := time.Now()
	z := ""
	for i := 0; i < c; i++ {
		s, _ := list[i]
		z = s
	}
	fmt.Println("---", z, "----")
	offset := time.Now().Sub(start)
	fmt.Println("couht: ", c, "---ms:", offset)
}
