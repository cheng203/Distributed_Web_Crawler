package view

import (
	"goSetUp/crawler/engine"
	"goSetUp/crawler/frontend/control"
	"goSetUp/crawler/model"
	"os"
	"testing"
)

func TestIndex(t *testing.T) {
	searchView := CreateSearchResultView("index.html")
	out, err := os.Create("template.test.html")
	page := control.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1844791299",
		Type: "zhenai",
		Id:   "1844791299",
		Payload: model.Profile{
			Name:     "真的爱你",
			Age:      33,
			Height:   175,
			Weight:   70,
			Income:   "1.2-2",
			Marriage: "未婚",
			Origin:   "四川成都",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = searchView.Render(out, page)
	if err != nil {
		panic(err)
	}
}
