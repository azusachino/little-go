package main

import (
	"context"
	"fmt"
	pb "github.com/azusachino/little-go/practices/grpc/proto/mail"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type service struct {
	// fallback solution
	pb.UnimplementedMailServiceServer
}

func (s *service) SendMail(_ context.Context, req *pb.MailRequest) (res *pb.MailResponse, err error) {
	fmt.Printf("mail: %s, msg detail: %s", req.Mail, req.Text)
	return &pb.MailResponse{
		Ok: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterMailServiceServer(s, &service{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
