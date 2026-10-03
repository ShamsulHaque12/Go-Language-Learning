package main

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var message string

	switch r.Method {
	case http.MethodGet:
		message = "Get all users"
	case http.MethodPost:
		message = "Create user"
	case http.MethodPut:
		message = "Update user"
	case http.MethodDelete:
		message = "Delete user"
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		message = "Method not allowed"
	}

	resp := Response{
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Println("Failed to encode response:", err)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server starting on port :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

*/