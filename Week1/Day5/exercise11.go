package main

/*
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	origName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	fmt.Print("Enter Email: ")
	origEmail, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading email:", err)
		return
	}

	fmt.Print("Enter Role: ")
	origRole, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading role:", err)
		return
	}

	// Remove newline characters from input for clean representation
	origNameCleanNL := strings.TrimRight(origName, "\r\n")
	origEmailCleanNL := strings.TrimRight(origEmail, "\r\n")
	origRoleCleanNL := strings.TrimRight(origRole, "\r\n")

	cleanName := strings.TrimSpace(origNameCleanNL)
	cleanEmail := strings.ToLower(strings.TrimSpace(origEmailCleanNL))
	cleanRole := strings.TrimSpace(origRoleCleanNL)

	isGmail := strings.HasSuffix(cleanEmail, "@gmail.com")

	fmt.Println("\n===== User Input Cleaner =====")
	fmt.Println()
	fmt.Printf("Original Name:   \"%s\"\n", origNameCleanNL)
	fmt.Printf("Clean Name:      \"%s\"\n", cleanName)
	fmt.Println()
	fmt.Printf("Original Email:  \"%s\"\n", origEmailCleanNL)
	fmt.Printf("Clean Email:     \"%s\"\n", cleanEmail)
	fmt.Println()
	fmt.Printf("Original Role:   \"%s\"\n", origRoleCleanNL)
	fmt.Printf("Clean Role:      \"%s\"\n", cleanRole)
	fmt.Println()
	fmt.Printf("Is Gmail: %t\n", isGmail)
}

*/