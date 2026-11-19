package main

/*
import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	fmt.Println("==== Struct Tags ====")

	// 1. Creating a User instance
	user := User{
		ID:    1,
		Name:  "Sujon",
		Email: "sujon@example.com",
	}

	// 2. Struct to JSON (Serialization / Marshaling)
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("\n--- Struct to JSON ---")
	fmt.Println(string(jsonData))

	// 3. JSON to Struct (Deserialization / Unmarshaling)
	rawJSON := `{"id": 2, "name": "Rahim", "email": "rahim@example.com"}`
	var newUser User
	err = json.Unmarshal([]byte(rawJSON), &newUser)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("\n--- JSON to Struct ---")
	fmt.Printf("User: %+v\n", newUser)

	// 4. Reading Struct Tags using Reflection
	fmt.Println("\n--- Reading Struct Tags via Reflection ---")
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("Field: %-6s | Tag (json): %s\n", field.Name, tag)
	}
}

*/
