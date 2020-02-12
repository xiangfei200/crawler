package engine

import (
	"crawler/fetcher"
	"log"
)
type SimpleEngine struct {}

func (e SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil{
			continue
		}
		//parseResult.Requests... = parseResult.Requests[0] parseResult.Requests[1] ...[end]
		requests = append(requests,parseResult.Requests...)
		//fmt.Println(requests)
		for _,item := range parseResult.Items{
			//%v 输出不转义的字符，如果item本身是ascii码的话，原方法string()转化下即可
			log.Printf("Got item %v",item)
		}
	}
}

func worker(r Request) (ParseResult,error) {
	log.Printf("Fetching %s",r.Url)
	//fetch 这一部分要去http请求花费了大量的时间，因此在这里可以做成并发
	body, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher : error"+"fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}

