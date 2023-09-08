package main

import (
	"cli/internal/domain"
	"flag"
	"fmt"
	"log"
	"os"
)

func checkFlags(fileF, directoryF, symlinksF *bool, extensionF *string) {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./myFind /path/to/dir")
		os.Exit(0)
	}
	if *extensionF != "" && *fileF == false {
		fmt.Println("usage: ./myFind -f -ext 'extension' /path/to/dir")
		os.Exit(0)
	}
	if len(os.Args) < 3 {
		*fileF, *directoryF, *symlinksF = true, true, true
	}
}

func main() {
	fileF := flag.Bool("f", false, "find files")
	directoryF := flag.Bool("d", false, "find directories")
	symlinksF := flag.Bool("sl", false, "find symlinks")
	extensionF := flag.String("ext", "", "find certain extension files")
	flag.Parse()
	checkFlags(fileF, directoryF, symlinksF, extensionF)

	path := os.Args[len(os.Args)-1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	finder := domain.Finder(path)
	finder.Find(*fileF, *directoryF, *symlinksF, *extensionF)
}
