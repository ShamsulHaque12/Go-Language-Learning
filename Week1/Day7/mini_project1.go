package main

/*
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

const dataFile = "users.json"

var users []User

// Helper: read trimmed line from input
func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// 7. Save Users -> JSON
func saveUsersToFile() error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

// 8. Load Users <- JSON
func loadUsersFromFile() error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			users = []User{}
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

// Helper: Generate next auto-increment ID
func getNextID() int {
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	return maxID + 1
}

// 1. Create User
func createUser(reader *bufio.Reader) {
	fmt.Println("\n--- 1. Create New User ---")
	name := readInput(reader, "Enter Name: ")
	if name == "" {
		fmt.Println("❌ Error: Name cannot be empty!")
		return
	}

	ageStr := readInput(reader, "Enter Age: ")
	age, err := strconv.Atoi(ageStr)
	if err != nil || age <= 0 {
		fmt.Println("❌ Error: Invalid age! Age must be a positive number.")
		return
	}

	email := strings.ToLower(readInput(reader, "Enter Email: "))
	if email == "" || !strings.Contains(email, "@") {
		fmt.Println("❌ Error: Invalid email address!")
		return
	}

	role := readInput(reader, "Enter Role (e.g. Developer, Admin): ")
	if role == "" {
		role = "User"
	}

	activeStr := readInput(reader, "Is Active? (true/false): ")
	isActive, err := strconv.ParseBool(activeStr)
	if err != nil {
		isActive = true
	}

	newUser := User{
		ID:       getNextID(),
		Name:     name,
		Age:      age,
		Email:    email,
		Role:     role,
		IsActive: isActive,
	}

	users = append(users, newUser)

	if err := saveUsersToFile(); err != nil {
		fmt.Println("⚠️ Warning: Failed to save to file:", err)
	} else {
		fmt.Printf("✅ User %q (ID: %d) created and saved successfully!\n", newUser.Name, newUser.ID)
	}
}

// 2. Get All Users
func getAllUsers() {
	fmt.Println("\n--- 2. All Users ---")
	if len(users) == 0 {
		fmt.Println("ℹ️ No users found.")
		return
	}

	fmt.Printf("%-5s %-20s %-5s %-25s %-15s %-8s\n", "ID", "Name", "Age", "Email", "Role", "Active")
	fmt.Println(strings.Repeat("-", 80))
	for _, u := range users {
		fmt.Printf("%-5d %-20s %-5d %-25s %-15s %-8t\n", u.ID, u.Name, u.Age, u.Email, u.Role, u.IsActive)
	}
	fmt.Println(strings.Repeat("-", 80))
}

// 3. Get User By ID
func getUserByID(reader *bufio.Reader) {
	fmt.Println("\n--- 3. Get User By ID ---")
	idStr := readInput(reader, "Enter User ID: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Error: Invalid ID format!")
		return
	}

	for _, u := range users {
		if u.ID == id {
			fmt.Println("\n✅ User Found:")
			fmt.Printf("ID       : %d\n", u.ID)
			fmt.Printf("Name     : %s\n", u.Name)
			fmt.Printf("Age      : %d\n", u.Age)
			fmt.Printf("Email    : %s\n", u.Email)
			fmt.Printf("Role     : %s\n", u.Role)
			fmt.Printf("IsActive : %t\n", u.IsActive)
			return
		}
	}

	fmt.Printf("❌ Error: User with ID %d not found.\n", id)
}

// 4. Update User
func updateUser(reader *bufio.Reader) {
	fmt.Println("\n--- 4. Update User ---")
	idStr := readInput(reader, "Enter User ID to update: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Error: Invalid ID format!")
		return
	}

	foundIndex := -1
	for i, u := range users {
		if u.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("❌ Error: User with ID %d not found.\n", id)
		return
	}

	u := &users[foundIndex]
	fmt.Printf("Updating user: %s (Press Enter to keep current value)\n", u.Name)

	newName := readInput(reader, fmt.Sprintf("Enter Name [%s]: ", u.Name))
	if newName != "" {
		u.Name = newName
	}

	newAgeStr := readInput(reader, fmt.Sprintf("Enter Age [%d]: ", u.Age))
	if newAgeStr != "" {
		if age, err := strconv.Atoi(newAgeStr); err == nil && age > 0 {
			u.Age = age
		}
	}

	newEmail := readInput(reader, fmt.Sprintf("Enter Email [%s]: ", u.Email))
	if newEmail != "" && strings.Contains(newEmail, "@") {
		u.Email = strings.ToLower(newEmail)
	}

	newRole := readInput(reader, fmt.Sprintf("Enter Role [%s]: ", u.Role))
	if newRole != "" {
		u.Role = newRole
	}

	newActiveStr := readInput(reader, fmt.Sprintf("Is Active? [%t]: ", u.IsActive))
	if newActiveStr != "" {
		if active, err := strconv.ParseBool(newActiveStr); err == nil {
			u.IsActive = active
		}
	}

	if err := saveUsersToFile(); err != nil {
		fmt.Println("⚠️ Warning: Failed to save updated user:", err)
	} else {
		fmt.Println("✅ User updated successfully!")
	}
}

// 5. Delete User
func deleteUser(reader *bufio.Reader) {
	fmt.Println("\n--- 5. Delete User ---")
	idStr := readInput(reader, "Enter User ID to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Error: Invalid ID format!")
		return
	}

	foundIndex := -1
	for i, u := range users {
		if u.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("❌ Error: User with ID %d not found.\n", id)
		return
	}

	deletedName := users[foundIndex].Name
	users = append(users[:foundIndex], users[foundIndex+1:]...)

	if err := saveUsersToFile(); err != nil {
		fmt.Println("⚠️ Warning: Failed to save after deletion:", err)
	} else {
		fmt.Printf("✅ User %q (ID: %d) deleted successfully!\n", deletedName, id)
	}
}

// 6. Search User
func searchUser(reader *bufio.Reader) {
	fmt.Println("\n--- 6. Search User ---")
	query := strings.ToLower(readInput(reader, "Enter search keyword (Name or Email): "))
	if query == "" {
		fmt.Println("❌ Error: Search query cannot be empty!")
		return
	}

	var results []User
	for _, u := range users {
		if strings.Contains(strings.ToLower(u.Name), query) || strings.Contains(strings.ToLower(u.Email), query) {
			results = append(results, u)
		}
	}

	if len(results) == 0 {
		fmt.Printf("ℹ️ No users found matching %q.\n", query)
		return
	}

	fmt.Printf("\nFound %d matching user(s):\n", len(results))
	fmt.Printf("%-5s %-20s %-5s %-25s %-15s %-8s\n", "ID", "Name", "Age", "Email", "Role", "Active")
	fmt.Println(strings.Repeat("-", 80))
	for _, u := range results {
		fmt.Printf("%-5d %-20s %-5d %-25s %-15s %-8t\n", u.ID, u.Name, u.Age, u.Email, u.Role, u.IsActive)
	}
	fmt.Println(strings.Repeat("-", 80))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Load existing data on startup
	if err := loadUsersFromFile(); err != nil {
		fmt.Println("⚠️ Warning loading users file:", err)
	} else {
		fmt.Printf("📂 Loaded %d user(s) from %s\n", len(users), dataFile)
	}

	for {
		fmt.Println("\n=========================================")
		fmt.Println("  Week 1 Mini Project: User Management  ")
		fmt.Println("=========================================")
		fmt.Println("1. Create User")
		fmt.Println("2. Get All Users")
		fmt.Println("3. Get User By ID")
		fmt.Println("4. Update User")
		fmt.Println("5. Delete User")
		fmt.Println("6. Search User")
		fmt.Println("7. Save Users -> JSON")
		fmt.Println("8. Load Users <- JSON")
		fmt.Println("9. Exit")
		fmt.Println("=========================================")

		choice := readInput(reader, "Select an option (1-9): ")

		switch choice {
		case "1":
			createUser(reader)
		case "2":
			getAllUsers()
		case "3":
			getUserByID(reader)
		case "4":
			updateUser(reader)
		case "5":
			deleteUser(reader)
		case "6":
			searchUser(reader)
		case "7":
			if err := saveUsersToFile(); err != nil {
				fmt.Println("❌ Error saving file:", err)
			} else {
				fmt.Println("✅ Users saved successfully to", dataFile)
			}
		case "8":
			if err := loadUsersFromFile(); err != nil {
				fmt.Println("❌ Error loading file:", err)
			} else {
				fmt.Printf("✅ Loaded %d user(s) from %s\n", len(users), dataFile)
			}
		case "9":
			fmt.Println("👋 Exiting... Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice! Please select an option between 1 and 9.")
		}
	}
}

*/
