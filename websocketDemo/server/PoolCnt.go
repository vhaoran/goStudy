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
	BUS_LEN = 2000000
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
		fmt.Println(" .......loop bus wait.....len(Bus)..", len(r.Bus))
		select {
		case data, ok := <-r.Bus:
			if ok {
				fmt.Println(" BUS read", string(data), " len:", len(r.Bus))

				if err := r.dispatch(data); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("message has leave bus-------")
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

func (r *PoolCnt) Push(key string, conn *PoolUnit) {
	fmt.Println("Push,id:", key)

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
	fmt.Println("PoolCnt-->dispatch")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("dispatch err:", err)
		}
	}()

	if len(buffer) == 0 {
		return errors.New("no data to send")
	}
	//
	bean := new(MsgData)
	if err := json.Unmarshal(buffer, bean); err != nil {
		return err
	}
	//
	fmt.Println("PoolCnt-->dispatch,before  11111111")
	_, dst := r.getSrcDst(bean)
	//
	if obj := r.GetUnit(dst); obj != nil {
		fmt.Println("PoolCnt-->dispatch,before  22222")
		if s, err := json.Marshal(bean); err == nil {
			fmt.Println("PoolCnt-->dispatch,before  33333")
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
