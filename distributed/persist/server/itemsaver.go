package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/recsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)
//在命令行启动端口
var port = flag.Int("port",0,"the prot for me to listen on ")

func main() {
	//使用log.Fatal避免写panic
	//err := ServerRpc(":1234", "data_profile")
	//if err != nil {
	//	panic(err)
	//}
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	//使用多端口
	//log.Fatal(ServerRpc(fmt.Sprintf(":%d",config.ItemSavePort), config.ElasticDataBase))
	log.Fatal(ServerRpc(fmt.Sprintf(":%d",*port), config.ElasticDataBase))
}

func ServerRpc(host,index string) error{
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticHost),
		elastic.SetSniff(false))
	if err !=nil{
		return err
	}
	return recsupport.ServerRpc(host,&persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}