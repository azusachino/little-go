package main

import "fmt"

func worker(ch chan struct{}) {
	// wait channel produce, block
	<-ch
	fmt.Println("do something")
	close(ch)
}

func main() {
	s := make(Set)
	s.Add("Tom")
	s.Add("Sam")
	fmt.Println(s.Has("Tom"))
	fmt.Println(s.Has("Jack"))

	ch := make(chan struct{})
	go worker(ch)

	ch <- struct{}{}
}
