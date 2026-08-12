package main

import "fmt"

func main() {
	x := 10
	if x > 12 {
		fmt.Println("you are good")
	} else if x < 5 {
		fmt.Println("you are bad")
	} else {
		fmt.Println("you are average")
	}
}
