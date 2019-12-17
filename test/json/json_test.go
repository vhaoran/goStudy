package json

import (
	"fmt"
	"testing"
	"time"

	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func Test_aaa(t *testing.T) {
	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		value := gjson.Get(json, "name.last")
		println(value.String())
	}
	fmt.Println("---", time.Since(t0))
}
