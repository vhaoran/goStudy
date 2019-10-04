package test

import (
	"fmt"
	"strings"
	"testing"
)

func In(a interface{}, l ...interface{}) bool {
	if l == nil || len(l) == 0 {
		return false
	}

	for _, v := range l {
		if a == v {
			return true
		}
	}

	return false
}

func Test_in(t *testing.T) {
	if In(7, 2, int64(3), int64(5), 7) {
		fmt.Println(" true")
	}
	fmt.Println(" end ")
}

func Test_f(t *testing.T) {

	f := float64(500.012345678)
	fmt.Print(fmt.Sprintf("%.2f", f))

}

func Test_strings_replacef(t *testing.T) {
	msg := "上分500000"
	src := "上分"
	s := strings.Replace(msg, src, "", -1)
	fmt.Print(fmt.Println("->", s))
}
