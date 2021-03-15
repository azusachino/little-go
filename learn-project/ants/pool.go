package ants

import (
	"github.com/little-go/learn-project/ants/internal"
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	capacity    int32
	running     int32
	workers     workerArray
	state       int32
	lock        sync.Locker
	cond        *sync.Cond
	workerCache sync.Pool
	blockingNum int
	options     *Options
}

func (p *Pool) purgePeriodically() {
	heartBeat := time.NewTicker(p.options.ExpiryDuration)
	defer heartBeat.Stop()

	for range heartBeat.C {
		if atomic.LoadInt32(&p.state) == CLOSED {
			break
		}
		p.lock.Lock()
		expiredWorkers := p.workers.retrieveExpiry(p.options.ExpiryDuration)
		p.lock.Unlock()

		for i := range expiredWorkers {
			expiredWorkers[i].task = nil
			expiredWorkers[i] = nil
		}

		if p.Running() == 0 {
			p.cond.Broadcast()
		}
	}
}

func NewPool(size int, options ...Option) (*Pool, error) {
	opts := loadOptions(options...)
	if size <= 0 {
		size = -1
	}

	if expiry := opts.ExpiryDuration; expiry < 0 {
		return nil, ErrInvalidPoolExpiry
	} else if expiry == 0 {
		opts.ExpiryDuration = DefaultCleanInterval
	}

	if opts.Logger == nil {
		opts.Logger = defaultLogger
	}

	p := &Pool{
		capacity: int32(size),
		lock:     internal.NewSpinLock(),
		options:  opts,
	}
	p.workerCache.New = func() interface{} {
		return &goWorker{
			pool: p,
			task: make(chan func(), workerChanCap),
		}
	}
	if p.options.PreAlloc {
		if size == -1 {
			return nil, ErrInvalidPreAllocSize
		}
		p.workers = newWorkerArray(loopQueueType, size)
	} else {
		p.workers = newWorkerArray(stackType, 0)
	}

	p.cond = sync.NewCond(p.lock)

	// Start a goroutine to clean up expired workers periodically.
	go p.purgePeriodically()

	return p, nil
}
