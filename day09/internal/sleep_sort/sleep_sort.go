package sort

import (
	"sync"
	"time"
)

func SleepSort(arr []int) <-chan int {
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
