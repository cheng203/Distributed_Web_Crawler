package main

import (
	"goSetUp/crawler/frontend/controler"
	"net/http"
)

func main() {
	http.Handle("/search", controler.CreateSearchResultHandler("crawler/frontend/view/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
