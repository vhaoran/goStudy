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

func NewPoolUnit(bus chan []byte) *PoolUnit {
	bean := &PoolUnit{
		out: make(chan []byte, 10000),
		bus: bus,
	}
	return bean
}

func (r *PoolUnit) Loop(conn *websocket.Conn, id string) {
	go r.loopDispatch(conn, id)
	r.loopReceive(conn, id)
}

func (r *PoolUnit) loopDispatch(conn *websocket.Conn, id string) {
	fmt.Println(id, "---loopDispatch----")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("loopDispatch err", err)
		}
	}()
	for {
		select {
		case data, ok := <-r.out:
			if ok {
				if err := conn.WriteMessage(1, data); err != nil {
					fmt.Println("loopDispatch to peer err ", err)
					return
				} else {
					fmt.Println("loopDispatch to peer ok ")
				}
			} else {
				return
			}
		}
	}
}

func (r *PoolUnit) Dispatch(data []byte) error {
	r.out <- data
	return nil
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
		fmt.Println("unit received from of  ID: ", id, "---data->", string(msgData), "len(bus):", len(r.bus))

		r.bus <- msgData
	}
}
