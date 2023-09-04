package main

import (
	"bakers/internal/domain"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"path/filepath"
)

func readData(file string) (*domain.Recipes, error) {
	var dbreader domain.DBReader
	if filepath.Ext(file) == ".json" {
		dbreader = domain.JsonFilename(file)
	} else if filepath.Ext(file) == ".xml" {
		dbreader = domain.XmlFilename(file)
	} else {
		return nil, errors.New("usage:  ./compareDB --old file.xml --new file.json")
	}
	return dbreader.Read()
}

func printResult(res []string) {
	for _, str := range res {
		fmt.Println(str)
	}
}

func main() {
	oldFile := flag.String("old", "", "old version of recipes")
	newFile := flag.String("new", "", "new version of recipes")
	flag.Parse()

	oldRecipes, err := readData(*oldFile)
	var newRecipes *domain.Recipes
	if err == nil {
		newRecipes, err = readData(*newFile)
	}
	if err != nil {
		log.Fatal(err)
	}

	res, err := oldRecipes.Compare(newRecipes)
	printResult(res)
}
