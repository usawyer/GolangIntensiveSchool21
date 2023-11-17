package db

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"search/internal/types"
)

type ElasticStore struct {
	client *elasticsearch.Client
}

func NewElasticStore() (*ElasticStore, error) {
	cfg := elasticsearch.Config{Addresses: []string{"http://localhost:9200"}}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ElasticStore{client: es}, nil
}

func (es *ElasticStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	res, err := es.client.Search(
		es.client.Search.WithIndex("places"),
		es.client.Search.WithSize(limit),
		es.client.Search.WithFrom(offset),
		es.client.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, 0, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()

	var responses types.Response
	if err := json.Unmarshal(bodyBytes, &responses); err != nil {
		return nil, 0, err
	}

	totalHits := responses.Hits.Total.Value
	var places []types.Place

	for _, response := range responses.Hits.Hits {
		places = append(places, response.Source)
	}

	return places, totalHits, nil
}
