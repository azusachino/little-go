package main

import (
	"fmt"
	"time"
)

func main() {
	A()
}

func Routine() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("Hello from Go Routine %d \n", i)
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func A() {
	for {
		fmt.Println("infinite loop " + time.Now().String())
	}
}
