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
		Addrs: []string{"172.2.2.11:7001",
			"localhost:7001",
			//"172.2.2.12:7002",
			//"172.2.2.13:7003",
			//"172.2.2.14:7004",
			//"172.2.2.15:7005",
			//"172.2.2.16:7006"
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
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           10,
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
			x := cnt.Set(fmt.Sprint("a_", k), k, time.Hour*100)
			err := x.Err()
			if err != nil {
				fmt.Println(err)
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
