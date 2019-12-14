package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	wsUrl := url.URL{Scheme: "ws", Host: "127.0.0.1:9999", Path: "/chat"}
	conn, _, err := (&websocket.Dialer{}).Dial(wsUrl.String(), nil)
	if err != nil {
		fmt.Printf("conn create err:%s\n", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("conn succ")
	for {
		err = conn.WriteMessage(websocket.TextMessage, []byte("hello, server"))
		if err != nil {
			fmt.Printf("send err:%s\n", err.Error())
			return
		}
		fmt.Println("send: hello, server")

		_, msgdata, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read err:%s\n", err.Error())
			return
		}
		fmt.Printf("recv: %s\n", string(msgdata))

		time.Sleep(time.Second * 5)
	}
}
