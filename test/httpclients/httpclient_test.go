package httpclients

import (
	"github.com/monaco-io/request"
	"log"
	"testing"
)

func Test_http_test(t *testing.T) {
	var body = struct {
		A string
		B int
	}{A: "A", B: 001}
	var result interface{}

	client := request.Client{
		URL:    "https://google.com",
		Method: "POST",
		Query:  map[string]string{"hello": "world"},
		JSON:   body,
	}
	resp := client.Send().Scan(&result)
	if !resp.OK() {
		// handle error
		log.Println(resp.Error())
	}

	// str := resp.String()
	// bytes := resp.Bytes()
}
