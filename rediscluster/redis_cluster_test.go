package rediscluster

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func ExampleNewClusterClient() {
	// See http://redis.io/topics/cluster-tutorial
	// how to setup Redis Cluster.
	cnt := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"172.2.2.11:7001",
			"172.2.2.12:7002",
			"172.2.2.13:7003",
			"172.2.2.14:7004",
			"172.2.2.15:7005",
			"172.2.2.16:7006"},
	})

	cnt.Ping()
	//
	t0 := time.Now()
	for i := 0; i < 100000; i++ {
		cnt.Set(fmt.Sprint("a_", i), i, time.Hour*100)
	}
	//

	//
	fmt.Println("---------aaa-----", time.Since(t0))
}

func Test_a(t *testing.T) {
	ExampleNewClusterClient()
}
