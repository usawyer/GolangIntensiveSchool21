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
	var buf bytes.Buffer
	for _, place := range places {
		action := fmt.Sprintf(`{ "index" : { "_index" : "places", "_id" : "%d" } }%s`, place.ID+1, "\n")
		if _, err := buf.WriteString(action); err != nil {
			return errors.Wrap(err, "error writing action: "+strconv.Itoa(place.ID))
		}
		if err := json.NewEncoder(&buf).Encode(place); err != nil {
			return errors.Wrap(err, "error encoding document: "+strconv.Itoa(place.ID))
		}
		log.Printf("[ %d place is being processed ]\n", place.ID+1)
	}

	req := esapi.BulkRequest{
		Body:    &buf,
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return errors.Wrap(err, "error performing bulk request")
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.New("error response for bulk request")
	}

	return nil

}
