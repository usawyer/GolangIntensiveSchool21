package domain

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

type JsonFilename string

func (filename JsonFilename) Read() (*Recipes, error) {
	file, err := os.Open(string(filename))
	if err != nil {
		return nil, errors.Wrap(err, "error opening JSON file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}

	var cakes Recipes
	err = json.Unmarshal(content, &cakes)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling JSON file")
	}

	return &cakes, nil
}
