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
	}

	PoolUnit struct {
		in  chan *MsgData
		bus chan<- []byte
	}
)

func (r *MsgData) Str() string {
	return fmt.Sprint("dst:", r.Dst, "data:", r.Data)
}

func NewPoolUnit(bus chan<- []byte) *PoolUnit {
	bean := &PoolUnit{
		in:  make(chan *MsgData, 100),
		bus: bus,
	}
	return bean
}

func (r *PoolUnit) Loop(conn *websocket.Conn, id string) {
	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("**********conn read err:%s\n", err.Error())
			continue
		}
		log.Println("unit received from: ", id, "---data->", msgType, " dadta:", string(msgData))
		//fmt.Printf("recv: %s\n", string(msgData))
		//conn.WriteMessage(msgType, []byte("hello, client"))
		//fmt.Println("send: hello, client")
		//
		//time.Sleep(time.Second * 1)
		r.bus <- msgData
	}
}
