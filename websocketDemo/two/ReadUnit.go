package main

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

type (
	SendUnit struct {
	}
)

func (r *SendUnit) Send(conn *websocket.Conn, buffer []byte) {
	var err error

	if len(buffer) == 0 {
		t0 := time.Now()
		h := 1000000
		for i := 0; i < h; i++ {
			s := fmt.Sprint(i, "----from server!")
			//for {
			if err = websocket.Message.Send(conn, []byte(s)); err != nil {
				fmt.Println("send to client:", err)
				break
			}
		}
		fmt.Println("total", time.Since(t0), " ", t0)

	} else {
		if err = websocket.Message.Send(conn, buffer); err != nil {
			fmt.Println("send to client:", err)
		} else {
			fmt.Println("send to client ok")
		}
	}
}
