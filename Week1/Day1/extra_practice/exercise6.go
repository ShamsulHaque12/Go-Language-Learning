package main

/*
import (
	"fmt"
	"strings"
)

// User represents a user entity
type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

// UserStore manages a collection of users without pointers
type UserStore struct {
	users []User
}

// NewUserStore creates a new UserStore value with initial mock data
func NewUserStore() UserStore {
	return UserStore{
		users: []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", IsActive: true},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com", IsActive: false},
			{ID: 3, Name: "Bob Smith", Email: "bob@example.com", IsActive: true},
			{ID: 4, Name: "Charlie Brown", Email: "charlie@example.com", IsActive: false},
			{ID: 5, Name: "Emily Davis", Email: "emily@example.com", IsActive: false},
		},
	}
}

// Display prints a single user's information cleanly formatted
func (u User) Display() {
	status := "Inactive"
	if u.IsActive {
		status = "Active"
	}
	fmt.Printf("[%d] %-16s | %-22s | Status: %s\n", u.ID, u.Name, u.Email, status)
}

// DisplayAll prints all users in the store
func (us UserStore) DisplayAll() {
	fmt.Println("\n--- All Users ---")
	if len(us.users) == 0 {
		fmt.Println("No users available.")
		return
	}
	for _, u := range us.users {
		u.Display()
	}
}

// DisplayFiltered prints users based on active status
func (us UserStore) DisplayFiltered(activeOnly bool) {
	statusStr := "Active"
	if !activeOnly {
		statusStr = "Inactive"
	}
	fmt.Printf("\n--- %s Users ---\n", statusStr)

	count := 0
	for _, u := range us.users {
		if u.IsActive == activeOnly {
			u.Display()
			count++
		}
	}
	if count == 0 {
		fmt.Printf("No %s users found.\n", strings.ToLower(statusStr))
	}
}

// FindByID searches for a user by ID and returns the User struct value
func (us UserStore) FindByID(id int) (User, bool) {
	for _, u := range us.users {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

// AddUser adds a new user and returns the updated UserStore value
func (us UserStore) AddUser(name, email string, isActive bool) UserStore {
	newID := 1
	if len(us.users) > 0 {
		newID = us.users[len(us.users)-1].ID + 1
	}
	newUser := User{
		ID:       newID,
		Name:     name,
		Email:    email,
		IsActive: isActive,
	}
	us.users = append(us.users, newUser)
	fmt.Printf("User successfully created with ID: %d\n", newID)
	return us
}

// reindexIDs updates the ID of each user sequentially from 1 to N
func (us UserStore) reindexIDs() UserStore {
	for i := range us.users {
		us.users[i].ID = i + 1
	}
	return us
}

// DeleteUser removes a user by ID, re-indexes remaining IDs, and returns the updated UserStore
func (us UserStore) DeleteUser(id int) (UserStore, bool) {
	for i, u := range us.users {
		if u.ID == id {
			us.users = append(us.users[:i], us.users[i+1:]...)
			us = us.reindexIDs()
			return us, true
		}
	}
	return us, false
}

// printMenu displays the system options
func printMenu() {
	fmt.Println("\n==============================================")
	fmt.Println("         USER MANAGEMENT SYSTEM               ")
	fmt.Println("==============================================")
	fmt.Println("1. Display All Users")
	fmt.Println("2. Display Active Users")
	fmt.Println("3. Display Inactive Users")
	fmt.Println("4. Search User by ID")
	fmt.Println("5. Add New User")
	fmt.Println("6. Delete User by ID")
	fmt.Println("7. Exit")
	fmt.Println("==============================================")
	fmt.Print("Enter your choice (1-7): ")
}

func main() {
	store := NewUserStore()

	for {
		printMenu()

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			store.DisplayAll()
		case 2:
			store.DisplayFiltered(true)
		case 3:
			store.DisplayFiltered(false)
		case 4:
			var id int
			fmt.Print("Enter User ID to search: ")
			fmt.Scan(&id)
			if user, found := store.FindByID(id); found {
				fmt.Println("\nUser Found:")
				user.Display()
			} else {
				fmt.Printf("\nUser with ID %d not found.\n", id)
			}
		case 5:
			var name, email string
			var activeChoice int
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Email: ")
			fmt.Scan(&email)
			fmt.Print("Is Active? (1 for Yes, 0 for No): ")
			fmt.Scan(&activeChoice)
			store = store.AddUser(name, email, activeChoice == 1)
		case 6:
			var id int
			fmt.Print("Enter User ID to delete: ")
			fmt.Scan(&id)
			var success bool
			store, success = store.DeleteUser(id)
			if success {
				fmt.Printf("\nUser with ID %d successfully deleted.\n", id)
			} else {
				fmt.Printf("\nUser with ID %d not found.\n", id)
			}
		case 7:
			fmt.Println("\nThank you for using User Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please select an option between 1 and 7.")
		}
	}
}
*/