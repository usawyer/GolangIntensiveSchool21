package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepSort(arr []int) <-chan int {
	resultChan := make(chan int, len(arr))

	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, value := range arr {
		go func(value int) {
			defer wg.Done()
			time.Sleep(time.Duration(value) * time.Second)
			resultChan <- value
		}(value)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 0, 2, 6, 5, 3, 5, 0}
	resultChan := sleepSort(arr)
	for v := range resultChan {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
