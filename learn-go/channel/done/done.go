package main

import "fmt"

func main() {

}

func worker(id int, c chan int, ok chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		ok <- true
	}
}
