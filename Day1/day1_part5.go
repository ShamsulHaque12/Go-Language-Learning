/* ...............Array , Slice , Map.............. */

package main

import "fmt"

func main() {

	fmt.Println("Task 1: Array")
	fmt.Println("=========Array========= ")

	number := [5]int{10, 20, 30, 40, 50}

	number[4] = 100

	fmt.Println(number[4])

	for i := 0; i < len(number); i++ {
		fmt.Println("Index: ", i+1, "==>", "Value: ", number[i])
	} 
	
}