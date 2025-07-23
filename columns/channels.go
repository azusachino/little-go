package columns

import (
	"fmt"
	"time"
)

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func Test_chan() {
	ch := getIntChan()
	for ele := range ch {
		fmt.Printf("%v\n", ele)
	}
}

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch
		fmt.Printf("Worker %d\n", id)
		time.Sleep(1 * time.Second)
		nextCh <- token
	}
}

func Test_chan2() {
	num := 5
	chs := make([]chan Token, num)
	for i := 0; i < num; i++ {
		chs[i] = make(chan Token, 1)
	}
	for i := 0; i < num; i++ {
		go newWorker(i, chs[i], chs[(i+1)%num])
	}
	chs[0] <- Token{}
	select {}
}

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (mu *Mutex) Lock() {
	<-mu.ch
}

func (mu *Mutex) Unlock() {
	select {
	case mu.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (mu *Mutex) TryLock() bool {
	select {
	case <-mu.ch:
		return true
	default:
		return false
	}
}

func (mu *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-mu.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

func (mu *Mutex) IsLocked() bool {
	return len(mu.ch) == 0
}
