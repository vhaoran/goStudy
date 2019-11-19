package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
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
			"192.168.0.99:7001",
			"192.168.0.99:7002",
			"192.168.0.99:7003",
			"192.168.0.99:7004",
			"192.168.0.99:7005",
			"192.168.0.99:7006"},
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
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           2,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	cnt.Ping()
	//
	h := 1000000
	var wg sync.WaitGroup
	wg.Add(h)

	t0 := time.Now()
	for i := 0; i < h; i++ {
		go func(k int, wg *sync.WaitGroup) {
			defer wg.Done()
			stat := cnt.Set(fmt.Sprint("a_", k), k, time.Hour*100)
			if stat.Err() != nil {
				fmt.Println("-----------------", stat.Err())
			} else {
				fmt.Println("ok->", k)
			}
		}(i, &wg)
	}
	//
	wg.Wait()
	//
	fmt.Println("---------aaa-----", time.Since(t0))
}
