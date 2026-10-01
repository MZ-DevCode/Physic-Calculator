package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Formula struct {
	Name    string                        //name
	Find    string                        //need find value
	Needed  []string                      //needed for formula
	Compute func(args map[string]float64) //math
}

func main() {
	fmt.Println("Calculator of kinematics")

	reader := bufio.NewReader(os.Stdin)
	formulas := []Formula{
		Name:	"Final Velocity",
		Find:	"V",
		Needed: []string{"V0", "a", "t"},
		Compute: func(args map[string]float64){
			V0 := args["V0"]
			a := args["a"]
			t := args["t"]
			result := V0 + (a * t)
			fmt.Printf("Result: V = %.2f\n", result)
		}
	}

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
