package cachedemo

import (
	"fmt"
	cache "github.com/patrickmn/go-cache"
	"testing"
	"time"
	// "github.com/patrickmn/go-cache"
)

func Test_go_cache(t *testing.T) {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:
	foo, found = c.Get("foo")
	if found {
		fmt.Println(foo.(string))
	}

	// This gets tedious if the value is used several times in the same function.
	// You might do either of the following instead:
	if x, found := c.Get("foo"); found {
		fmt.Println(x.(string))
		// ...
	}
	// or

	if x, found := c.Get("foo"); found {
		foo = x.(string)
		fmt.Println(foo)
	}

}

const expired = time.Second * 100

func Test_bench(t *testing.T) {
	c := cache.New(5*time.Minute, 10*time.Minute)

	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i)
		c.Set(key, i, expired)
	}
	fmt.Println("set()", h, " ", time.Since(t0))
	k := 0
	t0 = time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i)
		_, ok := c.Get(key)
		if ok {
			k++
		}
	}
	fmt.Println("set()", h, " ", time.Since(t0))
}
