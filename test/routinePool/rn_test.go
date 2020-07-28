package routinePool

import (
	"fmt"
	"testing"
	"time"
)

func Test_rn_n(t *testing.T) {

	for i := 0; i < 100000000; i++ {
		go func() {
			j := i
			fmt.Println("### ", j)
		}()
	}

	time.Sleep(5 * time.Second)
}
