package client

import (
	"crawler/distributed/config"
	"crawler/distributed/recsupport"
	"crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item,error) {
	client, err := recsupport.ClientRpc(host)
	if err!= nil{
		return nil,err
	}
	outer := make(chan engine.Item)
	go func() {
		result := ""
		itemCount := 0
		for {
			item := <-outer
			//log.Printf("Item Saver:got item"+"%d: %v", itemCount, item)
			itemCount++

			//保存数据到elasticsearch，此时调用rpc服务
			//在rpc server中注册了ItemSaverService，因此可以直接调用，不需要包名之类的
			//反正在goroutine中 所以等待返回 卡住也没关系
			err = client.Call(config.ItemSaverRpcMethod, item, &result)
			if err!= nil || result != "ok"{
				log.Printf("save result :%v has error:%s",result,err)
				continue
			}
		}

	}()
	return outer,nil
}