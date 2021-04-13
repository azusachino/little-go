package main

import (
	"context"
	pb "github.com/little-go/practices/grpc/helloworld/proto"
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
	// new GRPC SERVER
	s := grpc.NewServer()
	// register our service
	pb.RegisterGreeterServer(s, new(server))
	// listen call
	if err := s.Serve(tcp); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
