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
		{
			Name:   "Final Velocity",
			Find:   "V",
			Needed: []string{"V0", "a", "t"},
			Compute: func(args map[string]float64) {
				v0 := args["V0"]
				a := args["a"]
				t := args["t"]
				result := v0 + (a * t)
				fmt.Printf("Result: V = %.2f\n", result)
			},
		},
		{
			Name:   "Final Velocity from Distance and Time",
			Find:   "V",
			Needed: []string{"S", "t"},
			Compute: func(args map[string]float64) {
				s := args["S"]
				t := args["t"]
				if t == 0 {
					fmt.Println("Error: time cannot be zero!")
					return
				}
				result := s / t
				fmt.Printf("Result: V = %.2f\n", result)
			},
		},
		{
			Name:   "Acceleration",
			Find:   "a",
			Needed: []string{"V", "V0", "t"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				v0 := args["V0"]
				t := args["t"]
				if t == 0 {
					fmt.Println("Error: time cannot be zero!")
					return
				}
				result := (v - v0) / t
				fmt.Printf("Result: a = %.2f\n", result)
			},
		},
		{
			Name:   "Distance",
			Find:   "S",
			Needed: []string{"V", "t"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				t := args["t"]
				result := v * t
				fmt.Printf("Result: S = %.2f\n", result)
			},
		},
		{
			Name:   "Distance from Velocities and Acceleration",
			Find:   "S",
			Needed: []string{"V", "V0", "a"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				v0 := args["V0"]
				a := args["a"]
				if a == 0 {
					fmt.Println("Error: acceleration cannot be zero!")
					return
				}
				result := ((v * v) - (v0 * v0)) / (2 * a)
				fmt.Printf("Result: S = %.2f\n", result)
			},
		},
		{
			Name:   "Time",
			Find:   "t",
			Needed: []string{"S", "V"},
			Compute: func(args map[string]float64) {
				s := args["S"]
				v := args["V"]
				if v == 0 {
					fmt.Println("Error: velocity cannot be zero!")
					return
				}
				result := s / v
				fmt.Printf("Result: t = %.2f\n", result)
			},
		},
		{
			Name:   "Initial Velocity",
			Find:   "V0",
			Needed: []string{"V", "a", "t"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				a := args["a"]
				t := args["t"]
				result := v - (a * t)
				fmt.Printf("Result: V0 = %.2f\n", result)
			},
		},
	}

	fmt.Print("What needs to be found? (V, a, S, t, V0): ")
	input, _ := reader.ReadString('\n')
	target := strings.TrimSpace(input)

	var foundFormula *Formula
	for i := range formulas {
		if formulas[i].Find == target {
			foundFormula = &formulas[i]
			break
		}
	}

	if foundFormula == nil {
		fmt.Println("Formula not found")
		return
	}

	fmt.Printf("Using formula: %s\n", foundFormula.Name)

	args := make(map[string]float64)
	for _, param := range foundFormula.Needed {
		fmt.Printf("Enter %s: ", param)
		args[param] = readFloat(reader)
	}

	foundFormula.Compute(args)
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
