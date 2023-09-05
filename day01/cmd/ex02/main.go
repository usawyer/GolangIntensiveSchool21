package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func Read(filename string) error {
	file, err := os.Open(string(filename))
	if err != nil {
		return errors.Wrap(err, "error opening file")
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {

	}
	return nil
}

func main() {
	//oldFile := flag.String("old", "", "old version of filesystem")
	//newFile := flag.String("new", "", "new version of filesystem")
	//flag.Parse()
	//
	//if *oldFile == "" || *newFile == "" {
	//	fmt.Println("usage: ./compareFS --old file1.txt --new file2.txt")
	//	os.Exit(0)
	//} else {
	//	//Read(*oldFile)
	//	//read
	//	//compare
	//	//print
	//}

	res := map[string]bool{}
	fmt.Println(res)

}
