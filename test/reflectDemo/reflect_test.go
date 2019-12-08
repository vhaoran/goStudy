package reflectDemo

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"reflect"
	"testing"
)

type Good struct {
	ID   int
	Name string
}

func Test_slice_test(t *testing.T) {
	l := make([]*Good, 0)
	err := DoSet(&l)
	log.Println("----------", "aaa", "------------")
	fmt.Println(err)
	spew.Dump(l)
}

func DoSet(a interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(a))
	ptr := reflect.ValueOf(a)
	//
	switch v.Kind() {
	case reflect.Slice:
		{
			t := v.Type().Elem().Elem()
			log.Print(t.String())

			for i := 0; i < 10; i++ {
				v = reflect.Append(v, reflect.New(t))
			}

			//spew.Dump(v.Interface())
			ptr.Elem().Set(v)
			return nil
		}
	default:
		return errors.New("no support data Type,only for slice")
	}
	//

	return nil
}
