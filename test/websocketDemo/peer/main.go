package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	id := flag.String("id", "whr", "user id")
	host := flag.String("host", "192.168.0.99", "host name")
	flag.Parse()
	src := fmt.Sprint("[", *id, "]")
	fmt.Println("------[", *id, "]-----------")

	wsUrl := url.URL{Scheme: "ws", Host: *host + ":9999", Path: "/chat"}
	header := &http.Header{}
	header.Set("ID", *id)

	conn, _, err := (&websocket.Dialer{}).Dial(wsUrl.String(), *header)
	if err != nil {
		fmt.Printf("conn create err:%s\n", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("conn success", src)
	t0 := time.Now()
	i := int64(0)
	for {

		err = conn.WriteMessage(websocket.TextMessage, []byte("hello, server"+src))
		if err != nil {
			fmt.Printf("send err:%s\n", err.Error())
			return
		}

		i++
		fmt.Println("send: hello, server", src)
		offset := time.Since(t0)
		if offset.Seconds() > 0 {
			log.Println("count:", i, " second:", offset, " avg:", float64(i)/offset.Seconds())
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
