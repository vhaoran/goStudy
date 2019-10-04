package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Msg struct {
	Type    int64
	Content interface{}
}
type ContentTest struct {
	ID   int64
	Text string
}

func Test_json_unmashal(t *testing.T) {
	msg := &Msg{
		Type: 1,
		Content: &ContentTest{
			ID:   1,
			Text: "333333,,sss",
		},
	}

	x := fmt.Sprintf("aaaaa %p", msg)
	fmt.Println(x)
	x = fmt.Sprintf("bbb %+v", msg)
	fmt.Println(x)

	fmt.Println("----------VVV-------")

	s, _ := json.Marshal(msg)
	fmt.Println("s:", string(s))
	fmt.Println("-----------------")
	//-------- -----------------------------
	msg2 := new(Msg)
	json.Unmarshal(s, msg2)

	fmt.Println(fmt.Sprintf("msg2 %v", *msg2))
	fmt.Println(fmt.Sprintf("msg2 %+v", *msg2))
	fmt.Println(fmt.Sprintf("msg2 %#v", *msg2))

}
