package ants

import (
	"errors"
	"time"
)

var (
	errQueueIsFull     = errors.New("队列已满")
	errQueueIsReleased = errors.New("队列已释放")
)

type workerArray interface {
	len() int
	isEmpty() bool
	insert(worker *goWorker) error
	detach() *goWorker
	retrieveExpiry(duration time.Duration) []*goWorker
	reset()
}

type arrayType int

const (
	// 堆栈
	stackType arrayType = 1 << iota
	// 环形队列
	loopQueueType
)

func newWorkerArray(aType arrayType, size int) workerArray {
	switch aType {
	case stackType:
		return newWorkerStack(size)
	case loopQueueType:
		return newWorkerLoopQueue(size)
	default:
		return newWorkerStack(size)
	}
}
