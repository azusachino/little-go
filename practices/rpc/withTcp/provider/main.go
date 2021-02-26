package main

import (
	"github.com/little-go/practices/rpc/withTcp"
	"log"
	"net"
	"net/rpc"
)

// rpc.RegisterName()函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，所有注册的方法会放在HelloService服务的空间之下。
// 然后建立一个唯一的TCP链接，并且通过rpc.ServeConn()函数在该TCP链接上为对方提供RPC服务。

func main() {
	// 将HelloService注册为RPC服务
	_ = rpc.RegisterName("HelloService", new(withTcp.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept Error: ", err)
	}
	rpc.ServeConn(conn)
}
