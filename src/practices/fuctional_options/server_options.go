package fuctional_options

import (
	"crypto/tls"
	"time"
)

// 1. 定义函数类型
type Option func(*Server)

// 2. 将形参都转成方法
func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

// 3. 定义构造方法
func NewServer_(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func main() {
	s1, _ := NewServer_("localhost", 1024)
	s2, _ := NewServer_("localhost", 2048, Protocol("udp"))
	s3, _ := NewServer_("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
	print(s1, s2, s3)
}
