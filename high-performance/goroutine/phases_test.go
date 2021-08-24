package goroutine

import (
	"runtime"
	"testing"
	"time"
)

func TestTwoPhasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = timeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log(runtime.NumGoroutine())
}
