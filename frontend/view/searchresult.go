package view

import (
	"crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		// Must 的意思是认定不能出错
		template: template.Must(template.ParseFiles(filename)),
	}
}

//把模板文件生成html文件
func (s SearchResultView) Render(
	w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
