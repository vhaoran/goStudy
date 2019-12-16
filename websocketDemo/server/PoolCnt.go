package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	for {
		fmt.Println(" .......loop wait.......")
		select {
		case data, ok := <-r.Bus:
			if ok {
				if err := r.dispatch(data); err != nil {
					log.Println(err)
				}
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

func (r *PoolCnt) PushUnit(key string, conn *PoolUnit) {
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

func (r *PoolCnt) dispatch(buffer []byte) error {
	if len(buffer) == 0 {
		return errors.New("no data to send")
	}
	//
	bean := new(MsgData)
	if err := json.Unmarshal(buffer, bean); err != nil {
		return err
	}
	//
	_, dst := r.getSrcDst(bean)
	//
	if obj := r.GetUnit(dst); obj != nil {
		if s, err := json.Marshal(bean); err != nil {
			return obj.Dispatch(s)
		} else {
			return err
		}
	}
	return nil
}

func (r *PoolCnt) getSrcDst(bean *MsgData) (src, dst string) {
	return bean.Src, bean.Dst
}

func (r *PoolCnt) GetUnit(key string) *PoolUnit {
	obj, ok := r.m[key]
	if !ok || obj == nil {
		return nil
	}
	return obj
}
