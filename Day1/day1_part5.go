/* ...............Array , Slice , Map.............. */

package main

import "fmt"

/*
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
*/

/*
func main() {

	fmt.Println("Task 2: Slice")
	fmt.Println("=========Slice========= ")

	numer := []int{10, 20, 30, 40, 50}

	add_number := append(numer, 60, 100)

	remove_number := numer[0:2]

	fmt.Println()
	fmt.Println("Slice Numbers: ", numer)
	for i := 0; i < len(numer); i++ {
		fmt.Println("Index: ", i+1, "==>", "Value: ", numer[i])
	}
	fmt.Println()

	fmt.Println("After Appending: ", add_number)
	for i, a_number := range add_number {
		fmt.Println("Index: ", i+1, "==>", "Value: ", a_number)
	}
	fmt.Println()

	fmt.Println("After Removing: ", remove_number)
	for i, r_number := range remove_number {
		fmt.Println("Index: ", i+1, "==>", "Value: ", r_number)
	}
	fmt.Println()
}
*/

func main() {

	fmt.Println("Task 3: Map")
	fmt.Println("=========Map========= ")

	fmt.Println()
	student := map[string]string{
		"name":    "Sujon",
		"email":   "xyz@gmail.com",
		"phone":   "01700000000",
		"country": "Bangladesh",
	}
	fmt.Println(student)
	fmt.Println()

	fmt.Println("Student Details:")
	for i, a_number := range student {
		fmt.Println(i, ": ", a_number)
	}

	fmt.Println()
	fmt.Println("After Adding Country:")
	student["country"] = "Indonesia"
	fmt.Println(student)
	fmt.Println()

	fmt.Println("After Removing Country:")
	delete(student, "country")
	fmt.Println(student)
	fmt.Println()
}