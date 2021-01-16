package cachedemo

import (
	"fmt"
	"github.com/bluele/gcache"
	"testing"
	"time"
)

func Test_gcache_1(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.Set("key", "ok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}

func Test_gcache_3(t *testing.T) {
	gc := gcache.New(2000).
		LRU().
		Build()

	t0 := time.Now()
	z := 0
	for j := 0; j < 1000; j++ {
		for i := 0; i < 10000; i++ {
			k := fmt.Sprint("key_", i)
			_ = gc.Set(k, i)
			ret, _ := gc.Get(k)
			z += ret.(int)
		}
	}
	x := time.Now().Sub(t0).Milliseconds()
	fmt.Println("-----ok---------------ms: ", x)
	fmt.Println("-----ok---------------z: ", z)
}

func Test_gc_2(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.SetWithExpire("key", "ok", time.Second*1000)
	value, _ := gc.Get("key")
	fmt.Println("Get:", value)

	// Wait for value to expire
	time.Sleep(time.Second * 10)
	{
		value, err := gc.Get("key")
		if err != nil {
			panic(err)
		}
		fmt.Println("Get:", value) //
	}
}
