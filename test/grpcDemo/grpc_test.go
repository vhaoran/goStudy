package grpcDemo

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
	"time"
)

type (
	Person struct {
		ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
		Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	}
)

var high = 100000

func (p *Person) Reset() {
	p = &Person{}
}

func (p *Person) String() string {
	return proto.CompactTextString(p)
}

func (p Person) ProtoMessage() {}

func Test_proto_buffer(t *testing.T) {
	t0 := time.Now()
	for i := 0; i < high; i++ {
		bean := &Person{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprint("aa", i),
		}
		if buffer, err := proto.Marshal(bean); err == nil {
			p := new(Person)
			proto.Unmarshal(buffer, p)
		} else {
			fmt.Println("error:", err)
		}
	}
	fmt.Println("all time:", time.Since(t0))
}

func Test_json(t *testing.T) {
	t0 := time.Now()
	for i := 0; i < high; i++ {
		bean := &Person{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprint("aa", i),
		}
		if buffer, err := json.Marshal(bean); err == nil {
			p := new(Person)
			json.Unmarshal(buffer, p)
		} else {
			fmt.Println("error:", err)
		}
	}
	fmt.Println("all time:", time.Since(t0))

}

func Test_empty_noAction(t *testing.T) {
	t0 := time.Now()
	j := ""
	for i := 0; i < high; i++ {
		bean := &Person{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprint("aa", i),
		}
		j = bean.ID
	}

	fmt.Println("all time:", time.Since(t0), " // ", j)

}
