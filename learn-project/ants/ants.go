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
	DefaultPoolSize      = math.MaxInt32 // 默认协程池大小：1<<31 - 1
	DefaultCleanInterval = time.Second   // 默认清理周期：1s
)

// 协程池的状态
const (
	OPEN = iota
	CLOSED
)

var (
	ErrInvalidPoolSize     = errors.New("不合法的协程池容量")
	ErrLackPoolFunc        = errors.New("缺少自定义方法")
	ErrInvalidPoolExpiry   = errors.New("不合法的过期协议")
	ErrPoolClosed          = errors.New("协程池已关闭")
	ErrPoolOverload        = errors.New("协程池过载")
	ErrInvalidPreAllocSize = errors.New("不合法的PreAlloc容量")

	// 根据CPU数量判断工作channel是否是buffered chan
	workerChanCap = func() int {
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}
		return 1
	}()

	// 默认采用err日志级别
	defaultLogger = Logger(log.New(os.Stderr, "", log.LstdFlags))
	// 默认初始化的协程池
	defaultAntsPool, _ = NewPool(DefaultPoolSize)
)

// Logger 定义了logger的行为
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

func IsRunning() bool {
	return defaultAntsPool.Running() > 0
}
