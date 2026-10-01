package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Formula struct {
	Name    string
	Find    string
	Needed  []string
	Compute func(args map[string]float64)
}

func main() {
	fmt.Println("Calculator of kinematics")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What needs to be found?")
	input, _ := reader.ReadString('\n')
	target := strings.TrimSpace(input)
	fmt.Print("Enter the known data: ")
	readFloat(reader)
}

func readFloat(reader *bufio.Reader) float64 {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0
	}
	return val
}
