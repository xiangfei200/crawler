package main

import (
	"crawler/distributed/config"
	"crawler/distributed/recsupport"
	"crawler/engine"
	"crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	const host = ":1236"
	//start ItemSaveServer
	go ServerRpc(host, "test2")

	//偷懒强制睡眠1秒，确保rpcserver启动起来
	time.Sleep(time.Second)

	//start ItemSaveClent
	client, err := recsupport.ClientRpc(host)
	if err != nil{
		panic(err)
	}
	//Call save
	//把正确的结果，人工对比后贴进来即可
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/1238825159",
		Type: "zhenai",
		Id:   "1238825159",
		Payload: model.Profile{
			Name:      "Hanly_L",
			Age:       27,
			Height:    167,
			Income:    "2-5万",
			Marriage:  "未婚",
			Education: "本科",
			Hukou:     "四川遂宁",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpcMethod, expected, &result)
	if err !=nil || result != "ok"{
		t.Errorf("result: %s;err: %s",result,err)
	}
}