package goroutine

import (
	"fmt"
	"time"
)

func doBadThing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func doGoodThing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}

func timeOut(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)

	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)

	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
