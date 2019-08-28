package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_time_cut(t *testing.T) {
	now := time.Now()

	i := now.Unix()
	//
	j := i % int64(86400)

	tt := now.Add(time.Second*time.Duration(j)*(-1))
	fmt.Println("------aa-----------")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(tt)
	fmt.Println(now)
}
