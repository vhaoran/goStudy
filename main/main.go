package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello.world!")
	ExampleNewClusterClient()
}

func ExampleNewClusterClient() {
	// See http://redis.io/topics/cluster-tutorial
	// how to setup Redis Cluster.
	cnt := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:7001",
			"127.0.0.1:7002",
			"127.0.0.1:7003",
			"127.0.0.1:7004",
			"127.0.0.1:7005",
			"127.0.0.1:7006",
		},
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		OnNewNode:          nil,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        50 * time.Second,
		ReadTimeout:        50 * time.Second,
		WriteTimeout:       50 * time.Second,
		PoolSize:           300,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        100 * time.Second,
		IdleTimeout:        500 * time.Second,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	er1 := cnt.Ping().Err()
	if er1 != nil {
		fmt.Println(er1)
		return
	}
	//
	h := 500000
	var wg sync.WaitGroup
	wg.Add(h)

	t0 := time.Now()
	for i := 0; i < h; i++ {
		go func(k int, wg *sync.WaitGroup) {
			defer wg.Done()
			x := cnt.Set(fmt.Sprint("a_", k), k, time.Hour*100)
			err := x.Err()
			if err != nil {
				fmt.Println(x)
			} else {
				fmt.Println("ok-->", k)
			}
		}(i, &wg)
	}
	//
	wg.Wait()
	//
	fmt.Println("---------aaa-----", time.Since(t0))
}
