package main

import (
	"crawler/distributed/config"
	itemsaver "crawler/distributed/persist/client"
	"crawler/distributed/recsupport"
	worker "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/transaction/parser"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var Website string = "http://www.zhenai.com/zhenghun"
var SHWebsite string = "http://www.zhenai.com/zhenghun/shanghai"

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
	workerHosts = flag.String("worker_host","","worker hosts (comma separated)")
)

func main() {
	//指针接收者
	//itemsaver rpc
	flag.Parse()
	if *itemSaverHost  == ""|| *workerHosts == "" {
		fmt.Println("must specify itemsaver host and worker host")
		return
	}
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	//itemChan, err := persist.ItemSaver("data_profile")
	if err != nil {
		panic(err)
	}
	// worker rpc
	//create client pool
	pool := createClientPool(strings.Split(*workerHosts,","))
	processor := worker.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    Website,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _,h := range hosts{
		client, err := recsupport.ClientRpc(h)
		if err == nil{
			clients = append(clients,client)
			log.Printf("Connected to %s",h)
		}else{
			log.Printf("Error connecting to %s",h)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		//顺序分发，轮完一轮之后，还要第二轮，第三轮，因此外面再套一层for
		for {
			for _,client := range clients{
				out <- client
			}
		}

	}()
	return out
}