package test

import (
	"fmt"
	"testing"
)

func init() {

}

func IsIntX(a interface{}) bool {
	switch a.(type) {
	case int:
		{
			return true
		}
	}
	return false
}


func Test_int(t *testing.T) {
	a8 := int(8)
	a16 := int16(8)
	a32 := int32(8)
	a64 := int64(8)
	fmt.Println("",IsIntX(a8))
	fmt.Println("",IsIntX(a16))
	fmt.Println("",IsIntX(a32))
	fmt.Println("",IsIntX(a64))



}
