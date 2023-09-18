package domain

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func GoCount(files []string, flag rune) {
	var wg sync.WaitGroup
	wg.Add(len(files))
	res := make(map[string]int)
	for _, filename := range files {
		go func(filename string) {
			defer wg.Done()
			count(res, filename, flag)
		}(filename)
	}
	wg.Wait()
	printRes(res)
}

func count(res map[string]int, filename string, flag rune) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	switch flag {
	case 'l':
		scanner.Split(bufio.ScanLines)
	case 'm':
		scanner.Split(bufio.ScanRunes)
	case 'w':
		scanner.Split(bufio.ScanWords)
	}

	var counter int
	for scanner.Scan() {
		counter++
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	res[filename] = counter
}

func printRes(res map[string]int) {
	for key, value := range res {
		fmt.Printf("%d\t%s\n", value, key)
	}
}
