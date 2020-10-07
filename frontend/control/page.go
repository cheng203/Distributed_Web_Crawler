package control

import "goSetUp/crawler/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
