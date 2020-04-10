package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkerChan(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	in := make(chan Request)
	out := make(chan ParseResult)

	for i:=0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <- in
			res, err := worker(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
