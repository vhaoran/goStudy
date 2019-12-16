package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"goStudy/g"
)

type MsgData struct {
	Dst  string
	Data string
	Src  string
}

func main() {
	//id := flag.String("id", "whr", "user id")
	host := flag.String("host", "127.0.0.1", "host name")
	flag.Parse()

	h := 50
	obj := g.NewWaitGroupN(h)
	t0 := time.Now()

	for i := 0; i < h; i++ {
		id := fmt.Sprint(i, "_whr")
		obj.Call(func() error {
			send(id, *host)
			return nil
		})
	}
	obj.Wait()
	offset := time.Since(t0)
	time.Sleep(5 * time.Second)
	fmt.Println("------", "", "-----------")
	fmt.Println("------", "", "-----------")
	fmt.Println("BBBB-sec:", offset)
	log.Println("BBBB-count:", h*100)
}

func send(id, host string) {
	src := fmt.Sprint("[", id, "]")
	fmt.Println("------[", id, "]-----------")

	wsUrl := url.URL{Scheme: "ws", Host: host + ":9999", Path: "/chat"}
	header := &http.Header{}
	header.Set("ID", id)

	conn, _, err := (&websocket.Dialer{}).Dial(wsUrl.String(), *header)
	if err != nil {
		fmt.Printf("conn create err:%s\n", err.Error())
		return
	}
	// defer conn.Close()
	go func() {
		for {
			_, msgData, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("*********conn read err:%s\n", err.Error())
				continue
			} else {
				fmt.Println(time.Now(), "### readOK ###", string(msgData))
			}
		}

	}()

	fmt.Println("conn success", src)
	t0 := time.Now()
	i := int64(0)
	for j := 0; j < 10000; j++ {
		//id := fmt.Sprint(i%500, "_whr")
		bean := &MsgData{
			Dst:  id,
			Data: "aaaaaa",
			Src:  "aaaa",
		}
		ss, _ := json.Marshal(bean)

		err = conn.WriteMessage(websocket.TextMessage, []byte(ss))
		if err != nil {
			fmt.Printf("send err:%s\n", err.Error())
			return
		} else {
			fmt.Println("send ok:", j)
		}

		i++

	}
	offset := time.Since(t0)
	if offset.Seconds() > 0 {
		log.Println("--", float64(i)/offset.Seconds(), "count:", i, " second:", offset, " avg:")
	} else {
		log.Println("count: ", i, " second: ", time.Since(t0))
	}

	time.Sleep(time.Second * 50)

}
