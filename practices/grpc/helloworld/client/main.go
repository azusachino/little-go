package main

import (
	"context"
	pb "github.com/little-go/practices/grpc/helloworld/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// client call easily
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("call grpc failed: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
