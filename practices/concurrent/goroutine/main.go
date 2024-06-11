package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//demo()
	//timeAfter()
	timeAfterFor()
	// timeAfterForFixed()
}

// normal case (no loop)
func timeAfter() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "haha"
	}()
	select {
	case x := <-ch:
		fmt.Println(x)
	case <-time.After(time.Second):
		fmt.Print("timeout!!!")
	}
}

// memory up
func timeAfterFor() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			in := 1
			for {
				in++
				ch <- in
			}
		}()
	}

	for {
		select {
		case aaa := <-ch:
			fmt.Println(aaa)
		case <-time.After(3 * time.Second):
			fmt.Printf("now is: %d, do something!!!", time.Now().Unix())
		case <-time.After(100 * time.Second):
			return
		}
	}
}

func timeAfterForFixed() {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	ch := make(chan int, 10)
	go func() {
		in := 1
		for {
			in++
			ch <- in
		}
	}()

	for {
		select {
		case <-ch:
		case <-timer.C:
			fmt.Printf("now is %d, do something!!!", time.Now().Unix())
		}
	}
}

// 在 Go1.14 前，不会输出任何结果。
// 在 Go1.14 及之后，能够正常输出结果。
func demo() {
	// set to 1 process
	runtime.GOMAXPROCS(1)

	// Go1.14 实现了基于信号的抢占式调度，以此来解决上述一些仍然无法被抢占解决的场景。
	go func() {
		// with infinite loop
		for {

		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("mission clear")
}
