package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"goSetUp/crawler/engine"
	"goSetUp/crawler/model"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSaver(t *testing.T) {
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
	//save expected item

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(expected, client, index)
	if err != nil {
		panic(err)
	}
	//get item
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//verify
	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {

	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		fmt.Println("wrong")
	}
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
