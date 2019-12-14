package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"goStudy/g"
)

func main() {
	//id := flag.String("id", "whr", "user id")
	host := flag.String("host", "192.168.0.99", "host name")
	flag.Parse()

	obj := g.NewWaitGroupN(1000)
	t0 := time.Now()
	h := 100
	for i := 0; i < h; i++ {
		id := fmt.Sprint(i, "_whr")
		obj.Call(func() error {
			send(id, *host)
			return nil
		})
	}
	obj.Wait()
	fmt.Println("sec:", time.Since(t0))
	log.Println("count:", h*10000)
}

func send(id, host string) {
	src := fmt.Sprint("[", id, "]")
	fmt.Println("------[", id, "]-----------")

	wsUrl := url.URL{Scheme: "ws", Host: host + ":9999", Path: "/chat"}
	header := &http.Header{}
	header.Set("ID", id)

	conn, _, err := (&websocket.Dialer{}).Dial(wsUrl.String(), *header)
	if err != nil {
		fmt.Printf("conn create err:%s\n", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("conn success", src)
	t0 := time.Now()
	i := int64(0)
	for j := 0; j < 10000; j++ {
		err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint(i, "--hello_"+src)))
		if err != nil {
			fmt.Printf("send err:%s\n", err.Error())
			return
		}

		i++
		fmt.Println("send: hello, server", src)
		offset := time.Since(t0)
		if offset.Seconds() > 0 {
			log.Println("--", float64(i)/offset.Seconds(), "count:", i, " second:", offset, " avg:")
		} else {
			log.Println("count: ", i, " second: ", time.Since(t0))
		}

		//_, msgData, err := conn.ReadMessage()
		//if err != nil {
		//	fmt.Printf("conn read err:%s\n", err.Error())
		//	return
		//}
		//fmt.Printf("recv: %s\n", string(msgData))
		//time.Sleep(time.Second * 1)
	}
}
