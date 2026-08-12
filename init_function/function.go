package main

import "fmt"

// var a = 20

// func main(){
// 	fmt.Println("hello init function...........", a)
// }

// func init(){
// 	fmt.Println("hello 2nd init function............")
// }

// func init(){
// 	fmt.Println("hello 3nd init function............", a)
// 	a= 50
// }

/// anonymous function and LIFE..................................

//standerd function
func add(a, b int) {
	fmt.Println(a + b)
}
func main() {
	fmt.Println("hello main function.............")
	add(10, 20)

	// anonymous function
	func(a int, b int) {
		c := a + b
		fmt.Println("anonymous function calling.............", c)
	}(50, 20)

	// if expression
	d := 10
	if d > 0 {
		fmt.Println("d is greater than ", d)
	} else {
		fmt.Println("d is less than or equal to ", d)
	}

	// funstion expression
	minus := func(a int, b int) {
		c := a - b
		fmt.Println("minus function..............", c)
	}

	minus(40, 20)

	// IIFE
	func(a int, b int){
		d := a *b
		fmt.Println("IIFE function.............",d)
	}(10,20)
	
}

func init() {
	fmt.Println("hello anonymous init function.............")
}