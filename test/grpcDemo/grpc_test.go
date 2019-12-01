package grpcDemo

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
)

type (
	Person struct {
		ID   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
		Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	}
)

var high = int32(1000000)

func (p *Person) Reset() {
	p = &Person{}
}

func (p *Person) String() string {
	return proto.CompactTextString(p)
}

func (p Person) ProtoMessage() {}

func Test_proto_buffer(t *testing.T) {
	t0 := time.Now()
	ii := 0
	for i := int32(0); i < high; i++ {
		bean := &Person{
			ID:   i,
			Name: fmt.Sprint("aa", i),
		}
		if buffer, err := proto.Marshal(bean); err == nil {
			ii = len(buffer)

			p := new(Person)
			proto.Unmarshal(buffer, p)
			//fmt.Println("len: ", len(buffer))
		} else {
			fmt.Println("error:", err)
		}
	}
	fmt.Println("all time:", time.Since(t0), " ", ii)
}

func Test_json(t *testing.T) {
	t0 := time.Now()
	ii := 0
	for i := int32(0); i < high; i++ {
		bean := &Person{
			ID:   i,
			Name: fmt.Sprint("aa", i),
		}
		if buffer, err := json.Marshal(bean); err == nil {
			ii = len(buffer)
			p := new(Person)
			json.Unmarshal(buffer, p)
		} else {
			fmt.Println("error:", err)
		}
	}
	fmt.Println("all time:", time.Since(t0), "  ", ii)

}

func Test_empty_noAction(t *testing.T) {
	t0 := time.Now()
	j := int32(0)
	for i := int32(0); i < high; i++ {
		bean := &Person{
			ID:   i,
			Name: fmt.Sprint("aa", i),
		}
		j = bean.ID
	}

	fmt.Println("all time:", time.Since(t0), " // ", j)
}
