package scheduler

import "learn_project/crawler/types"

type Scheduler interface {
	Submit(types.Request)

	WorkerChan() chan types.Request

	ReadyNotifier

	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan types.Request)
}
