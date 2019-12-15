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
		bus  chan<- []byte
		conn *websocket.Conn
	}
)

func (r *MsgData) Str() string {
	return fmt.Sprint("dst:", r.Dst, "data:", r.Data)
}

func NewPoolUnit(bus chan<- []byte) *PoolUnit {
	bean := &PoolUnit{
		//conn: make(chan []byte, 100),
		bus: bus,
	}
	return bean
}

func (r *PoolUnit) Loop(conn *websocket.Conn, id string) {
	r.conn = conn
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

func (r *PoolUnit) Cast(buffer []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		r.conn.Close()
	}()

	err := r.conn.WriteMessage(1, buffer)
	if err != nil {
		fmt.Println("PoolUnit->Cast error:", err)
	} else {
		log.Println("PoolUnit->Cast ok", string(buffer))
	}

	return err
}
