package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Stock    int
	Quantity int
}

func (p Product) showProduct() {
	fmt.Println()
	fmt.Printf("Product ID: %d\n", p.ID)
	fmt.Printf("Product Name: %s\n", p.Name)
	fmt.Printf("Product Price: %.2f\n", p.Price)
	fmt.Printf("Product Stock: %d\n", p.Stock)
	fmt.Printf("Product Quantity: %d\n", p.Quantity)
}

func (p Product) isAvailable() bool {
	return p.Stock > 0
}

func (p Product) calculateTotalPrice() float64 {
	if !p.isAvailable() {
		return 0.0
	}
	return p.Price * float64(p.Quantity)
}

func (p Product) remainingStock() int {
	return p.Stock - p.Quantity
}

func (p Product) discount() float64 {
	if p.Price >= 100000 {
		return 0.10
	} else if p.Price >= 50000 {
		return 0.07
	} else {
		return 0.02
	}
}

func main() {
	fmt.Println("===================Exercise 4===================")

	product1 := Product{
		ID:       1,
		Name:     "Laptop",
		Price:    50000,
		Stock:    10,
		Quantity: 5,
	}

	product2 := Product{
		ID:       2,
		Name:     "Mouse",
		Price:    500,
		Stock:    100,
		Quantity: 20,
	}

	product3 := Product{
		ID:       3,
		Name:     "Keyboard",
		Price:    1000,
		Stock:    0,
		Quantity: 0,
	}

	products4 := Product{
		ID:       4,
		Name:     "Mobile",
		Price:    1000000,
		Stock:    50,
		Quantity: 10,
	}

	products5 := Product{
		ID:       5,
		Name:     "Mobile",
		Price:    10000,
		Stock:    5,
		Quantity: 0,
	}

	products := []Product{
		product1,
		product2,
		product3,
		products4,
		products5,
	}

	grandTotalWithoutDiscount := 0.0
	grandTotalWithDiscount := 0.0

	for _, product := range products {
		product.showProduct()

		if product.isAvailable() {
			total := product.calculateTotalPrice()
			discountRate := product.discount()
			discountAmount := total * discountRate
			finalPrice := total - discountAmount

			fmt.Println("Status: Available")
			fmt.Printf("Remaining Stock: %d\n", product.remainingStock())
			fmt.Printf("Total Price: %.2f\n", total)
			fmt.Printf("Discount: %.2f\n", discountRate)
			fmt.Printf("Total Discount: %.2f\n", discountAmount)
			fmt.Printf("Total Price After Discount: %.2f\n", finalPrice)

			grandTotalWithoutDiscount += total
			grandTotalWithDiscount += finalPrice
		} else {
			fmt.Println("Status: Out of Stock")
		}
	}

	fmt.Println()
	fmt.Println("======================================")
	fmt.Printf("Grand Total (Without Discount): %.2f\n", grandTotalWithoutDiscount)
	fmt.Printf("Grand Total (With Discount): %.2f\n", grandTotalWithDiscount)
	fmt.Println("======================================")
}
