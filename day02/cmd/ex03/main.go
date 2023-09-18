package main

import (
	"cli/internal/domain"
	"flag"
	"fmt"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("usage: ./myRotate /path/to/file.log")
	fmt.Println("       ./myRotate -a path/to/store/archived_files /path/to/file.log")
	os.Exit(0)
}

func main() {
	pathF := flag.String("a", "", "path to store archived files")
	flag.Parse()

	if len(os.Args) < 2 {
		printUsage()
	}

	var err error
	if *pathF != "" {
		err = domain.ArchiveFiles(os.Args[3:], *pathF)
	} else {
		err = domain.ArchiveFiles(os.Args[1:], "")
	}

	if err != nil {
		log.Fatal(err)
	}
}
