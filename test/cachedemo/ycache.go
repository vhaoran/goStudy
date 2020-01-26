package cachedemo

import (
	"fmt"
	cache "github.com/patrickmn/go-cache"
	"time"
)

type CacheX struct {
	cache    *cache.Cache
	callback func() (interface{}, error)
}

func NewCacheX(defaultExpiration, cleanupInterval time.Duration, loadCallback func() (interface{}, error)) *CacheX {
	bean := &CacheX{
		cache:    cache.New(defaultExpiration, cleanupInterval),
		callback: loadCallback,
	}
	return bean
}

func (r *CacheX) Set(k, v interface{}) {
	key := fmt.Sprint(k)
	r.cache.SetDefault(key, v)
}

func (r *CacheX) Get(k, v interface{}) (interface{}, bool) {
	key := fmt.Sprint(k)

	return r.cache.Get(key)
}
