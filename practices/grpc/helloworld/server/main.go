package main

import (
	"context"
	pb "github.com/azusachino/golong/practices/grpc/helloworld/proto"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	httpReporter "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello" + in.GetName()}, nil
}

func main() {
	// simulate start service
	tcp, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	reporter := httpReporter.NewReporter("http://localhost:9411")

	tracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		log.Fatalf("failed to start zipkin: %v", err)
	}
	// new GRPC SERVER
	s := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))
	// register our service
	pb.RegisterGreeterServer(s, new(server))
	// listen call
	if err := s.Serve(tcp); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
