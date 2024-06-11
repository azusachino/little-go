package examples

import (
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
	t.Log("example finished")
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
