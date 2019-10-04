package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_sleep(t *testing.T) {
	ms := time.Duration(5000) * time.Millisecond
	t0 := time.Now().UnixNano()

	if time.Duration(time.Now().UnixNano()-t0) < ms {
		fmt.Println("sleep before")
		time.Sleep(ms)
		fmt.Println("sleep after")
	}

}

func Test_for(t *testing.T) {
	l := []string{"b", "b", "c"}
	for i, v := range l {
		l[i] = v + "_" + fmt.Sprint(i)
	}
	fmt.Println("------aa-----------")
	fmt.Println(l)

}
