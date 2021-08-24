package goroutine

import (
	"runtime"
	"testing"
	"time"
)

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		_ = timeOut(f)
	}
	time.Sleep(2 * time.Second)
	t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) {
	test(t, doBadThing)
}

func TestBufferTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = timeoutWithBuffer(doBadThing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}
