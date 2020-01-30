package cachedemo

import (
	"fmt"
	"github.com/goburrow/cache"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_lru_2_go(t *testing.T) {
	load := func(k cache.Key) (cache.Value, error) {
		fmt.Println("load---->", k)
		return fmt.Sprintf("%d", k), nil
	}
	// Create a new cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterAccess(5*time.Second),
		cache.WithRefreshAfterWrite(5*time.Second),
	)

	for i := 0; i < 20; i++ {
		c.Put(i, i)
	}

	getTicker := time.Tick(500 * time.Millisecond)
	reportTicker := time.Tick(1 * time.Second)
	for {
		select {
		case <-getTicker:
			vv, er1 := c.Get(rand.Intn(5))
			fmt.Println(" #### get key:", vv, " ", er1)
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

func Test_rand_n(t *testing.T) {
	for i := 0; i < 100; i++ {
		log.Println("----------", rand.Intn(50000), "------------")
		log.Println("----####-", RandN6(), "------------")
	}
}

func RandN6() int32 {
	//rand.Seed(123456789)
REDO:
	j := rand.Int31n(123456789)
	for j < 100000 {
		j += rand.Int31n(123456789)
	}
	//
	if j > 100000 {
		j = j % 1000000
	}
	if j < 100000 {
		goto REDO
	}

	return j
}
