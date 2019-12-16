package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	var err error
	origin := "http://localhost/"
	server := "ws://localhost:9999/ws"

	cfg := websocket.Config{
	}
	if cfg.Location, err = url.ParseRequestURI(server); err != nil {
		return
	}
	if cfg.Origin, err = url.ParseRequestURI(origin); err != nil {
		return
	}
	cfg.Header = http.Header{"ID": []string{"whr"}}
	cfg.Version = websocket.ProtocolVersionHybi13

	ws, err := websocket.DialConfig(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			var msg = make([]byte, 512)
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Println(err, "*****")
				break
			}
			fmt.Printf("Received: %s.\n", msg[:n])
		}
	}()

	for i := 0; i < 3; i++ {
		s := fmt.Sprint("hello,world--", i)
		if _, err := ws.Write([]byte(s)); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * 5)
	}
	ws.Close()
	log.Println("closed")
	time.Sleep(time.Second * 3)
	log.Println("exited")
}
