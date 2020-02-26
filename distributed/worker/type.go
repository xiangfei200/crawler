package worker

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/transaction/parser"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

//创建Requst网络传输所需的结构体
type Request struct {
	Url    string
	Parser SerializedParser
}

//创建ParseResult网络传输所需的结构体
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

//本身的request和engine.request转换
// engine.request 转成本身的request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

//本身的request转成engine.request
func DeserializeRequest(r Request) (engine.Request, error) {
	result, err := deserializeParser(r.Parser)
	//fmt.Printf("%+v",result)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: result,
	}, nil
}

// 本身的parser转成engine.parser
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	//如何从字符串转成一个函数
	//1.把每个parse注册到全局的map中，然后从map中找到每个名字对应的parser
	//2.switch-case
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParse{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

//本身的parseResult和engine.parseResult转换
//engine.parseResult转换成本身的parseResult
func SerializeResult(p engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: p.Items,
	}
	//request需要转换
	for _, req := range p.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

//本身的parseResult转换成engine.parseResult
func DeserializeResult(p ParseResult) engine.ParseResult {
	parseResult := engine.ParseResult{
		Items: p.Items,
	}
	for _, request := range p.Requests {
		//fmt.Printf("%+v",request)
		deserializeRequest, err := DeserializeRequest(request)
		//因为有很多request 如果有报错，则不加入队列，只记录日志
		if err != nil {
			log.Printf("error deserializing request :%v", err)
			continue
		}
		parseResult.Requests = append(parseResult.Requests, deserializeRequest)
		//fmt.Printf("%+v",parseResult)
	}
	return parseResult
}
