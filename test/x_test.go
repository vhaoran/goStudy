package test

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

type Abc struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
	C string `json:"c,omitempty"`
}

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

func Test_abc_marshal(t *testing.T) {
	bean := Abc{
		A: "1",
		B: "2",
		C: "3",
	}
	s, err := json.Marshal(bean)
	if err != nil {
		log.Println("----------", err, "------------")
		return
	}

	log.Println("----------", string(s), "------------")
}

func Test_test_map_unmarshal(t *testing.T) {
	m := map[string]interface{}{"a": "1", "b": "2", "c": "3"}
	//
	s, err := json.Marshal(m)
	if err != nil {
		fmt.Println("x_test.go->", err)
		return
	}

	//
	fmt.Println("-----")
	fmt.Println(string(s))
	//-------- -----------------------------
	m1 := &Abc{}
	//
	err = json.Unmarshal(s, &m1)
	if err != nil {
		log.Println("----------", "------------")
		fmt.Println(err)
		return
	}

	log.Println("----------", "unmarshal", "------------")
	spew.Dump(m1)

}

func Test_map_unmarshal_2(t *testing.T) {
	m := map[string]interface{}{"a": "1", "b": "2", "c": "3"}
	//
	s, err := json.Marshal(m)
	if err != nil {
		fmt.Println("x_test.go->", err)
		return
	}

	//
	fmt.Println("-----")
	fmt.Println(string(s))
	//-------- -----------------------------
	m1 := new(map[string]interface{})
	//
	err = json.Unmarshal(s, &m1)
	if err != nil {
		log.Println("----------", "------------")
		fmt.Println(err)
		return
	}

	log.Println("----------", "unmarshal", "------------")
	spew.Dump(m1)
}
