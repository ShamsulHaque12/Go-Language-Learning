
/*
package main

import "fmt"

func main() {
	fmt.Println("BMI Calculator")

	var name string
	var weight float64
	var height float64

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your weight in Kg: ")
	fmt.Scanln(&weight)

	fmt.Print("Enter your height in cm: ")
	fmt.Scanln(&height)

	// Convert cm to meter
	height = height / 100

	// Calculate BMI
	bmi := weight / (height * height)

	fmt.Println()
	fmt.Println("Your name is:", name)
	fmt.Println("Your weight is:", weight, "Kg")
	fmt.Println("Your height is:", height, "m")
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("You are Underweight")
	} else if bmi < 25 {
		fmt.Println("You are Normal weight")
	} else if bmi < 30 {
		fmt.Println("You are Overweight")
	} else {
		fmt.Println("You are Obesity")
	}
}
*/

/*
package main

import "fmt"

func main() {
	
	fmt.Println("Grade Calculator")

	var name string
	var mark int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your mark: ")
	fmt.Scanln(&mark)

	fmt.Println()
	fmt.Println("Your name is:", name)
	fmt.Println("Your mark is:", mark)

	if mark > 100 || mark < 0 {
	fmt.Println("Invalid Mark")
} else if mark >= 80 {
	fmt.Println("A+")
} else if mark >= 70 {
	fmt.Println("A")
} else if mark >= 60 {
	fmt.Println("A-")
} else if mark >= 50 {
	fmt.Println("B")
} else if mark >= 40 {
	fmt.Println("C")
} else if mark >= 33 {
	fmt.Println("D")
} else {
	fmt.Println("F")
}
}
*/

/*
package main

import "fmt"

func main() {

	correctUsername := "sujon"
	correctPassword := "123456"

	fmt.Println("Login System")

	var username string
	var password string

	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	fmt.Println()
	fmt.Println("Your username is:", username)
	fmt.Println("Correct username is:", correctUsername)

	if username == correctUsername && password == correctPassword {
		fmt.Println(" Login successful \n Welcome to the system \n You are a valid user")
	} else {
		fmt.Println("Login failed \n You are not a valid user")
	}
	
}
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var choice int

	for {
		fmt.Println("\n======================================")
		fmt.Println("  Menu Calculator")
		fmt.Println("======================================")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Modulus")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 6 {
			fmt.Println("Exiting... Thank you!")
			break
		}

		switch choice {
		case 1:
			fmt.Println("======== Addition Selected ========")
			var num1, num2, sum float64

			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)

			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			sum = num1 + num2
			fmt.Println("Sum of", num1, "and", num2, "is", sum)

		case 2:
			fmt.Println("======== Subtraction Selected ========")
			var num1, num2, difference float64

			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)

			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			difference = num1 - num2
			fmt.Println("Difference of", num1, "and", num2, "is", difference)

		case 3:
			fmt.Println("======== Multiplication Selected ========")
			var num1, num2, product float64

			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)

			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			product = num1 * num2
			fmt.Println("Product of", num1, "and", num2, "is", product)

		case 4:
			fmt.Println("======== Division Selected ========")
			var num1, num2, quotient float64

			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)

			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			// Division validation using if condition
			if num2 == 0 {
				fmt.Println("Error: Cannot divide by zero (0)!")
			} else {
				quotient = num1 / num2
				fmt.Println("Quotient of", num1, "and", num2, "is", quotient)
			}

		case 5:
			fmt.Println("======== Modulus Selected ========")
			var num1, num2, modulus float64

			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)

			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			if num2 == 0 {
				fmt.Println("Error: Cannot calculate modulus with zero (0)!")
			} else {
				modulus = math.Mod(num1, num2)
				fmt.Println("Modulus of", num1, "and", num2, "is", modulus)
			}

		default:
			fmt.Println("Invalid choice! Please select between 1 and 6.")
		}
	}
}