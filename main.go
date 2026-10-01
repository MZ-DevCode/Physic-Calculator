package main

import "fmt"

var (
	V  = 0
	V0 = 0
	a  = 0
	t  = 0
	g  = 10.0
)

func main() {
	fmt.Println("Calculator of cinematics")
	fmt.Print("Enter the initial velocity(V0): ")
	fmt.Scan(&V0)

	fmt.Print("Enter acceleration(a): ")
	fmt.Scan(&a)

	fmt.Print("Enter time(t): ")
	fmt.Scan(&t)

	fmt.Printf("Result: V = %.2f\n", V)
}

func calculate() {
	V = V0 + (a * t)
}
