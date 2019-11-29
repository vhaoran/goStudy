package yredis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient(urls []string) (*redis.ClusterClient, error) {
	cnt := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              urls,
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

	if err := cnt.Ping().Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cnt, nil
}
