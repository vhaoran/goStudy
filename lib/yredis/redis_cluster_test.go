package yredis

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	url := []string{
		"192.168.0.99:7001",
		"192.168.0.99:7002",
		"192.168.0.99:7003",
		"192.168.0.99:7004",
		"192.168.0.99:7005",
		"192.168.0.99:7006",
	}

	red, err := NewRedisClient(url)
	if err != nil {
		log.Println(err)
		return
	}

	//
	t0 := time.Now()
	for i := 0; i < 10000; i++ {
		if err := red.Set(fmt.Sprint("a_", i), i, time.Hour*100).Err(); err != nil {
			log.Println(err)
		} else {
			if i%100 == 0 {
				log.Println("OK", i)
			}
		}
	}
	//

	//
	fmt.Println("---------aaa-----", time.Since(t0))
}
