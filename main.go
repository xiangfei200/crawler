package main

import (
	"crawler/engine"
	"crawler/fetcher/transaction/parser"
)

var Website string = "http://www.zhenai.com/zhenghun"


func main() {
	engine.Run(engine.Request{
		Url:        Website,
		ParserFunc: parser.ParseCityList,
	})
}

