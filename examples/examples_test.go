package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	c := Counter{
		sync.Mutex{},
		0,
	}
	if c.TryLock() {
		c.Lock()
		defer c.Unlock()
	}
	fmt.Printf("finished %d", c.Count)
}

func TestTryMutex(t *testing.T) {
	var mu sync.Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		t.Log("got the lock")
		mu.Unlock()
		return
	}
	t.Log("cannot get the lock")
}

func TestCount(t *testing.T) {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("%d, %t, %t, %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
