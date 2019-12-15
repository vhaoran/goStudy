package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/atomic"
	"log"
	"sync"
	"time"
)

type (
	PoolCnt struct {
		Bus chan []byte

		sync.Mutex
		m map[string]*PoolUnit

		threadMin   atomic.Int32
		threadMax   atomic.Int32
		threadCount atomic.Int32
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
	bean.threadMin.Store(1000)
	bean.threadMax.Store(10000)
	bean.threadCount.Store(0)

	go bean.loop()
	return bean
}

func (r *PoolCnt) loop() {
	for {
		r.autoAdd()
		time.Sleep(time.Second * 1)
	}
}

func (r *PoolCnt) autoAdd() {
	cur := int32(len(r.Bus))
	count := r.threadCount.Load()
	//
	if count < r.threadMin.Load() || (count < r.threadMax.Load() && cur > (count)*5) {
		for i := 0; i < 100; i++ {
			r.addOne()
		}
	}
}

func (r *PoolCnt) addOne() {
	fmt.Println("#####  ", r.threadCount.Load(), "   ####")

	if r.threadCount.Load() >= r.threadMax.Load() {
		return
	}

	r.threadCount.Inc()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		r.threadCount.Dec()
	}()

	exit := false
	for {
		fmt.Println(" .......loop wait.......")
		select {
		case data, ok := <-r.Bus:
			if ok {
				if err := r.dispatch(data); err != nil {
					log.Println(err)
				}
				//r.autoAdd()
			}
		case <-time.After(time.Second * 60):
			exit = true
		}

		if exit {
			break
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
			return obj.Cast(s)
		} else {
			return err
		}
	}
	//if local not exists key,then brocast to other node

	//save offline
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
