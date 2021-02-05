package main

import (
	"github.com/little-go/learn-go/crawler/engine"
	"github.com/little-go/learn-go/crawler/scheduler"
	"github.com/little-go/learn-go/crawler/types"
	"github.com/little-go/learn-go/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
