package main

import "fmt"

func main() {
	const (
		name        = "Sujon"
		age         = 27
		email       = "sujon@gmail.com"
		department  = "CSE"
		university  = "BUBT"
		gpa         = 3.00
		isGraduated = true
	)

	fmt.Println("Name is", name)
	fmt.Println("Age is", age)
	fmt.Println("Email is", email)
	fmt.Println("Department is", department)
	fmt.Println("University is", university)
	fmt.Println("GPA is", gpa)
	fmt.Println("Is Graduated?", isGraduated)
}
