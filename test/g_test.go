package test

import (
	"fmt"
	"testing"

	"goStudy/lib/g"
)

func Test_is_zero(t *testing.T) {
	i := 0
	fmt.Println(g.IsZero(&i))
}

func Test_isZeroAll(t *testing.T) {
	fmt.Println(g.IsZeroAll())
}
