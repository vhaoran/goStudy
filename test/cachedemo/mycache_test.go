package cachedemo

import (
	"fmt"
	"testing"
	"time"
)

func Test_mycache(t *testing.T) {

	obj := NewCacheX(5*time.Second, 5*time.Second, nil)

	fmt.Println(obj)

}
