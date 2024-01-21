package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	handler()
}

func handler() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		txt := scanner.Text()
		str := strings.Split(txt, " ")

		switch str[0] {
		case "GET":
			fmt.Println("GET operation")
		case "SET":
			fmt.Println("SET operation")
		case "DELETE":
			fmt.Println("DELETE operation")
		default:
			fmt.Println("Unknown operation")
		}
	}
}
