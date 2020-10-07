package frontend

import (
	"goSetUp/crawler/frontend/control"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (s *SearchResultView) Render(w io.Writer, data control.SearchResult) error {
	return s.template.Execute(w, data)
}
