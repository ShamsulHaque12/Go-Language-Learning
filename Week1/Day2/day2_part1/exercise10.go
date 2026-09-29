package main
/*
// Panic vs error

import (
	"errors"
	"fmt"
)

func divideNumber(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func changeAge(age int) int {
	if age < 0 {
		panic("Age can not be negative")
	}
	return age
}

func main() {
	// 1. Defer MUST be declared BEFORE any statement that might cause a panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("============= Exercise 10 =============")

	fmt.Print("Enter the First Number : ")
	var a int
	fmt.Scanln(&a)

	fmt.Print("Enter the Second Number : ")
	var b int
	fmt.Scanln(&b)

	// 2. Handle normal error from divideNumber
	res, err := divideNumber(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	fmt.Print("Enter the Age : ")
	var age int
	fmt.Scanln(&age)

	// 3. Call changeAge (will trigger panic if age < 0)
	updatedAge := changeAge(age)
	fmt.Println("Updated Age:", updatedAge)
}

*/
