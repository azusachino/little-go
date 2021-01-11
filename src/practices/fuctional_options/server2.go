package fuctional_options

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Server2 struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	//...
	return nil, nil
}

func test() {

	//Using the default configuratrion
	srv1, _ := NewServer("localhost", 9000, nil)
	conf := Config{Protocol: "tcp", Timeout: 60}
	srv2, _ := NewServer("locahost", 9000, &conf)
	print(srv1, srv2)
}
