package imageValidate

import (
	"fmt"
	"log"
)

type MyStore struct {
	m map[string][]byte
}

func NewMyStore() *MyStore {
	return &MyStore{
		m: make(map[string][]byte),
	}
}

func (r *MyStore) Set(id string, digits []byte) {
	log.Println("set:", id, ":", digits)
	r.m[id] = digits
}

func (r *MyStore) Get(id string, clear bool) (digits []byte) {
	digits, _ = r.m[id]
	fmt.Println("get digits", digits)
	return
}
