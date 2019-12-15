package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func chat(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{}).Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		fmt.Printf("conn creat err:%s\n", err.Error())
		return
	}
	defer conn.Close()
	for {
		msgtype, msgdata, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read err:%s\n", err.Error())
			return
		}
		fmt.Printf("recv: %s\n", string(msgdata))
		conn.WriteMessage(msgtype, []byte("hello, client"))
		fmt.Println("send: hello, client")

		time.Sleep(time.Second * 5)
	}
}

func main() {
	fmt.Println("start ws... ...")
	http.HandleFunc("/chat", chat)
	http.ListenAndServe(":9999", nil)
}
