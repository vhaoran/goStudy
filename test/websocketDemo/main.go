package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	log.SetPrefix("wsDemo: ")
	http.HandleFunc("/", Home)
	http.Handle("/echo", websocket.Handler(Echo))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Echo(ws *websocket.Conn) {
	defer ws.Close()
	buf := make([]byte, 64*1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			return
		}
		_, err = ws.Write(buf[:n])
		if err != nil {
			return
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	body := strings.NewReader(home)
	http.ServeContent(w, r, "index.html", time.Time{}, body)
}
