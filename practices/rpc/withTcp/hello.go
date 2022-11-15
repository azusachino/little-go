package withTcp

import "net/rpc"

// 简单的Go RPC withTcp
type HelloService struct {
}

// Hello()方法必须满足Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (h *HelloService) Hello(req string, reply *string) error {
	*reply = "Hello: " + req
	return nil
}

// 重构

// 1. service name
const HelloServiceName = "github.com/azusachino/little-go/practices/rpc/withTcp.HelloService"

// 2. service interface
type HelloServiceInterface interface {
	// implementation
	Hello(req string, res *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, srv)
}
