package engine

import (
	"crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult,error) {
	log.Printf("Fetching %s",r.Url)
	//fetch 这一部分要去http请求花费了大量的时间，因此在这里可以做成并发
	body, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher : error"+"fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body,r.Url),nil
}
