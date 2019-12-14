package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var pool = new(PoolCnt)

func main() {
	fmt.Println("start ws... ...")
	http.HandleFunc("/chat", chat)

	http.ListenAndServe(":9999", nil)
}

func chat(w http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		fmt.Printf("conn creat err:%s\n", err.Error())
		return
	}

	fmt.Println("------", "header", "-----------")
	for key, v := range req.Header {
		fmt.Println(key, v)
	}

	id, ok := req.Header["id"]
	if !ok {
		//return
	}

	log.Println("connect ok of id :", id)
	defer conn.Close()

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read err:%s\n", err.Error())
			return
		}

		fmt.Printf("recv: %s\n", string(msgData))
		conn.WriteMessage(msgType, []byte("hello, client"))
		fmt.Println("send: hello, client")

		time.Sleep(time.Second * 1)
	}
}
