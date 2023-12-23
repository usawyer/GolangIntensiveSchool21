package main

import (
	fanin "day09/internal/fan_in"
	"fmt"
)

func intGenerator(begin, end int) chan interface{} {
	ch := make(chan interface{})
	go func() {
		for i := begin; i < end; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func stringGenerator(begin, end string) chan interface{} {
	ch := make(chan interface{})
	go func() {
		for i := begin; i < end; i = incrementString(i) {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func incrementString(s string) string {
	r := []rune(s)
	r[len(r)-1]++
	return string(r)
}

func main() {
	fanIn := fanin.Multiplex(
		intGenerator(0, 5),
		intGenerator(5, 10),
	)

	for val := range fanIn {
		fmt.Println(val)
	}

	fmt.Println()

	fanIn = fanin.Multiplex(
		stringGenerator("a", "d"),
		stringGenerator("e", "g"),
	)

	for val := range fanIn {
		fmt.Println(val)
	}

}
