package scheduler

import "learn_project/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s SimpleScheduler) ConfigureWorkerChan(chan engine.Request) {

}

func (s SimpleScheduler) Submit(request engine.Request) {
	s.workerChan <- request
}