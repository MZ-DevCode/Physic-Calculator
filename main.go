package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	V  float64 = 0.0
	V0 float64 = 0.0
	a  float64 = 0.0
	t  float64 = 0.0
)

func main() {
	fmt.Println("Calculator of kinematics")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the initial velocity (V0): ")
	V0 = readFloat(reader)

	fmt.Print("Enter acceleration (a): ")
	a = readFloat(reader)

	fmt.Print("Enter time (t): ")
	t = readFloat(reader)

	calculateKinematics()

	fmt.Printf("Result: V = %.2f\n", V)
}

func calculateKinematics() {
	V = V0 + (a * t)
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
