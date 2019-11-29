package yredis

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	host := "127.0.0.1"
	url := []string{
		fmt.Sprintf("%s:7001", host),
		fmt.Sprintf("%s:7002", host),
		fmt.Sprintf("%s:7003", host),
		fmt.Sprintf("%s:7004", host),
		fmt.Sprintf("%s:7005", host),
		fmt.Sprintf("%s:7006", host),
	}

	red, err := NewRedisClient(url)
	if err != nil {
		log.Println(err)
		return
	}

	//
	h := 10000
	var wg sync.WaitGroup
	wg.Add(h)
	t0 := time.Now()
	for i := 0; i < h; i++ {
		go func(k int) {
			if err := red.Set(fmt.Sprint("a_", k), k, time.Hour*100).Err(); err != nil {
				log.Println(err)
			} else {
				if i%100 == 0 {
					//log.Println("OK", i)
				}
			}
			wg.Done()
		}(i)
	}
	//
	wg.Wait()
	fmt.Println("---------aaa---count--", h, " time:", time.Since(t0))
}
