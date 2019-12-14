package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var pool = NewPoolCnt()

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
	defer conn.Close()
	fmt.Println("------", "header", "-----------")
	for k, v := range req.Header {
		log.Println(k, v)
	}
	fmt.Println("------", "endHeader", "-----------")

	//-------- -----------------------------
	l, ok := req.Header["Id"]
	fmt.Println("------", "result", "-----------")
	log.Println(l, ok)

	if !ok || len(l) == 0 {
		log.Println("not login,ensure pass id in header")
		return
	}
	id := l[0]

	//-------------------------------------
	unit := NewPoolUnit(pool.Bus)
	pool.Push(id, unit)
	unit.Loop(conn, id)
}
