package cachedemo

import (
	"fmt"
	mcache "github.com/OrlovEvgeny/go-mcache"
	"log"
	"testing"
	"time"
)

type User struct {
	Name string
	Age  uint
	Bio  string
}

func Test_mcache_test_go(t *testing.T) {
	//Start mcache instance
	MCache := mcache.New()

	//Create custom key
	key := "custom_key1"
	//Create example struct
	user := &User{
		Name: "John",
		Age:  20,
		Bio:  "gopher 80 lvl",
	}

	//args - key, &value, ttl (or you need never delete, set ttl is mcache.TTL_FOREVER)
	err := MCache.Set(key, user, time.Minute*20)
	if err != nil {
		log.Fatal(err)
	}

	if data, ok := MCache.Get(key); ok {
		objUser := data.(*User)
		fmt.Printf("User name: %s, Age: %d, Bio: %s\n", objUser.Name, objUser.Age, objUser.Bio)
	}
}

func Test_bench_(t *testing.T) {
	l := mcache.New()
	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i%10000)
		l.Set(key, i, expired)
	}

	fmt.Println("", time.Since(t0))
	//
	t0 = time.Now()
	for i := 0; i < h; i++ {
		key := fmt.Sprint("key_", i)
		_, _ = l.Get(key)
	}
	fmt.Println("", time.Since(t0))
}
