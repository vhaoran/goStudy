package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type (
	MsgData struct {
		Dst  string
		Data string
		Src  string
	}

	PoolUnit struct {
		bus chan<- []byte
		out chan []byte
	}
)

func (r *MsgData) Str() string {
	return fmt.Sprint("dst:", r.Dst, "data:", r.Data)
}

func NewPoolUnit(bus chan<- []byte) *PoolUnit {
	bean := &PoolUnit{
		out: make(chan []byte, 100),
		bus: bus,
	}
	return bean
}

func (r *PoolUnit) loopOut(conn *websocket.Conn, id string) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err:", err)
			}
		}()

		for {
			select {
			case data := <-r.out:
				{
					fmt.Printf("recv: %s\n", string(data))
					conn.WriteMessage(1, data)
				}
			}
		}
	}()
}

func (r *PoolUnit) Loop(conn *websocket.Conn, id string) {
	r.loopOut(conn, id)
	r.loopIn(conn, id)
}

func (r *PoolUnit) loopIn(conn *websocket.Conn, id string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("**********conn read err:%s\n", err.Error())
			continue
		}
		log.Println("unit received from: ", id, "---data->", msgType, " dadta:", string(msgData))
		r.bus <- msgData
	}
}

func (r *PoolUnit) Cast(buffer []byte) error {
	r.out <- buffer
	return nil
}
