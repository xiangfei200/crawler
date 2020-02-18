package parser

import (
	"io/ioutil"
	"testing"
)
//var Website string = "http://www.zhenai.com/zhenghun"
/**
表格测试citylist的解析器
 */
func TestParseCityList(t *testing.T) {
	//get到文件内容 手动保存内容到html文件中,不要换行
	//contents, err := fetcher.Fetch(Website)
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",contents)

	result := ParseCityList(contents,"")
	//470城市列表/请求
	const resultSize = 470
	//把正确的结果，人工对比后贴进来即可
	expectedUrls := []string{
	"","","",
	}
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d "+"results;but had %d",resultSize,len(result.Requests))
	}
	for i,url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d: %s; but "+"was %s",i,url,result.Requests[i].Url)
		}
	}
	if len(result.Items) != resultSize{
		t.Errorf("result should have %d "+"results;but had %d",resultSize,len(result.Items))
	}
}
