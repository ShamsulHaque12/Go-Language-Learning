package main

import "fmt"

func main() {
	fmt.Println("welcome go language")

	var name string
	fmt.Println("Enter your name ___")
	fmt.Scanln(&name)

	var num1 int
	var num2 int
	fmt.Println("emter your 1st numer: ____")
	fmt.Println("enter your 2nd number:____")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	sum := num1 - num2

	fmt.Println("result: ", sum)
}
