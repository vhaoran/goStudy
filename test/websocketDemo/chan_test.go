package websocketDemo

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"

	"goStudy/g"
)

type ABC struct {
	sync.RWMutex
	m map[string]string
}

func NewABC() *ABC {
	bean := &ABC{
		m: make(map[string]string),
	}
	return bean
}

func (r *ABC) Push(key, v string) {
	r.Lock()
	defer r.Unlock()
	r.m[key] = v
}

func Test_chan_test(t *testing.T) {
	h := int64(1000 * 10000)
	c := make(chan string, h)
	t0 := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		doExit := false
		for {
			select {
			case <-c:
			case <-time.After(30 * time.Millisecond):
				doExit = true
			}
			if doExit {
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := int64(0); i < h; i++ {
			c <- fmt.Sprint("c", i)
		}
	}()

	wg.Wait()
	log.Println("since:", time.Since(t0))
}

func Test_map_high(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	h := 1000 * 10000
	bean := NewABC()
	t0 := time.Now()
	obj := g.NewWaitGroupN(1000)
	for i := 0; i < h; i++ {
		obj.Call(func() error {
			bean.Push(fmt.Sprint(i), fmt.Sprint(i))
			return nil
		})
	}
	obj.Wait()
	fmt.Println("since:", time.Since(t0))
}

func Test_m_high(t *testing.T) {
	h := 1000 * 10000
	bean := NewABC()
	t0 := time.Now()
	for i := 0; i < h; i++ {
		bean.Push(fmt.Sprint(i), fmt.Sprint(i))
	}
	fmt.Println("since:", time.Since(t0))
}
