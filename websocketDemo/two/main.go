package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.Handle("/ws", websocket.Handler(chat))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func chat(ws *websocket.Conn) {
	{
		h := ws.Request().Header
		fmt.Println("------", "", "-----------")
		for k, v := range h {
			log.Println(k, v)
		}
	}

	obj := new(SendUnit)
	go obj.Send(ws, nil)

	for {
		var reply string
		buffer := make([]byte, 1024)
		fmt.Println(time.Now(), "......block before Read")
		if n, err := ws.Read(buffer); err != nil {
			log.Println(err, "***********************")
			break
		} else {
			reply = string(buffer[:n])
			log.Println("received:", reply)
		}

		new(SendUnit).Send(ws, []byte(reply))

		//if err = websocket.Message.Receive(ws, &reply); err != nil {
		//	fmt.Println(err)
		//	continue
		//}

		//ret := strings.ToUpper(reply) + "-----[l'mfrom server]"
		//fmt.Println(time.Now(), "#### block before Send")
		//if err = websocket.Message.Send(ws, ret); err != nil {
		//	fmt.Println("send to client:", err)
		//	//continue
		//}
	}
}
