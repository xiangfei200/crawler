package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/transaction/parser"
)

var Website string = "http://www.zhenai.com/zhenghun"
var SHWebsite string = "http://www.zhenai.com/zhenghun/shanghai"


func main() {
	//指针接收者
	itemChan, err := persist.ItemSaver("data_profile")
	if err !=nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:        Website,
		ParserFunc: parser.ParseCityList,
	})

}

