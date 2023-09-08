package domain

import (
	"fmt"
	"os"
	"path/filepath"
)

type Finder string

func (f Finder) Find(fileF, directoryF, symlinksF bool, extensionF string) {
	errs := filepath.Walk(string(f), func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if info.Mode() != 0 {
				if symlinksF && info.Mode().Type() == os.ModeSymlink {
					if destination, er := filepath.EvalSymlinks(path); er != nil {
						fmt.Println(path + "->" + "[broken]")
					} else {
						fmt.Println(path + "->" + destination)
					}
				} else if fileF && info.Mode().IsRegular() {
					if extensionF == "" || (extensionF != "" && filepath.Ext(path) == "."+(extensionF)) {
						fmt.Println(path)
					}
				} else if directoryF && info.IsDir() {
					fmt.Println(path)
				}
			}
		}
		return nil
	})
	_ = errs
}
