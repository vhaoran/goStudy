package merge

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/imdario/mergo"
	"log"
	"testing"
)

type Foo struct {
	A string
	B int64
}

type FooA struct {
	A int64
	B int64
	C string
}

func Test_a_merge(t *testing.T) {
	src := Foo{
		A: "one",
		B: 2,
	}
	dest := FooA{
		A: 1111,
		C: "good",
	}
	mergo.Merge(&dest, src)
	fmt.Println(dest)
	log.Println("----------", "aa", "------------")
	spew.Dump("dest", dest)
}

func Test_a_merge_2(t *testing.T) {
	src := map[string]interface{}{"A": 5, "B": 4, "C": "good"}
	dest := FooA{
		A: 1111,
		C: "good",
	}
	if err := mergo.Map(&dest, src, mergo.WithOverride); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dest)
	log.Println("----------", "aa", "------------")
	spew.Dump("dest", dest)
}
