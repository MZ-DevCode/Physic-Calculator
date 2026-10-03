package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Formula struct {
	Name     string                        //name
	Equation string                        //equation
	Find     string                        //need find value
	Needed   []string                      //needed for formula
	Compute  func(args map[string]float64) //math
}

func main() {
	fmt.Println("Calculator of kinematics & Free Fall")

	reader := bufio.NewReader(os.Stdin)

	units := map[string]string{
		"V":  "m/s",
		"V0": "m/s",
		"a":  "m/s^2",
		"t":  "s",
		"S":  "m",
		"h":  "m",
		"g":  "m/s^2",
	}

	formulas := []Formula{
		{
			Name:     "Final Velocity",
			Equation: "V = V0 + a * t",
			Find:     "V",
			Needed:   []string{"V0", "a", "t"},
			Compute: func(args map[string]float64) {
				v0 := args["V0"]
				a := args["a"]
				t := args["t"]
				result := v0 + (a * t)
				fmt.Printf("Result: V = %.2f %s\n", result, units["V"])
			},
		},
		{
			Name:     "Final Velocity from Distance and Time",
			Equation: "V = S / t",
			Find:     "V",
			Needed:   []string{"S", "t"},
			Compute: func(args map[string]float64) {
				s := args["S"]
				t := args["t"]
				if t == 0 {
					fmt.Println("Error: time cannot be zero!")
					return
				}
				result := s / t
				fmt.Printf("Result: V = %.2f %s\n", result, units["V"])
			},
		},
		{
			Name:     "Acceleration",
			Equation: "a = (V - V0) / t",
			Find:     "a",
			Needed:   []string{"V", "V0", "t"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				v0 := args["V0"]
				t := args["t"]
				if t == 0 {
					fmt.Println("Error: time cannot be zero!")
					return
				}
				result := (v - v0) / t
				fmt.Printf("Result: a = %.2f %s\n", result, units["a"])
			},
		},
		{
			Name:     "Distance",
			Equation: "S = V * t",
			Find:     "S",
			Needed:   []string{"V", "t"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				t := args["t"]
				result := v * t
				fmt.Printf("Result: S = %.2f %s\n", result, units["S"])
			},
		},
		{
			Name:     "Distance from Velocities and Acceleration",
			Equation: "S = (V^2 - V0^2) / (2 * a)",
			Find:     "S",
			Needed:   []string{"V", "V0", "a"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				v0 := args["V0"]
				a := args["a"]
				if a == 0 {
					fmt.Println("Error: acceleration cannot be zero!")
					return
				}
				result := ((v * v) - (v0 * v0)) / (2 * a)
				fmt.Printf("Result: S = %.2f %s\n", result, units["S"])
			},
		},
		{
			Name:     "Time",
			Equation: "t = S / V",
			Find:     "t",
			Needed:   []string{"S", "V"},
			Compute: func(args map[string]float64) {
				s := args["S"]
				v := args["V"]
				if v == 0 {
					fmt.Println("Error: velocity cannot be zero!")
					return
				}
				result := s / v
				fmt.Printf("Result: t = %.2f %s\n", result, units["t"])
			},
		},
		{
			Name:     "Initial Velocity",
			Equation: "V0 = V - a * t",
			Find:     "V0",
			Needed:   []string{"V", "a", "t"},
			Compute: func(args map[string]float64) {
				v := args["V"]
				a := args["a"]
				t := args["t"]
				result := v - (a * t)
				fmt.Printf("Result: V0 = %.2f %s\n", result, units["V0"])
			},
		},
		{
			Name:     "Free Fall Height (from time)",
			Equation: "h = (g * t^2) / 2",
			Find:     "h",
			Needed:   []string{"g", "t"},
			Compute: func(args map[string]float64) {
				g := args["g"]
				t := args["t"]
				result := (g * (t * t)) / 2
				fmt.Printf("Result: h = %.2f %s\n", result, units["h"])
			},
		},
		{
			Name:     "Free Fall Time (from height)",
			Equation: "t = sqrt(2 * h / g)",
			Find:     "t_fall",
			Needed:   []string{"h", "g"},
			Compute: func(args map[string]float64) {
				h := args["h"]
				g := args["g"]
				if g == 0 {
					fmt.Println("Error: gravity cannot be zero!")
					return
				}
				result := math.Sqrt((2 * h) / g)
				fmt.Printf("Result: t = %.2f %s\n", result, units["t"])
			},
		},
		{
			Name:     "Free Fall Final Velocity",
			Equation: "V = g * t",
			Find:     "V_ff",
			Needed:   []string{"g", "t"},
			Compute: func(args map[string]float64) {
				g := args["g"]
				t := args["t"]
				result := g * t
				fmt.Printf("Result: V = %.2f %s\n", result, units["V"])
			},
		},
	}

	fmt.Print("What needs to be found? (V, a, S, t, V0, h, t_fall, V_ff): ")
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

	fmt.Printf("Using formula: %s (%s)\n", foundFormula.Name, foundFormula.Equation)

	args := make(map[string]float64)
	for _, param := range foundFormula.Needed {
		if unit, exists := units[param]; exists {
			fmt.Printf("Enter %s (%s): ", param, unit)
		} else {
			fmt.Printf("Enter %s: ", param)
		}
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
