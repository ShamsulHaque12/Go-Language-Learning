package main

import (
	"fmt"
	"exercise1/calculator" // আমাদের তৈরি custom package import করা হলো
)

func main() {
	fmt.Println("=== Go Package & Import Demonstration ===")

	// 1. Exported Variable ব্যবহার (Capital 'A')
	fmt.Println("App Name:", calculator.AppName)

	// 2. Exported Functions ব্যবহার (Capital 'A', 'S', 'M', 'D')
	sum := calculator.Add(10, 5)
	fmt.Println("Addition (10 + 5):", sum)

	sub := calculator.Sub(10, 5)
	fmt.Println("Subtraction (10 - 5):", sub)

	mul := calculator.Mul(10, 5)
	fmt.Println("Multiplication (10 * 5):", mul)

	div, err := calculator.Div(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division (10 / 2):", div)
	}

}
