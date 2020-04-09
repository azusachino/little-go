package main

import (
	"learn_project/crawler/engine"
	"learn_project/crawler/parser/zhenai"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenhun",
		ParseFunc: zhenai.ParseCityList,
	})
}
