package internal

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

func (sl *spinLock) Lock() {
	// 地址 -> 1，认为是有锁状态
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		// Gosched yields the processor, allowing other goroutines to run.
		runtime.Gosched()
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

func NewSpinLock() sync.Locker {
	// 取一块地址，初始化为0
	return new(spinLock)
}
