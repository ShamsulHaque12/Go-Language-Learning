package main

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Method  string `json:"method"`
	Path    string `json:"path"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request on %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Message: "User list will come here",
		Method:  r.Method,
		Path:    r.URL.Path,
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