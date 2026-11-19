package main
/*

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("===== Environment Variables ======")



	// Set an environment variable
	err := os.Setenv("APP_ENV", "production")
	if err != nil {
		fmt.Println("Error setting env var:", err)
		return
	}

	// Get an environment variable
	val := os.Getenv("APP_ENV")
	fmt.Println("APP_ENV:", val)

	// Get all environment variables
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Remove an environment variable
	os.Unsetenv("APP_ENV")
	val = os.Getenv("APP_ENV")
	fmt.Println("APP_ENV:", val)
}

*/