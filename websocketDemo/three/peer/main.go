package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"time"

	"go.uber.org/atomic"
	"golang.org/x/net/websocket"
)

var allCount atomic.Int32

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	allCount.Store(0)

	//host := flag.String("host", "127.0.0.1", "host pos")
	host := flag.String("host", "39.99.169.246", "host pos")
	flag.Parse()

	h := 100
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < h; i++ {
		go call(*host)
	}

	wg.Wait()
}

func call(host string) {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error
	origin := fmt.Sprintf("http://%s/", host)
	server := fmt.Sprintf("ws://%s:9999/ws", host)

	fmt.Println("------", "aaa", "-----------")
	fmt.Println(origin, server)
	fmt.Println("------", "aaa", "-----------")

	cfg := new(websocket.Config)
	if cfg.Location, err = url.ParseRequestURI(server); err != nil {
		return
	}
	if cfg.Origin, err = url.ParseRequestURI(origin); err != nil {
		return
	}
	cfg.Header = http.Header{"ID": []string{"whr"}}
	cfg.Version = websocket.ProtocolVersionHybi13

	ws, err := websocket.DialConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		h := 1000000
		t0 := time.Now()
		for i := 0; i < h; i++ {
			var msg = make([]byte, 512)
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Println(err, "*****")
				wg.Done()
				break
			}
			allCount.Inc()

			fmt.Println(fmt.Sprint(allCount.Load(), "   ", time.Now(), " Received: ", string(msg[:n])))
		}
		fmt.Println("total:", time.Since(t0))
	}()

	for i := 0; i < 10000; i++ {
		s := fmt.Sprint("hello,world--", i)
		if _, err := ws.Write([]byte(s)); err != nil {
			log.Fatal(err)
		}
		allCount.Inc()
		fmt.Println(fmt.Sprint(allCount.Load(), "   ", time.Now(), "##############"))
	}

	//ws.Close()
	log.Println("closed")
	log.Println("exited")
	wg.Wait()
}
