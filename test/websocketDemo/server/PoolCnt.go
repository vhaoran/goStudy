package main

import (
	"log"
	"sync"
)

type (
	PoolCnt struct {
		Bus chan []byte

		sync.Mutex
		m map[string]*PoolUnit
	}
)

const (
	BUS_LEN = 1000000
)

func NewPoolCnt() *PoolCnt {
	bean := &PoolCnt{
		Bus: make(chan []byte, BUS_LEN),
		m:   make(map[string]*PoolUnit),
	}
	go bean.loop()
	return bean
}

func (r *PoolCnt) loop() {
	for {
		select {
		case buffer, ok := <-r.Bus:
			if ok {
				log.Println(" bus received======= ", string(buffer))
			}
		}
	}
}

func (r *PoolCnt) GetCnt(key string) *PoolUnit {
	r.Lock()
	defer r.Unlock()
	if conn, ok := r.m[key]; ok {
		return conn
	}
	return nil
}

func (r *PoolCnt) Push(key string, conn *PoolUnit) {
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
