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
			"192.168.0.99:7006",

			//"127.0.0.1:7001",
			//"127.0.0.1:7002",
			//"127.0.0.1:7003",
			//"127.0.0.1:7004",
			//"127.0.0.1:7005",
			//"127.0.0.1:7006",
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
		DialTimeout:        5 * time.Second,
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       5 * time.Second,
		PoolSize:           300,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        10 * time.Second,
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
			<<<<<<< HEAD
			stat := cnt.Set(fmt.Sprint("a_", k), k, time.Hour*100)
			if stat.Err() != nil {
				fmt.Println("-----------------", stat.Err())
			} else {
				fmt.Println("ok->", k)
				====== =
				x := cnt.Set(fmt.Sprint("a_", k), k, time.Hour*100)
				err := x.Err()
				if err != nil {
					fmt.Println(x)
				} else {
					fmt.Println("ok-->", k)
					>>>>>>> 307
					b400b457cc33950419e0dc830f3c34f19ad93
				}
			}
			(i, &wg)
		}
		//
		wg.Wait()
		//
		fmt.Println("---------aaa-----", time.Since(t0))
	}
