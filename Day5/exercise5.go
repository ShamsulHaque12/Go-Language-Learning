package main

/*

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

// Helper function to save users slice to JSON file
func saveUsers(filename string, users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Helper function to read users slice from JSON file
func loadUsers(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func main() {
	fmt.Println("===== JSON File CRUD Operations =====")

	filename := "users.json"

	// ----------------------------------------------------
	// 1. CREATE
	// ----------------------------------------------------
	fmt.Println("\n====== 1. CREATE ======")
	users := []User{
		{
			ID:    1,
			Name:  "Sujon",
			Email: "sujon@gmail.com",
			Age:   27,
			Role:  "Flutter Developer",
		},
		{
			ID:    2,
			Name:  "Rahim",
			Email: "rahim@gmail.com",
			Age:   25,
			Role:  "Go Developer",
		},
		{
			ID:    3,
			Name:  "Karim",
			Email: "karim@gmail.com",
			Age:   30,
			Role:  "Backend Engineer",
		},
	}

	err := saveUsers(filename, users)
	if err != nil {
		fmt.Println("Error creating/saving users:", err)
		return
	}
	fmt.Println("Created initial users and saved to", filename)

	// ----------------------------------------------------
	// 2. READ
	// ----------------------------------------------------
	fmt.Println("\n====== 2. READ ======")
	readUsers, err := loadUsers(filename)
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	}

	fmt.Println("Loaded Users from file:")
	for _, u := range readUsers {
		fmt.Printf("ID: %d | Name: %-6s | Email: %-16s | Age: %d | Role: %s\n", u.ID, u.Name, u.Email, u.Age, u.Role)
	}

	// ----------------------------------------------------
	// 3. UPDATE
	// ----------------------------------------------------
	fmt.Println("\n====== 3. UPDATE ======")
	targetIDToUpdate := 2
	updated := false

	// Update user with ID 2
	for i, u := range readUsers {
		if u.ID == targetIDToUpdate {
			readUsers[i].Age = 26
			readUsers[i].Role = "Senior Go Developer"
			updated = true
			break
		}
	}

	if updated {
		err = saveUsers(filename, readUsers)
		if err != nil {
			fmt.Println("Error saving updated users:", err)
			return
		}
		fmt.Printf("User with ID %d updated successfully!\n", targetIDToUpdate)
	} else {
		fmt.Printf("User with ID %d not found for update.\n", targetIDToUpdate)
	}

	// Read and display after update
	readUsers, _ = loadUsers(filename)
	fmt.Println("Users after Update:")
	for _, u := range readUsers {
		fmt.Printf("ID: %d | Name: %-6s | Email: %-16s | Age: %d | Role: %s\n", u.ID, u.Name, u.Email, u.Age, u.Role)
	}

	// ----------------------------------------------------
	// 4. DELETE
	// ----------------------------------------------------
	fmt.Println("\n====== 4. DELETE ======")
	targetIDToDelete := 3
	var updatedUsers []User
	deleted := false

	for _, u := range readUsers {
		if u.ID == targetIDToDelete {
			deleted = true
			continue // skip this user (delete)
		}
		updatedUsers = append(updatedUsers, u)
	}

	if deleted {
		err = saveUsers(filename, updatedUsers)
		if err != nil {
			fmt.Println("Error saving after delete:", err)
			return
		}
		fmt.Printf("User with ID %d deleted successfully!\n", targetIDToDelete)
	} else {
		fmt.Printf("User with ID %d not found for deletion.\n", targetIDToDelete)
	}

	// Read and display final list after delete
	finalUsers, _ := loadUsers(filename)
	fmt.Println("\nFinal Users List in file:")
	for _, u := range finalUsers {
		fmt.Printf("ID: %d | Name: %-6s | Email: %-16s | Age: %d | Role: %s\n", u.ID, u.Name, u.Email, u.Age, u.Role)
	}
}

*/
