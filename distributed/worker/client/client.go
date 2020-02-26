package client

import (
	"crawler/distributed/config"
	"crawler/distributed/worker"
	"crawler/engine"
	"net/rpc"
)

//创建worker客户端
//传统的需要很多数量的worker使用一个slice（数组），然后加锁，同步等等来解决并发问题
//go 可以直接使用chan来解决
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := recsupport.ClientRpc(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil{
	//	return nil,err
	//}

	return func(req engine.Request) (engine.ParseResult,error) {
		// 转换成用于网络传输的格式
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <- clientChan
		//实际请求调用
		err := c.Call(config.CrawlServiceMethod, sReq, &sResult)
		if err !=nil{
			return engine.ParseResult{},err
		}
		//// 网络传输转换成适用本地的格式
		return worker.DeserializeResult(sResult),nil

	}
}