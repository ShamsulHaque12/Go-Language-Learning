package main

import (
	"errors"
	"fmt"
)

var errorMessage = errors.New("Something went wrong!")

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) withdraw(amount float64) error {
	if a.Balance < amount {
		return errorMessage
	}

	a.Balance -= amount
	return nil
}

func main() {
	fmt.Println("======= Exercise 6 =======")

}
