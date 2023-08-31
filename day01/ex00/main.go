package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"stovetime"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name" xml:"itemname"`
			IngredientCount string `json:"ingredient_count" xml:"itemcount"`
			IngredientUnit  string `json:"ingredient_unit" xml:"itemunit"`
		} `json:"ingredients" xml:"ingredients>item"`
	} `json:"cake" xml:"cake"`
}

type DBReader interface {
	read() Recipes
}

type jsonFilename string
type xmlFilename string

func (filename xmlFilename) read() Recipes {
	file, err := os.Open(string(filename))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var cakes Recipes
	err = xml.Unmarshal(content, &cakes)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	toJson, err := json.MarshalIndent(cakes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(toJson))

	return cakes
}

func (filename jsonFilename) read() Recipes {
	file, err := os.Open(string(filename))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var cakes Recipes
	err = json.Unmarshal(content, &cakes)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	toXml, err := xml.MarshalIndent(cakes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(toXml))

	return cakes
}

func main() {
	file := flag.String("f", "", "File to read")
	flag.Parse()

	var dbreader DBReader

	if filepath.Ext(*file) == ".json" {
		dbreader = jsonFilename(*file)
		dbreader.read()
	} else if filepath.Ext(*file) == ".xml" {
		dbreader = xmlFilename(*file)
		dbreader.read()
	} else {
		fmt.Println("usage: ./readDB -f file.json")
		fmt.Println("       ./readDB -f file.xml")
	}
}
