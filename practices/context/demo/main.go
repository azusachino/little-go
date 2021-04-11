package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	demo()
}

func demo() {
	parentCtx := context.TODO()
	ctx, cancel := context.WithTimeout(parentCtx, 1*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
