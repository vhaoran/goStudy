package main

import (
	"fmt"

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

func (r *PoolUnit) Loop(conn *websocket.Conn, id string) {
	r.loopDispatch(conn, id)
	r.loopReceive(conn, id)
}

func (r *PoolUnit) loopDispatch(conn *websocket.Conn, id string) {
	for {
		select {
		case data, ok := <-r.out:
			if ok {
				if err := conn.WriteMessage(0, []byte(data)); err != nil {
					return
				}
			} else {
				return
			}
		}
	}
}

func (r *PoolUnit) Dispatch(data []byte) {
	r.out <- data
}

func (r *PoolUnit) loopReceive(conn *websocket.Conn, id string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	for {
		_, msgData, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("**********conn read err:%s\n", err.Error())
			return
		}
		// log.Println("unit received from: ", id, "---data->", string(msgData))

		r.bus <- msgData
	}
}
