package db

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"search/internal/types"
	"strings"
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

	settings := `{"index": {"max_result_window": 20000}}`
	req := esapi.IndicesPutSettingsRequest{
		Index: []string{"places"},
		Body:  strings.NewReader(settings),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &ElasticStore{client: es}, nil
}

func (es *ElasticStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	res, err := es.client.Search(
		es.client.Search.WithIndex("places"),
		es.client.Search.WithSize(limit),
		es.client.Search.WithFrom(offset),
		es.client.Search.WithTrackTotalHits(true),
	)

	if err != nil || res.IsError() {
		return nil, 0, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

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

func (es *ElasticStore) GetClosestPlaces(lat float64, lon float64) ([]types.Place, error) {
	query := map[string]interface{}{
		"sort": []map[string]interface{}{
			{
				"_geo_distance": map[string]interface{}{
					"location": map[string]interface{}{
						"lat": lat,
						"lon": lon,
					},
					"order":           "asc",
					"unit":            "km",
					"mode":            "min",
					"distance_type":   "arc",
					"ignore_unmapped": true,
				},
			},
		},
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	res, err := es.client.Search(
		es.client.Search.WithIndex("places"),
		es.client.Search.WithSize(3),
		es.client.Search.WithBody(strings.NewReader(string(queryBytes))),
	)

	if err != nil || res.IsError() {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var responses types.Response
	if err := json.Unmarshal(bodyBytes, &responses); err != nil {
		return nil, err
	}

	var places []types.Place
	for _, response := range responses.Hits.Hits {
		places = append(places, response.Source)
	}

	return places, nil
}
