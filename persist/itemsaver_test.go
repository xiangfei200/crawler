package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
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
	//保存数据
	const index = "data_test"
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.101:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = save(client,index,expected)
	if err != nil {
		panic(err)
	}

	//拿出来对比
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	//t.Logf("%s",resp.Source)
	var acutal engine.Item
	err = json.Unmarshal(*resp.Source, &acutal)

	//把playload转化成profile类型
	transferProfile, _ := model.FromJsonObj(acutal.Payload)

	acutal.Payload = transferProfile
	//验证是否相同
	if acutal != expected {
		t.Errorf("got:%+v;expected:%+v", acutal, expected)
	}
}
