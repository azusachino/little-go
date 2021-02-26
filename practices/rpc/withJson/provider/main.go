package main

import (
	"github.com/little-go/practices/rpc/withJson"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 将HelloService注册为RPC服务
	_ = rpc.RegisterName("HelloService", new(withJson.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept Error: ", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
