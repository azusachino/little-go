package examples

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) Count() int {
	// state
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	// waitting
	v = v >> mutexWaiterShift
	// add up count of lock holder
	v += (v & mutexLocked)
	return int(v)
}

func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
