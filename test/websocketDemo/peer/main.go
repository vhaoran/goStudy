package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	id := flag.String("id", "whr", "user id")
	flag.Parse()
	src := fmt.Sprint("[", *id, "]")
	fmt.Println("------[", *id, "]-----------")

	wsUrl := url.URL{Scheme: "ws", Host: "127.0.0.1:9999", Path: "/chat"}
	header := &http.Header{}
	header.Set("id", *id)

	conn, _, err := (&websocket.Dialer{}).Dial(wsUrl.String(), *header)
	if err != nil {
		fmt.Printf("conn create err:%s\n", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("conn success", src)
	for {
		err = conn.WriteMessage(websocket.TextMessage, []byte("hello, server"+src))
		if err != nil {
			fmt.Printf("send err:%s\n", err.Error())
			return
		}

		fmt.Println("send: hello, server", src)

		_, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read err:%s\n", err.Error())
			return
		}
		fmt.Printf("recv: %s\n", string(msgData))
		time.Sleep(time.Second * 5)
	}
}
