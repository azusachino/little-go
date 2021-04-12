package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var littleBuf = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 64)
		return &buf
	},
}
var goroutineSpace = []byte("goroutine")

func main() {
	go func() {
		fmt.Println("the goroutineId:", curGoRoutineId())
	}()

	time.Sleep(time.Second)
}

func curGoRoutineId() uint64 {
	bp := littleBuf.Get().(*[]byte)
	defer littleBuf.Put(bp)

	b := *bp
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No sapce found in %q", b))
	}
	b = b[:i]
	n, err := parseUintBytes(b, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func parseUintBytes(bytes []byte, a, b int) (uint64, error) {
	return 0, nil
}
