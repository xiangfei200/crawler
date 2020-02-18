package parser

import (
	"crawler/engine"
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//get到文件内容 手动保存内容到html文件中,不要换行
	//contents, err := fetcher.Fetch(Website)
	contents,err := ioutil.ReadFile("profile_test_data.html")

	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",contents)
	//名字人肉判断，拉下来
	result := parseProfile(contents,"https://album.zhenai.com/u/1238825159","Hanly_L")

	if len(result.Items) != 1{
		t.Errorf("result Items should contain 1 "+"elements;but had %d",len(result.Items))
	}

	actual := result.Items[0]
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
	//rProfile := reflect.TypeOf(profile)
	//vProfile := reflect.ValueOf(profile)
	//vexpected := reflect.ValueOf(expected)
	//for k := 0;k< rProfile.NumField();k++{
	//	filed := rProfile.Field(k).Name
	//	current := vProfile.Field(k).Interface()
	//	excurrent := vexpected.Field(k).Interface()
	//	if current != excurrent {
	//		t.Errorf("expected value of attribute #%s: %v; but "+"current result was %v",filed,excurrent,current)
	//	}
	//}
	//将interface转成字符串 .(string)
	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
