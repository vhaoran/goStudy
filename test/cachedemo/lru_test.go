package cachedemo

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"testing"
	"time"
)

func Test_lru_1(t *testing.T) {
	l, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
}

func Test_bench_lru(t *testing.T) {
	h := 10000
	l, _ := lru.New(h)
	t0 := time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i%10000)
		l.Add(key, i)
	}

	fmt.Println("", time.Since(t0))
	//
	t0 = time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i)
		l.Contains(key)
	}
	fmt.Println("", time.Since(t0))
}
