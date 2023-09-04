package domain

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

type XmlFilename string

func (filename XmlFilename) Read() (*Recipes, error) {
	file, err := os.Open(string(filename))
	if err != nil {
		return nil, errors.Wrap(err, "error opening XML file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading XML file: %v", err)
	}

	var cakes Recipes
	err = xml.Unmarshal(content, &cakes)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling XML file")
	}

	return &cakes, nil
}
