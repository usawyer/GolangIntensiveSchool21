package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var args []string
	args = os.Args[2:]

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		args = append(args, strings.Fields(in.Text())...)
	}

	cmd := exec.Command(os.Args[1], args...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", stdout)
}
