package main

/*
import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("===== Application Config System =====")

	// Load .env file into environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file:", err)
	}

	// 1. Fetching individual env variables using os.Getenv
	appName := os.Getenv("APP_NAME")
	appEnv := os.Getenv("APP_ENV")
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")

	if port == "" {
		port = "8080" // default fallback
	}

	fmt.Println("\n--- Individual Configs ---")
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("App Env:  %s\n", appEnv)
	fmt.Printf("Port:     %s\n", port)
	fmt.Printf("DB Host:  %s\n", dbHost)

	// 2. Reading all key-value pairs directly from .env file using godotenv.Read
	fmt.Println("\n--- All .env File Contents ---")
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Println("Error reading .env file:", err)
		return
	}

	for key, value := range envMap {
		fmt.Printf("%s = %s\n", key, value)
	}
}

*/
