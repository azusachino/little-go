package ants

import "time"

// 类似于Java的Runnable
type goWorker struct {
	pool        *Pool
	task        chan func()
	recycleTime time.Time
}

func (w *goWorker) run() {
	// TODO

	go func() {
		// 回收资源
		defer func() {

		}()
		for f := range w.task {
			if f == nil {
				return
			}
			f()
			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}
