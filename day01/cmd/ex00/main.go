package main

import (
	"bakers/internal/domain"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	file := flag.String("f", "", "File to read")
	flag.Parse()

	var dbreader domain.DBReader

	if filepath.Ext(*file) == ".json" {
		dbreader = domain.JsonFilename(*file)
	} else if filepath.Ext(*file) == ".xml" {
		dbreader = domain.XmlFilename(*file)
	} else {
		fmt.Println("usage: ./readDB -f file.json")
		fmt.Println("       ./readDB -f file.xml")
		os.Exit(0)
	}

	recipes, err := dbreader.Read()
	if err != nil {
		log.Fatal(err)
	}

	if filepath.Ext(*file) == ".json" {
		recipes.PrintXML()
	} else if filepath.Ext(*file) == ".xml" {
		recipes.PrintJSON()
	}

}
