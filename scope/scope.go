package main

import "fmt"

var a = 26
var b = 45

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	p := 12
	q := 34

	add(p, q)
	add(a, b)
	add(a, p)

}
