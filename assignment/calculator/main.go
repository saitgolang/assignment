package main

import (
	"fmt"
	"calculator/calculatorop"
)

func main() {
	a, b := 10.0, 5.0

	fmt.Printf("Addition: %.2f\n", calculatorop.Add(a, b))
	fmt.Printf("Subtraction: %.2f\n", calculatorop.Subtract(a, b))
	fmt.Printf("Multiplication: %.2f\n", calculatorop.Multiply(a, b))
	fmt.Printf("Division: %.2f\n", calculatorop.Divide(a, b))
}
