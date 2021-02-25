package main

import "fmt"

func main() {
	intChan := make(chan int)
	boolChan := make(chan bool, 1)
	go worker(1, intChan, boolChan)
	select {
	case <-boolChan:
		fmt.Println("ha ha ha")
	default:
		fmt.Println("he he he")
	}
	<-intChan
	close(intChan)
	close(boolChan)
}

func worker(id int, c chan int, ok chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		ok <- true
	}
}
