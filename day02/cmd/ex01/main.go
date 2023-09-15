package main

import (
	"cli/internal/domain"
	"flag"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("usage: ./myWc -l file")
	fmt.Println("       ./myWc -m file")
	fmt.Println("       ./myWc -w file")
}

func main() {
	linesF := flag.Bool("l", false, "counts lines")
	charactersF := flag.Bool("m", false, "counts characters")
	wordsF := flag.Bool("w", false, "counts words")
	flag.Parse()

	if (!*linesF && !*charactersF && !*wordsF) || (!*linesF && !*charactersF && *wordsF) {
		domain.GoCount(flag.Args()[:], 'w')
	} else if *linesF && !*charactersF && !*wordsF {
		domain.GoCount(flag.Args()[:], 'l')
	} else if !*linesF && *charactersF && !*wordsF {
		domain.GoCount(flag.Args()[:], 'm')
	} else {
		printUsage()
		os.Exit(0)
	}
}
