package cachedemo

import (
	"fmt"
	"github.com/muesli/cache2go"
	"testing"
	"time"
)

// Keys & values in cache2go can be of arbitrary types, e.g. a struct.
type myStruct struct {
	text     string
	moreData []byte
}

func Test_cache2go(t *testing.T) {
	// Accessing a new cache table for the first time will create it.
	cache := cache2go.Cache("myCache")

	// We will put a new item in the cache. It will expire after
	// not being accessed via Value(key) for more than 5 seconds.
	val := myStruct{"This is a test!", []byte{}}
	cache.Add("someKey", 5*time.Second, &val)

	// Let's retrieve the item from the cache.
	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Wait for the item to expire in cache.
	time.Sleep(6 * time.Second)
	res, err = cache.Value("someKey")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	// Add another item that never expires.
	cache.Add("someKey", 0, &val)

	// cache2go supports a few handy callbacks and loading mechanisms.
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache.
	cache.Delete("someKey")

	// And wipe the entire cache table.
	cache.Flush()
}

func Test_bench_cache2go(t *testing.T) {
	cache := cache2go.Cache("myCache")
	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key", i)
		cache.Add(key, 50*time.Second, i)
	}
	fmt.Println("", time.Since(t0))
	t0 = time.Now()
	k := 0
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key", i)
		_, err := cache.Value(key)
		if err == nil {
			k++
		}
	}
	fmt.Println("", time.Since(t0))

}
