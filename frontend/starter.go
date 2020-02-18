package main

import (
	"crawler/frontend/controller"
	"net/http"
)

func main() {
	// 根目录下拉取css等前端静态文件 打开搜索入口页面index.html
	http.Handle("/",
		http.FileServer(http.Dir("./frontend/view")))
	// 搜索结果页面
	http.Handle("/search",
		controller.CreateSearchResultHandler("./frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
