package parser

import (
	"fmt"
	"goSetUp/crawler/engine"
	"goSetUp/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data2.html")
	if err != nil {
		panic(err)
	}
	results := ParseProfile(contents, "https://album.zhenai.com/u/1641748432", "雕刻时光")
	fmt.Println(results.Items)
	if len(results.Items) != 1 {
		t.Errorf("Items should contain 1 element "+"but it was %v", results.Items)
	}
	actual := results.Items[0]

	//Name     string
	//	Age      int
	//	Height   int
	//	Weight   int
	//	Income   string
	//	Marriage string
	//	Origin   string

	expected := engine.Item{
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
	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
	fmt.Println(actual)
}
