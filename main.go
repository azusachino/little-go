package main

import (
	"context"
	"fmt"
	"github.com/little-go/learn-project/ants"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(ants.Running())
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			goto label
		default:
			_ = ants.Submit(func() {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
			})
		}
	}
label:

	time.Sleep(time.Second * 3)
	fmt.Println(ants.Running())
	ants.Release()
	fmt.Println(ants.Running())
}
