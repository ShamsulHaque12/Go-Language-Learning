package main

/*
import "fmt"

type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

func (u User) showUser() {
	fmt.Println()
	fmt.Printf("User ID: %d\n", u.ID)
	fmt.Printf("User Name: %s\n", u.Name)
	fmt.Printf("User Email: %s\n", u.Email)
	fmt.Printf("User IsActive: %t\n", u.IsActive)
}

func (u User) isActive() bool {
	return u.IsActive
}

func findUser(user []User, id int) (User, bool) {
	for _, u := range user {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

func main() {
	fmt.Println("===================Exercise 5===================")

	user1 := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "[EMAIL_ADDRESS]",
		IsActive: true,
	}

	user2 := User{
		ID:       2,
		Name:     "Jane Doe",
		Email:    "[EMAIL_ADDRESS]",
		IsActive: false,
	}

	user3 := User{
		ID:       3,
		Name:     "Bob Smith",
		Email:    "[EMAIL_ADDRESS]",
		IsActive: true,
	}

	user4 := User{
		ID:       4,
		Name:     "Charlie Brown",
		Email:    "[EMAIL_ADDRESS]",
		IsActive: false,
	}

	user5 := User{
		ID:       5,
		Name:     "Emily Davis",
		Email:    "[EMAIL_ADDRESS]",
		IsActive: false,
	}

	users := []User{user1, user2, user3, user4, user5}

	activeUsers := 0
	for _, user := range users {
		user.showUser()
		if user.isActive() {
			fmt.Println("Status: Active")
			activeUsers++
		} else {
			fmt.Println("Status: Inactive")
		}
	}

	fmt.Println()
	fmt.Println("======================================")
	fmt.Printf("Active Users: %d\n", activeUsers)
	fmt.Println("======================================")

	fmt.Println()
	fmt.Println("================ Search User ================")
	var searchID int
	fmt.Print("Enter User ID to search: ")
	fmt.Scan(&searchID)

	if user, found := findUser(users, searchID); found {
		fmt.Printf("\nUser found with ID %d:\n", searchID)
		user.showUser()
	} else {
		fmt.Printf("\nUser with ID %d not found.\n", searchID)
	}
}
*/