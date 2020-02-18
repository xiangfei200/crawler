package view

import (
	"os"
	"testing"

	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1238825159",
		Type: "zhenai",
		Id:   "1238825159",
		Payload: common.Profile{
			Name:      "Hanly_L",
			Age:       27,
			Height:    167,
			Income:    "2-5万",
			Marriage:  "未婚",
			Education: "本科",
			Hukou:     "四川遂宁",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}

	// TODO: verify contents in template.test.html
}
