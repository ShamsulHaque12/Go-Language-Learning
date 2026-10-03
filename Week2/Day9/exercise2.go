package main

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Response struct {
	Message string `json:"message"`
	User    *User  `json:"user,omitempty"`
	Users   []User `json:"users,omitempty"`
}

// In-memory user database
var users = []User{
	{ID: 1, Name: "Sujon", Email: "sujon@gmail.com", Age: 27},
	{ID: 2, Name: "Rahim", Email: "rahim@gmail.com", Age: 25},
	{ID: 3, Name: "Karim", Email: "karim@gmail.com", Age: 30},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Clean up path trailing slash for easy comparison
	path := strings.TrimSuffix(r.URL.Path, "/")

	switch r.Method {
	case http.MethodGet:
		// Case 1: GET /users -> Return all users
		if path == "/users" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response{
				Message: "All users fetched successfully",
				Users:   users,
			})
			return
		}

		// Case 2: GET /users/{id} -> Return single user by ID
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: "Invalid user ID",
			})
			return
		}

		// Search for user with the given ID
		var foundUser *User
		for i := range users {
			if users[i].ID == id {
				foundUser = &users[i]
				break
			}
		}

		// If user not found, return 404 Not Found
		if foundUser == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Message: "User not found",
			})
			return
		}

		// If user found, return 200 OK with user object
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: "User fetched successfully",
			User:    foundUser,
		})

	case http.MethodPut:
		// Case 3: PUT /users/{id} -> Update user by ID
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: "Invalid user ID",
			})
			return
		}

		var updatedUser User
		err = json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: "Invalid JSON format",
			})
			return
		}

		if updatedUser.Name == "" || updatedUser.Email == "" || updatedUser.Age <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: "Invalid user data",
			})
			return
		}

		// Find existing user index
		foundIndex := -1
		for i := range users {
			if users[i].ID == id {
				foundIndex = i
				break
			}
		}

		// If user not found, return 404 Not Found
		if foundIndex == -1 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Message: "User not found",
			})
			return
		}

		// Update existing user data
		updatedUser.ID = id
		users[foundIndex] = updatedUser

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: "User updated successfully",
			User:    &users[foundIndex],
		})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{
			Message: "Method not allowed",
		})
	}
}

func main() {
	// Note: "/users/" with trailing slash matches both "/users" and "/users/{id}"
	http.HandleFunc("/users/", usersHandler)

	fmt.Println("Server starting on port :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

*/