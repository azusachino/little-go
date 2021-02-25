package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func myHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	time.Sleep(10 * time.Second)
	cancel()
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

func main() {
	logger := log.New(os.Stdout, "[main] ", log.Ltime)
	myHandler()
	logger.Printf("finished")
}
