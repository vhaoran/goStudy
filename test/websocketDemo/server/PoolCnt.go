package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type (
	PoolCnt struct {
		sync.Mutex
		m map[string]*websocket.Conn
	}
)

func (r *PoolCnt) GetCnt(key string) *websocket.Conn {
	r.Lock()
	defer r.Unlock()
	if conn, ok := r.m[key]; ok {
		return conn
	}
	return nil
}

func (r *PoolCnt) Push(key string, conn *websocket.Conn) {
	r.Lock()
	defer r.Unlock()
	r.m[key] = conn
}

func (r *PoolCnt) Exist(key string) bool {
	r.Lock()
	defer r.Unlock()
	_, ok := r.m[key]
	return ok
}
