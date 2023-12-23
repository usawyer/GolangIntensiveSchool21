package main

import (
	sort "day09/internal/sleep_sort"
	"fmt"
)

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 0, 2, 6, 5, 3, 5, 0}
	resultChan := sort.SleepSort(arr)
	for v := range resultChan {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
