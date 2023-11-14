package domain

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"sync"
)

type Place struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location Location `json:"location"`
}

type Location struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

func ParseDataFromCsv(path string) ([]Place, error) {
	var data []Place

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = '\t'
	info, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range info[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("id %s converting error: %s", record[0], err)
			continue
		}

		lon, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Printf("id %d latitude converting error: %s", id, err)
			continue
		}

		lat, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Printf("id %d longitude converting error: %s", id, err)
			continue
		}

		data = append(data, Place{
			ID:      id,
			Name:    record[1],
			Address: record[2],
			Phone:   record[3],
			Location: Location{
				Lon: lon,
				Lat: lat}})
	}

	return data, nil
}

func InsertDataToElastic(es *elasticsearch.Client, places []Place) error {
	var wg sync.WaitGroup

	errs := make(chan error, len(places))
	wg.Add(len(places))

	for _, place := range places {
		go func(place Place) {
			defer wg.Done()
			errs <- IndexPlace(es, place)
		}(place)
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return errors.Wrap(err, "error indexing document: ")
		}
	}

	return nil
}

func IndexPlace(es *elasticsearch.Client, place Place) error {
	doc, err := json.Marshal(place)
	if err != nil {
		return errors.Wrap(err, "error marshaling document: "+strconv.Itoa(place.ID))
	}

	req := esapi.IndexRequest{
		Index:      "places",
		DocumentID: fmt.Sprintf("%d", place.ID),
		Body:       bytes.NewReader(doc),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return errors.Wrap(err, "error performing request: "+strconv.Itoa(place.ID))
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.Wrap(err, "error response: "+strconv.Itoa(place.ID))
	}

	return nil
}
