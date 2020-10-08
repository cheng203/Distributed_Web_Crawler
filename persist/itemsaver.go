package persist

import (
	"context"
	"errors"
	"goSetUp/crawler/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++
			err := Save(item, client, index)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}

		}
	}()
	return out, nil
}

func Save(item engine.Item, client *elastic.Client, index string) error {
	if item.Type == "" {
		return errors.New("Must supply type")
	}
	indexService := client.Index().Index(index).Type(item.Type)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v", resp)
	return nil
}
