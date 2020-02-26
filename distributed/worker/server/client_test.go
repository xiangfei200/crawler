package main

import (
	"crawler/distributed/config"
	"crawler/distributed/recsupport"
	"crawler/distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawService(t *testing.T) {
	const port = ":8000"
	go recsupport.ServerRpc(port,worker.CrawlService{})
	time.Sleep(time.Second)
	client, err := recsupport.ClientRpc(port)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url:    "https://album.zhenai.com/u/1238825159",
		Parser: worker.SerializedParser{
			Name:config.ParseProfile,
			Args:"Hanly_L",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceMethod, req, &result)
	if err != nil{
		t.Error(err)
	}else{
		fmt.Println(result)
	}
}
