package test

import (
	"fmt"
	"testing"
)

type AAA struct {
	ID   int
	Name string
	Age  int
}

func Test_v(t *testing.T) {
	a := AAA{
		ID:   1,
		Name: "aaaa",
		Age:  3333,
	}

	fmt.Println("------ +v -----------")
	fmt.Println(fmt.Sprintf("%+v", a))
	fmt.Println("----- #v ------------")
	fmt.Println(fmt.Sprintf("%#v", a))
	fmt.Println("-----------v------")
	fmt.Println(fmt.Sprintf("%v", a))
}
