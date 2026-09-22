package main

/*
==== আজ প্রথমে শিখবো ====
1. os.Create()
2. os.WriteFile()
3. os.ReadFile()
4. os.Open()
5. os.Close()
6. File error handling



import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("====== File Handling ======")

	// 1. Create "users.txt" file
	file, err := os.Create("users.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("users.txt created successfully!")

	// 2. Write details to the file
	content := "Name: Sujon\nAge: 27\nRole: Flutter Developer"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		file.Close()
		return
	}
	fmt.Println("Data written to users.txt successfully!")

	// 3. Close the file
	file.Close()

	// 4. Read file content
	data, err := os.ReadFile("users.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 5. Print entire content to terminal
	fmt.Println("\n--- File Content ---")
	fmt.Println(string(data))
}

*/