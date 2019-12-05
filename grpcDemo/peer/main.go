package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "goStudy/grpcDemo/pb"
)

const (
	address     = "localhost:8888"
	defaultName = "whr"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	msg := ":"
	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		name = fmt.Sprint("whr_", i)
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			break
		}
		msg = r.Message
	}
	log.Println("all times:", time.Since(t0))
	log.Printf("Greeting: %s", msg)
}
