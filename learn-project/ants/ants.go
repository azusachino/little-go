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
	ErrPoolClosed          = errors.New("")
	ErrPoolOverload        = errors.New("")
	ErrInvalidPreAllocSize = errors.New("")

	workerChanCap = func() int {
		// GOMAXPROCS sets the maximum number of CPUs that can be executing
		// simultaneously and returns the previous setting.
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

// Submit submits a task to pool.
func Submit(task func()) error {
	return defaultAntsPool.Submit(task)
}

// Running returns the number of the currently running goroutines.
func Running() int {
	return defaultAntsPool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return defaultAntsPool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return defaultAntsPool.Free()
}

// Release Closes the default pool.
func Release() {
	defaultAntsPool.Release()
}

// Reboot reboots the default pool.
func Reboot() {
	defaultAntsPool.Reboot()
}
