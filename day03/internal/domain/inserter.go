package domain

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pkg/errors"
	"log"
	"os"
	"search/internal/types"
	"strconv"
)

const (
	indexName string = "places"
	batch     int    = 2500
)

func ParseDataFromCsv(path string) ([]types.Place, error) {
	var data []types.Place

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

		data = append(data, types.Place{
			ID:      id + 1,
			Name:    record[1],
			Address: record[2],
			Phone:   record[3],
			Location: types.Location{
				Lon: lon,
				Lat: lat}})
	}

	log.Printf("→ Generated %s places", humanize.Comma(int64(len(data))))
	return data, nil
}

func InsertDataToElastic(es *elasticsearch.Client, places []types.Place) error {
	fmt.Print("→ Sending batch ")
	var buf bytes.Buffer
	for i, place := range places {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, place.ID, "\n"))

		data, err := json.Marshal(place)
		if err != nil {
			return errors.Wrap(err, "Cannot encode place "+strconv.Itoa(place.ID))
		}
		data = append(data, "\n"...)

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)

		if (i+1)%batch == 0 || i == len(places)-1 {
			fmt.Printf("[%d/%d] ", i+1, len(places))

			res, err := es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
			if err != nil {
				return errors.Wrap(err, "failure indexing batch "+strconv.Itoa(place.ID))
			}
			if res.IsError() {
				var raw map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
					return errors.Wrap(err, "failure to to parse response body")
				} else {
					log.Printf("  Error: [%d] %s: %s",
						res.StatusCode,
						raw["error"].(map[string]interface{})["type"],
						raw["error"].(map[string]interface{})["reason"],
					)
				}
			}
			res.Body.Close()
			buf.Reset()
		}
	}

	return nil
}
