package examples

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
	once        *Once
)

// 利用标志位
func Instance() *singleton {
	// 初始化标志位
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

// 利用Once
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
