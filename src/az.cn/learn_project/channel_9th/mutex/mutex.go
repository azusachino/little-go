package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInteger struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInteger) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInteger) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInteger
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.value)

}
