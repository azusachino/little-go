package main

import (
	"fmt"
	"time"
)

func main() {
	chanDemo()
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func worker(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func bufferedChannel() {
	c := make(chan int, 3) // buffer size = 3

	c <- 1
}

func channelClose() {
	c := make(chan int)

	c <- 1
	close(c)
}
