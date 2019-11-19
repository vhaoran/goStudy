package test

import (
	"container/ring"
	"fmt"
	"log"
	"testing"
)

func Test_ring_1(t *testing.T) {
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()
	log.Println("len:", n)

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		//r.Prev()
		r = r.Next()
	}

	for i := 1; i < 3; i++ {
		r.Value = i * 300
		//r.Prev()
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		if p != nil {
			fmt.Println("p:", p)
		}
	})
}
