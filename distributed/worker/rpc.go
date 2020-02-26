package worker

import "crawler/engine"

type CrawlService struct {
}

//engine.Requst在网络上无法传输，因为有个interface类型，里面有些函数
//因此需要创建Requst网络传输所需的结构体
// 跟并发版和单机版的独立
func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err !=nil{
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err !=nil{
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
