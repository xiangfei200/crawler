package parser

import (
	"crawler/model"
	"io/ioutil"
	"reflect"
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
	result := parseProfile(contents,"真诚点")

	if len(result.Items) != 1{
		t.Errorf("result Items should contain 1 "+"elements;but had %d",len(result.Items))
	}

	profile := result.Items[0].(model.Profile)
	//把正确的结果，人工对比后贴进来即可
	expected := model.Profile{
		Name:"真诚点",
		Age : 32,
		Height :162,
		Income :"1.2-2万",
		Marriage :"未婚",
		Education :"大专",
		Hukou :"四川成都",
	}
	rProfile := reflect.TypeOf(profile)
	vProfile := reflect.ValueOf(profile)
	vexpected := reflect.ValueOf(expected)
	for k := 0;k< rProfile.NumField();k++{
		filed := rProfile.Field(k).Name
		current := vProfile.Field(k).Interface()
		excurrent := vexpected.Field(k).Interface()
		if current != excurrent {
			t.Errorf("expected value of attribute #%s: %v; but "+"current result was %v",filed,excurrent,current)
		}
	}
	//将interface转成字符串 .(string)
	//if expected.(string) != city{
	//	t.Errorf("expected city #%d: %s; but "+"was %s",i,city,result.Items[i].(string))
	//}
}
