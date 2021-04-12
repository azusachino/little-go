package main

import (
	"fmt"
	"time"
)

func main() {
	demo3()
}

func demo1() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("2 second after")
}

func demo2() {
	v := make(chan struct{})
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("now do something")
		v <- struct{}{}
	})
	defer timer.Stop()
	// ensure send happens after receive
	<-v
}

func demo3() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	// background goroutine => send done signal
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	// main process
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-ticker.C:
			fmt.Println("hen shin: ", t.Unix())
		}
	}
}
