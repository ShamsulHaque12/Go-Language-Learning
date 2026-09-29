package main

/*
import (
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	Name   string
	Price  float64
	Stock  int
	Active bool
}

func main() {
	fmt.Println("====== Challenge 2: Product Processing =======")

	productName := "   iPhone 17 Pro   "
	priceText := " 1299.99 "
	stockText := " 25 "
	activeText := " true "

	fmt.Println("\n--- Raw Inputs ---")
	fmt.Printf("Raw Name   : %q\n", productName)
	fmt.Printf("Raw Price  : %q\n", priceText)
	fmt.Printf("Raw Stock  : %q\n", stockText)
	fmt.Printf("Raw Active : %q\n", activeText)

	// 1. Clean Name -> TrimSpace
	cleanName := strings.TrimSpace(productName)
	if cleanName == "" {
		fmt.Println("\nError: Product name cannot be empty!")
		return
	}

	// 2. Clean & Convert Price -> ParseFloat
	price, err := strconv.ParseFloat(strings.TrimSpace(priceText), 64)
	if err != nil {
		fmt.Println("\nError: Invalid price input!", err)
		return
	}
	if price <= 0 {
		fmt.Println("\nError: Price must be greater than 0!")
		return
	}

	// 3. Clean & Convert Stock -> Atoi
	stock, err := strconv.Atoi(strings.TrimSpace(stockText))
	if err != nil {
		fmt.Println("\nError: Invalid stock input!", err)
		return
	}
	// Check Stock > 0
	if stock <= 0 {
		fmt.Println("\nError: Stock must be greater than 0!")
		return
	}

	// 4. Clean & Convert Active -> ParseBool
	active, err := strconv.ParseBool(strings.TrimSpace(activeText))
	if err != nil {
		fmt.Println("\nError: Invalid active state!", err)
		return
	}

	// Create Product struct
	prod := Product{
		Name:   cleanName,
		Price:  price,
		Stock:  stock,
		Active: active,
	}

	// Print product information nicely
	fmt.Println("\n===== Validated Product Information =====")
	fmt.Printf("Product Name : %s\n", prod.Name)
	fmt.Printf("Price        : $%.2f\n", prod.Price)
	fmt.Printf("Stock        : %d units\n", prod.Stock)
	fmt.Printf("Active Status: %t\n", prod.Active)
	fmt.Printf("Struct       : %+v\n", prod)
}

*/