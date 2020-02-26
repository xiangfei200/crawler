package main

import (
	"crawler/distributed/recsupport"
	"crawler/distributed/worker"
	"flag"
	"fmt"
	"log"
)
//在命令行启动端口
var port = flag.Int("port",0,"the prot for me to listen on ")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	//使用多端口
	//log.Fatal(recsupport.ServerRpc(fmt.Sprintf(":%d",config.WorkerPort0), worker.CrawlService{}))
	log.Fatal(recsupport.ServerRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
