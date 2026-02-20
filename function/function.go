package main

import "fmt"

// func main() {
// 	a := 10
// 	b := 12
// 	sum := a + b

// 	fmt.Println(sum)
// }

/// function...........

func add(num1 int, num2 int) {
	sum := num1 + num2

	fmt.Println(sum)
}
func minus(num1 int, num2 int) {
	minu := num1 - num2

	fmt.Println(minu)
}

func main() {
	a := 122
	b := 22

	add(a, b)
	minus(a, b)
}
