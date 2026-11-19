package main

/*
import (
	"errors"
	"fmt"
)

var ErrInvalidAmount = errors.New("invalid withdrawal amount")
var ErrInsufficientBalance = errors.New("insufficient balance")

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientBalance
	}

	a.Balance -= amount
	return nil
}

func testWithdraw(acc *Account, amount float64) {
	fmt.Printf("Withdraw: %.0f\n", amount)
	err := acc.withdraw(amount)
	if err != nil {
		fmt.Printf("→ Error: %v\n", err)
	} else {
		fmt.Printf("→ New Balance: %.0f\n", acc.Balance)
	}
}

func main() {
	acc := Account{Owner: "User", Balance: 5000}
	fmt.Printf("Balance: %.0f\n\n", acc.Balance)

	testWithdraw(&acc, 2000)
	fmt.Println()
	testWithdraw(&acc, 5000)
	fmt.Println()
	testWithdraw(&acc, -100)
}

*/
