package test

import (
	"fmt"
	"strconv"
	"testing"
)

func init() {

}

func Test_float(t *testing.T) {
	s := "100.243"
	f, err := strconv.ParseFloat(s,10)
	fmt.Println("r:",f," err:",err)

}
