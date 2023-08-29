package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func parseInput() []int {
	var data []int
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		tmp := in.Text()

		if tmp == "" {
			fmt.Println("Empty input")
			continue
		}
		num, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		if num > 100000 || num < -100000 {
			fmt.Println("Out of range input")
			continue
		}
		data = append(data, num)
	}
	return data
}

func mean(data []int) float64 {
	var sum float64
	for _, element := range data {
		sum += float64(element)
	}
	return sum / float64(len(data))
}

func median(data []int) float64 {
	var res float64
	sort.Ints(data)

	if len(data)%2 != 0 {
		res = float64(data[len(data)/2])
	} else {
		res = float64(data[len(data)/2]+data[len(data)/2-1]) / 2.0
	}
	return res
}

func mode(data []int) int {
	dict := make(map[int]int)
	var res = []int{0, 0}

	for _, num := range data {
		dict[num]++
		if dict[num] > res[1] || (dict[num] == res[1] && num < res[0]) {
			res[0] = num
			res[1] = dict[num]
		}
	}
	return res[0]
}

func standardDeviation(data []int) float64 {
	var mean = mean(data)
	var sum, sd float64

	for _, num := range data {
		sum += math.Pow(float64(num)-mean, 2.0)
	}

	sd = math.Sqrt(sum / float64(len(data)))
	return sd
}

func getResult(meanF, medianF, modeF, sdF *bool, data []int) {
	if *meanF {
		fmt.Printf("Mean = %.2f\n", mean(data))
	}
	if *medianF {
		fmt.Printf("Median = %.2f\n", median(data))
	}
	if *modeF {
		fmt.Printf("Mode = %d\n", mode(data))
	}
	if *sdF {
		fmt.Printf("Standard deviation = %.2f\n", standardDeviation(data))
	}
}

func main() {
	var meanF = flag.Bool("mean", false, "The average of a data set")
	var medianF = flag.Bool("median", false, "The middle value of a set of numbers")
	var modeF = flag.Bool("mode", false, "The most frequent number")
	var sdF = flag.Bool("sd", false, "The average amount of variability in your data set.")
	flag.Parse()
	if len(os.Args) < 2 {
		*meanF = true
		*medianF = true
		*modeF = true
		*sdF = true
	}

	data := parseInput()
	if len(data) > 0 {
		fmt.Println(data)
		getResult(meanF, medianF, modeF, sdF, data)
	}
}
