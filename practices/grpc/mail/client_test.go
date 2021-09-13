package main

import (
	"context"
	pb "github.com/little-go/practices/grpc/proto/mail"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestService_SendMail(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//addr := fmt.Sprintf("%s://%s", "tcp", "127.0.0.1:8972")
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8972", grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	c := pb.NewMailServiceClient(conn)

	res, err := c.SendMail(context.TODO(), &pb.MailRequest{
		Mail: "abc@protonmail.com",
		Text: "Hello",
	})
	log.Println(res)
}
