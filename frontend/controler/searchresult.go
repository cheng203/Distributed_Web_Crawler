package controler

import (
	"context"
	"goSetUp/crawler/engine"
	"goSetUp/crawler/frontend/control"
	"goSetUp/crawler/frontend/view"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client}
}

//localhost:8888/_search?q=nan,yigoufang&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
func (h SearchResultHandler) getSearchResult(q string, from int) (control.SearchResult, error) {
	var result control.SearchResult
	resp, err := h.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	return result, nil
}
