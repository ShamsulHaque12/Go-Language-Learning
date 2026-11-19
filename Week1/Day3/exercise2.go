package main

/*
import "fmt"

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Animal interface {
	Speak()
	Move()
}

func (d Dog) Speak() {
	fmt.Printf("%s says: Woof! Woof!\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s is running happily.\n", d.Name)
}

func (c Cat) Speak() {
	fmt.Printf("%s says: Meow! Meow!\n", c.Name)
}

func (c Cat) Move() {
	fmt.Printf("%s is walking softly.\n", c.Name)
}

// Function using single Type Assertion with comma-ok idiom
func checkTypeWithAssertion(a Animal) {
	dog, ok := a.(Dog)
	if ok {
		fmt.Printf("[Type Assertion] Found a Dog with name: %s\n", dog.Name)
		return
	}

	cat, ok := a.(Cat)
	if ok {
		fmt.Printf("[Type Assertion] Found a Cat with name: %s\n", cat.Name)
		return
	}

	fmt.Println("[Type Assertion] Unknown animal type")
}

// Function using Type Switch
func checkTypeWithSwitch(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("[Type Switch] It's a Dog! Name: %s\n", v.Name)
	case Cat:
		fmt.Printf("[Type Switch] It's a Cat! Name: %s\n", v.Name)
	default:
		fmt.Println("[Type Switch] Unknown type")
	}
}

func main() {
	fmt.Println("==== Type Assertion ====")

	var a1 Animal = Dog{Name: "Buddy"}
	var a2 Animal = Cat{Name: "Kitty"}

	fmt.Println("\n--- Testing Animal 1 (Dog) ---")
	checkTypeWithAssertion(a1)
	checkTypeWithSwitch(a1)

	fmt.Println("\n--- Testing Animal 2 (Cat) ---")
	checkTypeWithAssertion(a2)
	checkTypeWithSwitch(a2)
}

*/