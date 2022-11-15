package main

import "github.com/azusachino/little-go/practices/pipeline"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go func() {
		for _, a := range arr {
			c <- a
		}
	}()
	print(pipeline.Odd(c))
}
