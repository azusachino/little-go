package ants

import (
	"runtime"
	"time"
)

// 类似于Java的Runnable
type goWorker struct {
	pool        *Pool       // 所属协程池
	task        chan func() // 实际运行的方法
	recycleTime time.Time
}

func (w *goWorker) run() {
	// count + 1
	w.pool.increaseRunning()

	go func() {
		// 回收资源
		defer func() {
			// count - 1
			w.pool.decreaseRunning()
			// 放回snyc.Pool等待下次retrieveGoWorker
			w.pool.workerCache.Put(w)
			if p := recover(); p != nil {
				if ph := w.pool.options.PanicHandler; ph != nil {
					ph(p)
				} else {
					w.pool.options.Logger.Printf("worker exits from a panic: %v \n", p)
					var buf [4096]byte
					n := runtime.Stack(buf[:], false) // 异常堆栈信息
					w.pool.options.Logger.Printf("worker exits from panic: %s\n", string(buf[:n]))
				}
			}
		}()
		for f := range w.task {
			if f == nil {
				return
			}
			// real function
			f()
			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}
