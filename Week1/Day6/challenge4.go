package main

/*
import (
	"fmt"
	"strconv"
	"strings"
)

// 1. Struct for Product
type Product struct {
	Name  string
	Price float64
	Stock int
}

// 2. Struct for Order
type Order struct {
	ProductName string
	Quantity    int
	Total       float64
}

// Function to print all products
func printProducts(products []Product) {
	fmt.Println("\n--- Available Products ---")
	for i, p := range products {
		fmt.Printf("%d. %s - Price: $%.2f | Stock: %d\n", i+1, p.Name, p.Price, p.Stock)
	}
}

// Function to process order with error handling
func processOrder(products []Product, name string, qty int) (Order, error) {
	// Validation: Quantity check
	if qty <= 0 {
		return Order{}, fmt.Errorf("quantity must be greater than 0")
	}

	// Search product in slice using loop
	cleanName := strings.TrimSpace(strings.ToLower(name))
	for i := range products {
		if strings.ToLower(products[i].Name) == cleanName {
			// Validation: Stock check
			if qty > products[i].Stock {
				return Order{}, fmt.Errorf("insufficient stock! (Available: %d)", products[i].Stock)
			}

			// Reduce stock
			products[i].Stock -= qty

			// Calculate total bill
			total := products[i].Price * float64(qty)

			// Return new Order struct
			return Order{
				ProductName: products[i].Name,
				Quantity:    qty,
				Total:       total,
			}, nil
		}
	}

	// If product not found
	return Order{}, fmt.Errorf("product %q not found", name)
}

func main() {
	fmt.Println("=========== Mini Order System =========")

	// Slice of Products
	products := []Product{
		{Name: "Laptop", Price: 1200.00, Stock: 5},
		{Name: "Mouse", Price: 25.50, Stock: 20},
		{Name: "Keyboard", Price: 45.00, Stock: 10},
	}

	// Slice of Orders
	var orders []Order

	// Show Product List
	printProducts(products)

	// User Input
	var nameInput, qtyInput string

	fmt.Print("\nEnter Product Name: ")
	fmt.Scanln(&nameInput)

	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&qtyInput)

	// Convert string -> int using strconv.Atoi
	qty, err := strconv.Atoi(strings.TrimSpace(qtyInput))
	if err != nil {
		fmt.Println("❌ Error: Quantity must be a valid number!")
		return
	}

	// Place order using function
	order, err := processOrder(products, nameInput, qty)
	if err != nil {
		fmt.Println("❌ Order Failed:", err)
	} else {
		orders = append(orders, order)
		fmt.Println("\n✅ Order Placed Successfully!")
		fmt.Printf("Item: %s | Qty: %d | Total: $%.2f\n", order.ProductName, order.Quantity, order.Total)
	}

	// Show Updated Stock
	printProducts(products)
}

*/