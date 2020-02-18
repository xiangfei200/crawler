package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item,error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.101:9200"),
		elastic.SetSniff(false))
	if err !=nil{
		return nil,err
	}
	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item"+"%d: %v", itemCount, item)
			itemCount++

			//保存数据到elasticsearch
			err := save(client,index,item)
			if err != nil{
				log.Printf("Item saver : error"+"saving item %v:%v",item,err)
				continue
			}

		}

	}()
	return out,nil
}

//保存到elasticsearch
func save(client *elastic.Client,index string,item engine.Item) error{
	//1.post方式   2.elasticsearch client
	//http.Post()
	// must turn off the sniff 关闭客户端集群维护状态，内网无法维护

	//if err != nil {
	//	return err
	//}
	//存数据 我们想把Type和Id自定义，Index由配置文件配置
	//resp, err := client.Index().
	//	Index("data_profile").
	//	Type("zhenai").
	//	Id(id).
	//	BodyJson(item). //数据本身
	//	// Background returns a non-nil, empty Context. It is never canceled, has no
	//	// values, and has no deadline. It is typically used by the main function,
	//	// initialization, and tests, and as the top-level Context for incoming
	//	// requests.
	//	Do(context.Background())
	//爬出来的数据没有Type或者没有id的解决
	if item.Type == ""{
		return errors.New("must supply type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item) //数据本身
	if item.Id != ""{
		indexService.Id(item.Id)
	}
	_, err := indexService.
		// Background returns a non-nil, empty Context. It is never canceled, has no
		// values, and has no deadline. It is typically used by the main function,
		// initialization, and tests, and as the top-level Context for incoming
		// requests.
		Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v",resp)
	return nil
}
