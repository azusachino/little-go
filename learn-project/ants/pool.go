package ants

import (
	"github.com/little-go/learn-project/ants/internal"
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	capacity    int32       // 存储容量的地址
	running     int32       // 存储运行中协程数量的地址
	workers     workerArray // 任务池
	state       int32       // 存储协程池状态的地址
	lock        sync.Locker // 协程池锁
	cond        *sync.Cond  // Cond => 达到某种条件，Signal所有阻塞协程
	workerCache sync.Pool   // help func:retrieveWorker
	blockingNum int         // pool.Submit之后阻塞的任务数量
	options     *Options    // 协程池配置
}

// NewPool 创建新的协程池
func NewPool(size int, options ...Option) (*Pool, error) {
	opts := loadOptions(options...)

	// 不限制容量
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
	// 采用标准库中的sync.Pool装载goWorker
	p.workerCache.New = func() interface{} {
		return &goWorker{
			pool: p,
			task: make(chan func(), workerChanCap),
		}
	}
	// 如果定义了PreAlloc，采用环形队列
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

// Submit 向池提交任务
func (p *Pool) Submit(task func()) error {
	if atomic.LoadInt32(&p.state) == CLOSED {
		return ErrPoolClosed
	}
	var w *goWorker
	if w = p.retrieveWorker(); w == nil {
		return ErrPoolOverload
	}
	w.task <- task
	return nil
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *Pool) Free() int {
	return p.Cap() - p.Running()
}

func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) Tune(size int) {
	if capacity := p.Cap(); capacity == -1 || size <= 0 || size == capacity || p.options.PreAlloc {
		return
	}
	atomic.StoreInt32(&p.capacity, int32(size))
}

func (p *Pool) Release() {
	atomic.StoreInt32(&p.state, CLOSED)
	p.lock.Lock()
	p.workers.reset()
	p.lock.Unlock()
}

func (p *Pool) Reboot() {
	if atomic.CompareAndSwapInt32(&p.state, CLOSED, OPEN) {
		go p.purgePeriodically()
	}
}

func (p *Pool) increaseRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) decreaseRunning() {
	atomic.AddInt32(&p.running, -1)
}

// 获取一个可用的goWorker
func (p *Pool) retrieveWorker() (w *goWorker) {
	// 获取goWorker并执行任务
	spawnWorker := func() {
		w = p.workerCache.Get().(*goWorker)
		w.run()
	}
	p.lock.Lock()

	w = p.workers.detach()
	if w != nil {
		p.lock.Unlock()
	} else if capacity := p.Cap(); capacity == -1 || p.Running() < capacity {
		p.lock.Unlock()
		spawnWorker()
	} else {
		if p.options.NonBlocking {
			p.lock.Unlock()
			return
		}
	ReEntry:
		if p.options.MaxBlockingTasks != 0 && p.blockingNum >= p.options.MaxBlockingTasks {
			p.lock.Unlock()
			return
		}
		p.blockingNum++
		p.cond.Wait()
		p.blockingNum--
		if p.Running() == 0 {
			p.lock.Unlock()
			spawnWorker()
			return
		}

		// 从workArray中获取一个goWorker
		w = p.workers.detach()
		// 重试
		if w == nil {
			goto ReEntry
		}

		p.lock.Unlock()

	}
	return
}

// revertWorker puts a worker back into free pool, recycling the goroutines.
func (p *Pool) revertWorker(worker *goWorker) bool {
	if capacity := p.Cap(); (capacity > 0 && p.Running() > capacity) || atomic.LoadInt32(&p.state) == CLOSED {
		return false
	}
	worker.recycleTime = time.Now()
	p.lock.Lock()

	// To avoid memory leaks, add a double check in the lock scope.
	if atomic.LoadInt32(&p.state) == CLOSED {
		p.lock.Unlock()
		return false
	}

	err := p.workers.insert(worker)
	if err != nil {
		p.lock.Unlock()
		return false
	}

	// Notify the invoker stuck in 'retrieveWorker()' of there is an available worker in the worker queue.
	p.cond.Signal()
	p.lock.Unlock()
	return true
}

// 定期清除过期任务
func (p *Pool) purgePeriodically() {

	// 定时器
	heartBeat := time.NewTicker(p.options.ExpiryDuration)

	defer heartBeat.Stop()

	for range heartBeat.C {
		// 直到协程池CLOSED为止
		if atomic.LoadInt32(&p.state) == CLOSED {
			break
		}
		// 获取已超时任务（线程安全）
		p.lock.Lock()
		expiredWorkers := p.workers.retrieveExpiry(p.options.ExpiryDuration)
		p.lock.Unlock()

		// 直接清除
		for i := range expiredWorkers {
			expiredWorkers[i].task = nil // help GC
			expiredWorkers[i] = nil
		}

		// 目前没有任何正在运行的任务，唤醒那些阻塞在 p.cond.Wait() 的任务
		if p.Running() == 0 {
			p.cond.Broadcast()
		}
	}
}
