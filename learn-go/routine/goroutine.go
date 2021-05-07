package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	otherPractice()
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
	// infinite loop
	for {
		fmt.Println("infinite loop " + time.Now().String())
	}
}

func bestPractice() {
	counter := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	ch := make(chan bool)

	for {
		select {
		case <-counter:
			fmt.Println("tick tock")
		case <-after:
			ch <- true
		}
	}

	done := <-ch

	if done {
		os.Exit(1)
	}

}

func otherPractice() {
	ch := make(chan string)

	// main process
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Check me out"
	}()

	select {
	case s := <-ch:
		fmt.Println(s)
	case <-time.After(time.Second * 4):
		fmt.Println("Should be done here")
	}
}
