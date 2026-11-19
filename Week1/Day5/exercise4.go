package main

/*
import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

func main() {
	fmt.Println("===== Struct -> JSON -> File ===== ")

	user := User{
		Id:    101,
		Name:  "Sujon",
		Email: "sujon@gmail.com",
		Age:   27,
		Role:  "Flutter Developer",
	}

	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	file, err := os.Create("user.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to user.json")

	// Read JSON from file
	readData, err := os.ReadFile("user.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("\n===== User JSON File =====")
	fmt.Println("\nJSON File Content:")
	fmt.Println(string(readData))

	// Convert (Unmarshal) JSON back to Struct
	var readUser User
	err = json.Unmarshal(readData, &readUser)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("\n===== Converted Back To Struct =====")
	fmt.Printf("\nID: %d\nName: %s\nEmail: %s\nAge: %d\nRole: %s\n", readUser.Id, readUser.Name, readUser.Email, readUser.Age, readUser.Role)
}

*/