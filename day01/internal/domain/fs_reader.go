package domain

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type FSReader map[string]string

//func open(filename string) (*os.File, error) {
//	file, err := os.Open(filename)
//	if err != nil {
//		return nil, errors.Wrap(err, "error opening file")
//	}
//	defer file.Close()
//	return file, nil
//}

func (fs FSReader) Read(filename string) error {
	//file, err := open(filename)
	//if err != nil {
	//	return err
	//}

	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "error opening file")
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		//fmt.Println(scan.Text())
		if str := scan.Text(); str != "" {
			fs[scan.Text()] = "REMOVED"
		}
	}
	return nil
}

func (fs FSReader) Compare(filename string) error {
	//file, err := open(filename)
	//if err != nil {
	//	return err
	//}

	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "error opening file")
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if str := scan.Text(); str != "" {
			if _, isExist := fs[str]; !isExist {
				fs[str] = "ADDED"
			} else {
				delete(fs, str)
			}
		}
	}
	return nil
}

func (fs FSReader) Print() {
	for key, value := range fs {
		fmt.Println(value, key)
	}
}
