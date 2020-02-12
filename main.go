package main

import (
	"crawler/engine"
	"crawler/fetcher/transaction/parser"
	"crawler/scheduler"
)

var Website string = "http://www.zhenai.com/zhenghun"


func main() {
	//指针接收者
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        Website,
		ParserFunc: parser.ParseCityList,
	})
}

