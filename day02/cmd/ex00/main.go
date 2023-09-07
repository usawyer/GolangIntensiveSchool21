package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	//"github.com/pkg/errors"
	"os"
)

func print(split []string) {
	for _, str := range split {
		fmt.Println(str)
	}
}

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

	var res []string

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if info.Mode() != 0 {
				if *symlinksF && info.Mode().Type() == os.ModeSymlink {
					//fmt.Print(path + "->")
					if destination, er := filepath.EvalSymlinks(path); er != nil {
						fmt.Println(path + "->" + "[broken]")
						//res = append(res, path+"->"+"[broken]")
					} else {
						fmt.Println(path + "->" + destination)
						//res = append(res, path+"->"+destination)
					}
				} else if *fileF && info.Mode().IsRegular() {
					if *extensionF == "" || (*extensionF != "" && filepath.Ext(path) == "."+(*extensionF)) {
						//res = append(res, path)
						fmt.Println(path)
					}
				} else if *directoryF && info.IsDir() {
					//res = append(res, path)
					fmt.Println(path)

				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	print(res)
	//fmt.Println(len(res))

}
