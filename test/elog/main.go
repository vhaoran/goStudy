package main

import (
	"flag"
	"github.com/starjiang/elog"
	"log"
	"time"
)

func main() {
	flag.Parse()

	x := elog.NewEasyLogger("debug", false, 1,
		elog.NewEasyFileHandler("./logs", 100))

	defer x.Flush()

	h := 100 * 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		x.Info("hello", "world")
	}

	log.Println("time: ", time.Since(t0))
}
