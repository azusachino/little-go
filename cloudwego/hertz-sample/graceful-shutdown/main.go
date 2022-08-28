package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("172.0.0.1:8090"),
		server.WithExitWaitTime(3*time.Second))

	// add shutdown hook
	h.OnShutdown = append(h.OnShutdown,
		func(ctx context.Context) {
			fmt.Println("shutdown hook")
			<-ctx.Done()
			fmt.Println("exit timeout!")
		})

    h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
        fmt.Println("another hook")
    })

    h.Spin()
}
