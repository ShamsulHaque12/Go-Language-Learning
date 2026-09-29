package main

/*
import (
	"errors"
	"fmt"
)

// Define sentinel error for user not found
var errUserNotFound = errors.New("user not found")

type User struct {
	ID   int
	Name string
	Age  int
}

func findUserByID(id int) (User, error) {
	users := []User{
		{ID: 1, Name: "Shamsul", Age: 24},
		{ID: 2, Name: "Sujon", Age: 24},
		{ID: 3, Name: "Zihad", Age: 24},
		{ID: 4, Name: "Rafi", Age: 25},
		{ID: 5, Name: "Nabil", Age: 24},
	}

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}

	// Fix: pass errUserNotFound directly (not pointer &errUserNotFound)
	// %w wraps the error so it can be checked using errors.Is()
	return User{}, fmt.Errorf("failed to find user with ID %d && %w", id, errUserNotFound)
}

func main() {
	fmt.Println("========= Exercise 8 ==========")

	// 1. Success Case: Searching for an existing user (ID: 3)
	user, err := findUserByID(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User Found:", user)
	}
	fmt.Println()

	// 2. Error Case: Searching for a non-existent user (ID: 6)
	user, err = findUserByID(6)
	if err != nil {
		fmt.Println("Error:", err)

		// Demonstrate how errors.Is works with wrapped errors (%w)
		if errors.Is(err, errUserNotFound) {
			fmt.Println("Note: Root cause is 'errUserNotFound'")
		}
	} else {
		fmt.Println("User Found:", user)
	}
}

*/
