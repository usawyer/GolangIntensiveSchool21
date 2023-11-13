package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"search/internal/domain"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	data, err := domain.ParseDataFromCsv("../../test/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = domain.InsertDataToElastic(es, data)
	if err != nil {
		log.Fatal(err)
	}
}
