package reflectDemo

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_tm_m(t *testing.T) {
	a := []int{1, 2, 3}
	s, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("------", "", "-----------")
	log.Println(string(s))
}

func Test__r3_aaa(t *testing.T) {
	s := "abc"
	vType(1, s)
	fmt.Println("------", "", "-----------")
	vType(2, []byte(s))
	//
	vType(2, []string{"a", "b"})
}

func vType(k int, v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println(k, "is string", vv)
	case int:
		fmt.Println(k, "is int ", vv)
	case float64:
		fmt.Println(k, "is float64 ", vv)
	case []byte:
		fmt.Println(k, "is byte array:")
	case []interface{}:
		fmt.Println(k, "is array:")
		for i, j := range vv {
			fmt.Println(i, j)
		}
	}

}
