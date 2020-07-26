package main

import (
	"learnGo/crawler/engine"
	"learnGo/crawler/scheduler"
	"learnGo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimplerScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
