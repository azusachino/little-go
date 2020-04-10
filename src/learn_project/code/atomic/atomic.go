package main

import (
	"fmt"
	"sync"
	"time"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicInt) increment() {
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *AtomicInt) get() int {

	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a AtomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
