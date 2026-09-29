package main

/*
1. os.Open()
2. defer file.Close()
3. file.Read()
4. Content print


import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=============== os.Open() + file.Read() =============== ")

	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Total byte read:", count)
	fmt.Println("Content:\n", string(data[:count]))

}

*/
