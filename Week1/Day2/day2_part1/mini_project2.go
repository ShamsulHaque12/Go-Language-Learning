package main

/*
import (
	"errors"
	"fmt"
)

// 1. errors.New() - সেনটিনেল এরর (Sentinel Errors)
var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount       = errors.New("invalid amount")
)

// 2. struct - একাউন্ট ডাটা স্ট্রাকচার
type Account struct {
	ID      int
	Owner   string
	Balance float64
}

// 9. panic - ফাকা নাম বা নেগেটিভ ব্যালেন্স দিলে panic ট্রিগার হবে
func NewAccount(id int, owner string, initialBalance float64) Account {
	if owner == "" {
		panic("Account owner name cannot be empty!")
	}
	if initialBalance < 0 {
		panic("Initial balance cannot be negative!")
	}
	return Account{
		ID:      id,
		Owner:   owner,
		Balance: initialBalance,
	}
}

// 3. method & pointer receiver - একাউন্টে টাকা জমা করা
// 4. error, 5. fmt.Errorf(), 6. %w - এরর র‍্যাপ করা
func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("failed to deposit %.2ftk: %w", amount, ErrInvalidAmount)
	}
	acc.Balance += amount
	return nil
}

// 3. method & pointer receiver - একাউন্ট থেকে টাকা তোলা
func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("failed to withdraw %.2ftk: %w", amount, ErrInvalidAmount)
	}
	if amount > acc.Balance {
		return fmt.Errorf("failed to withdraw %.2ftk: %w", amount, ErrInsufficientBalance)
	}
	acc.Balance -= amount
	return nil
}

func main() {
	// 8. defer & 10. recover() - panic হলে প্রোগ্রাম ক্র্যাশ হওয়া আটকাবে
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n[System Recovered from Panic]:", r)
		}
	}()

	fmt.Println("============ Bank Account System ============")

	// ইউজারের থেকে তথ্য নেওয়া
	var id int
	var owner string
	var initialBalance float64

	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&id)

	fmt.Print("Enter Owner Name: ")
	fmt.Scanln(&owner)

	fmt.Print("Enter Initial Balance: ")
	fmt.Scanln(&initialBalance)

	// একাউন্ট তৈরি (ভুল ইনপুট দিলে panic হবে)
	acc := NewAccount(id, owner, initialBalance)
	fmt.Printf("\n--> Account Created: %+v\n\n", acc)

	// মেনু চালানোর জন্য লুপ
	for {
		fmt.Println("--------------------------------------------")
		fmt.Println("1. Deposit  | 2. Withdraw  | 3. Check Balance | 4. Exit")
		fmt.Print("Choose option (1-4): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scanln(&amount)

			err := acc.Deposit(amount)
			if err != nil {
				fmt.Println("Error:", err)
				// 7. errors.Is() - এরর চেক করা
				if errors.Is(err, ErrInvalidAmount) {
					fmt.Println("-> Check: Amount must be greater than 0.")
				}
			} else {
				fmt.Printf("Deposit Successful! New Balance: %.2ftk\n", acc.Balance)
			}

		case 2:
			var amount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scanln(&amount)

			err := acc.Withdraw(amount)
			if err != nil {
				fmt.Println("Error:", err)
				// 7. errors.Is() - এরর চেক করা
				if errors.Is(err, ErrInsufficientBalance) {
					fmt.Println("-> Check: Insufficient balance in account.")
				} else if errors.Is(err, ErrInvalidAmount) {
					fmt.Println("-> Check: Invalid withdrawal amount.")
				}
			} else {
				fmt.Printf("Withdrawal Successful! Remaining Balance: %.2ftk\n", acc.Balance)
			}

		case 3:
			fmt.Printf("Owner: %s | Current Balance: %.2ftk\n", acc.Owner, acc.Balance)

		case 4:
			fmt.Println("Exiting Bank System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice! Please select 1 to 4.")
		}
	}
}

*/
