package main

import "fmt"

/// This program checks if a person is an adult or not based on their age................
// func main() {
// 	age := 20

// 	if age >= 18 {
// 		fmt.Println("Access Granted.")
// 	} else {
// 		fmt.Println("Access Denied.")
// 	}
// }

/// positive or negative number check..................................
// func main() {
// 	n := -5
// 	if n > 0 {
// 		fmt.Println("The number is positive.")
// 	} else if n < 0 {
// 		fmt.Println("The number is negative.")
// 	} else {
// 		fmt.Println("The number is zero.")
// 	}
// }

/// leap year check...........................................
// func main() {
// 	year := 2020
// 	if year%4 == 0 {
// 		fmt.Println("The year is a leap year.")
// 	} else {
// 		fmt.Println("The year is not a leap year.")
// 	}
// }

/// leap year check with more conditions...........................
// func main() {
// 	year := 2000
// 	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
// 		fmt.Println("The year is a leap year.")
// 	} else {
// 		fmt.Println("The year is not a leap year.")
// 	}
// }

// / shopping discount based on total amount...........................
func main() {
	amount := 4000.50
	if amount > 5000 {
		fmt.Println("You get a 10% discount.")
	} else if 2000 < amount && amount < 5000 {
		fmt.Println("You get a 5% discount.")
	} else if amount < 2000 {
		fmt.Println("You get no discount.")
	}
}
