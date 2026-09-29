package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"exercise2/calculator"
	"exercise2/user"
)

func main() {
	fmt.Println("===== User & Calculator Demo =====")

	// User package usage
	u1 := user.NewUser(1, "John Doe", "doe@gmail.com")
	fmt.Println("User 1:", u1.GetID(), u1.GetName(), u1.GetEmail())

	u2 := user.NewUser(2, "Jane Doe", "jane@gmail.com")
	fmt.Println("User 2:", u2.GetID(), u2.GetName(), u2.GetEmail())

	// Calculator package usage
	fmt.Println("\nApp Name:", calculator.AppName)

	sum := calculator.Add(15, 25)
	fmt.Println("Addition:", sum)

	sub := calculator.Sub(15, 25)
	fmt.Println("Subtraction:", sub)

	mul := calculator.Mul(15, 25)
	fmt.Println("Multiplication:", mul)

	result, err := calculator.Div(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

	// Gin server
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Go Backend",
		})
	})

	router.Run(":8080")
}