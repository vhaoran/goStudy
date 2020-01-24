package cachedemo

import (
	"fmt"
	"github.com/goburrow/cache"
	"math/rand"
	"testing"
	"time"
)

func Test_lru_2_go(t *testing.T) {
	load := func(k cache.Key) (cache.Value, error) {
		return fmt.Sprintf("%d", k), nil
	}
	// Create a new cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterAccess(10*time.Second),
		cache.WithRefreshAfterWrite(60*time.Second),
	)

	getTicker := time.Tick(10 * time.Millisecond)
	reportTicker := time.Tick(1 * time.Second)
	for {
		select {
		case <-getTicker:
			_, _ = c.Get(rand.Intn(2000))
		case <-reportTicker:
			st := cache.Stats{}
			c.Stats(&st)
			fmt.Printf("%+v\n", st)
		}
	}
}

func Test_bench_lru_guava(t *testing.T) {
	load := func(k cache.Key) (cache.Value, error) {
		return fmt.Sprintf("%d", k), nil
	}
	// Create a new cache
	l := cache.NewLoadingCache(load,
		cache.WithMaximumSize(10000),
		cache.WithExpireAfterAccess(10*time.Second),
		cache.WithRefreshAfterWrite(60*time.Second),
	)

	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i%10000)
		l.Put(key, 1)
	}

	fmt.Println("", time.Since(t0))
	//
	t0 = time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", 1)
		l.Get(key)
	}
	fmt.Println("", time.Since(t0))
}
