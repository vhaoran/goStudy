package test

import (
	"fmt"
	"testing"
)

type (
	RedisClient struct {
	}

	ClassB struct {
		RedisClient
	}
)

func NewClassA() RedisClient {
	b := RedisClient{}
	return b
}

func (r *RedisClient) Exec() {
	fmt.Println("A is called")
}

func (r *ClassB) Exec() {
	fmt.Println("B is called")
	r.RedisClient.Exec()
}

func Test_A_B_class(t *testing.T) {
	new(ClassB).Exec()
}
