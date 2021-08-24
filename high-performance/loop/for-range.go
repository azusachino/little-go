package main

import "fmt"

func main() {
	chn()
}

func chn() {
	ch := make(chan string)
	go func() {
		ch <- "Hello"
		ch <- "World"
		ch <- "the truth"
		close(ch)
	}()

	for s := range ch {
		fmt.Println(s)
	}
}
