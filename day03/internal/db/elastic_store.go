package db

import (
	"github.com/elastic/go-elasticsearch/v8"
	"search/internal/types"
)

type ElasticStore struct {
	client *elasticsearch.Client
}

func NewElasticStore(client *elasticsearch.Client) *ElasticStore {
	return &ElasticStore{client: client}
}

func (es *ElasticStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	return nil, 0, nil
}
