package domain

import (
	"encoding/csv"
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
