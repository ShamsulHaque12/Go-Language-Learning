package main

import "fmt"

var (
	a int = 10
	b int = 20
)
func totalNumber(num int)  {
	fmt.Println("total number is ", num)
}
func add(a int , b int)  {
	result := a + b
	totalNumber(result)
}
func main()  {
	add(a, b)
}