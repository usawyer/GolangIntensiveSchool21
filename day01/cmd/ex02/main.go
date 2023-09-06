package main

import (
	"bakers/internal/domain"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	oldFile := flag.String("old", "", "old version of filesystem")
	newFile := flag.String("new", "", "new version of filesystem")
	flag.Parse()

	if *oldFile == "" || *newFile == "" {
		fmt.Println("usage: ./compareFS --old file1.txt --new file2.txt")
		os.Exit(0)
	} else {
		oldFS := make(domain.FSReader)
		oldFS.Read(*oldFile)
		err := oldFS.Read(*oldFile)
		if err == nil {
			err = oldFS.Compare(*newFile)
		}
		if err != nil {
			log.Fatal(err)
		}

		oldFS.Print()
	}
}
