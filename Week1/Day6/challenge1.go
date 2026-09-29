package main

/*
import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Active bool
}

func main() {
	fmt.Println("===== Challenge 1: User Validation & Processing ===========")

	// Raw input strings
	rawName := "   Sujon Islam   "
	rawEmail := "   SUJON@GMAIL.COM   "
	rawAge := " 25 "
	rawActive := " true "

	fmt.Println("\n--- Raw Inputs ---")
	fmt.Printf("Raw Name  : %q\n", rawName)
	fmt.Printf("Raw Email : %q\n", rawEmail)
	fmt.Printf("Raw Age   : %q\n", rawAge)
	fmt.Printf("Raw Active: %q\n", rawActive)

	// 1. Clean Name
	cleanName := strings.TrimSpace(rawName)
	if cleanName == "" {
		fmt.Println("\nError: Name cannot be empty!")
		return
	}

	// 2. Clean Email
	cleanEmail := strings.ToLower(strings.TrimSpace(rawEmail))

	// 3. Gmail Check
	isGmail := strings.HasSuffix(cleanEmail, "@gmail.com")
	if !isGmail {
		fmt.Println("\nError: Email must be a valid @gmail.com address!")
		return
	}

	// 4. Convert Age (string -> int)
	age, err := strconv.Atoi(strings.TrimSpace(rawAge))
	if err != nil {
		fmt.Println("\nError: Invalid age input!", err)
		return
	}

	// 5. Convert Active (string -> bool)
	active, err := strconv.ParseBool(strings.TrimSpace(rawActive))
	if err != nil {
		fmt.Println("\nError: Invalid active state!", err)
		return
	}

	// Create User struct if all validations pass
	user := User{
		Name:   cleanName,
		Email:  cleanEmail,
		Age:    age,
		Active: active,
	}

	// Output validated user information nicely
	fmt.Println("\n===== Validated User Profile =====")
	fmt.Printf("Name    : %s\n", user.Name)
	fmt.Printf("Email   : %s\n", user.Email)
	fmt.Printf("Age     : %d\n", user.Age)
	fmt.Printf("Active  : %t\n", user.Active)
	fmt.Printf("Is Gmail: %t\n", isGmail)
	fmt.Printf("Struct  : %+v\n", user)
}

*/