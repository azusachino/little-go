package ants

import (
	"errors"
	"log"
	"math"
	"os"
	"runtime"
	"time"
)

const (
	// 默认goroutine池大小
	DefaultPoolSize = math.MaxInt32
	// 默认清理周期
	DefaultCleanInterval = time.Second
)

const (
	OPEN = iota
	CLOSED
)

var (
	ErrInvalidPoolSize     = errors.New("协程池size不合法")
	ErrLackPoolFunc        = errors.New("")
	ErrInvalidPoolExpiry   = errors.New("")
	ErrLackPoolFunc        = errors.New("")
	ErrPoolClosed          = errors.New("")
	ErrPoolOverload        = errors.New("")
	ErrInvalidPreAllocSize = errors.New("")

	workerChanCap = func() int {
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}
		return 1
	}()
	defaultLogger      = Logger(log.New(os.Stderr, "", log.LstdFlags))
	defaultAntsPool, _ = NewPool(DefaultPoolSize)
)

type Logger interface {
	Printf(format string, args ...interface{})
}

func Submit(task func()) error {
	return nil
}
