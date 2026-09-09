package main

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var users = []User{
	{ID: 1, Name: "Shamsul", Email: "shamsul@example.com"},
	{ID: 2, Name: "Sujon", Email: "sujon@example.com"},
	{ID: 3, Name: "Zihad", Email: "zihad@example.com"},
	{ID: 4, Name: "Rafi", Email: "rafi@example.com"},
	{ID: 5, Name: "Nabil", Email: "nabil@example.com"},
}

func findUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID %d not found", id)
}

func findUserByName(name string) (User, error) {
	for _, u := range users {
		if u.Name == name {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with name '%s' not found", name)
}

func displayUser(u User) {
	fmt.Println("ID:   ", u.ID)
	fmt.Println("Name: ", u.Name)
	fmt.Println("Email:", u.Email)
}

func displayAllUsers() {
	fmt.Println("--- All Users List ---")
	for _, u := range users {
		fmt.Printf("ID: %d | Name: %-8s | Email: %s\n", u.ID, u.Name, u.Email)
	}
	fmt.Println("----------------------")
}

func main() {
	fmt.Println("============= Exercise 7 =============")

	// 1. Initial display of all users
	displayAllUsers()
	fmt.Println()

	// 2. User choice for search type
	var choice int
	fmt.Println("Search options:")
	fmt.Println("1. Search by ID")
	fmt.Println("2. Search by Name")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var searchID int
		fmt.Print("Enter User ID: ")
		fmt.Scanln(&searchID)

		user, err := findUserByID(searchID)
		if err != nil {
			fmt.Println("\nError:", err)
		} else {
			fmt.Println("\nUser Found:")
			displayUser(user)
		}

	case 2:
		var searchName string
		fmt.Print("Enter User Name: ")
		fmt.Scanln(&searchName)

		user, err := findUserByName(searchName)
		if err != nil {
			fmt.Println("\nError:", err)
		} else {
			fmt.Println("\nUser Found:")
			displayUser(user)
		}

	default:
		fmt.Println("\nInvalid option selected!")
	}
}
