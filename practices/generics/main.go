package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	//now := time.Now()
	for i := 0; i < 100000; i++ {
		go p(i)
	}
	for {
	}
}

func p(i interface{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println(i, time.Now())
}
