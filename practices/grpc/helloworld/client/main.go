package main

import (
	"context"
	pb "github.com/little-go/practices/grpc/helloworld/proto"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	httpReporter "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "World"
)

func main() {
	reporter := httpReporter.NewReporter("http://localhost:9411")
	tracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		log.Fatalf("failed to start zikpin: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithStatsHandler(zipkingrpc.NewServerHandler(tracer)), grpc.WithInsecure(), grpc.WithBlock())
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
