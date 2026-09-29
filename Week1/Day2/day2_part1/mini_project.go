package main

/*
import (
	"errors"
	"fmt"
)

// 1. errors.New() - Define sentinel errors
var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount       = errors.New("invalid amount")
)

// 2. struct - Define Account structure
type Account struct {
	ID      int
	Owner   string
	Balance float64
}

// Function to construct Account (uses panic for invalid arguments)
func NewAccount(id int, owner string, initialBalance float64) Account {
	if owner == "" {
		// 9. panic - Trigger panic for unrecoverable/invalid state
		panic("Account owner name cannot be empty")
	}
	if initialBalance < 0 {
		panic("Initial balance cannot be negative")
	}
	return Account{
		ID:      id,
		Owner:   owner,
		Balance: initialBalance,
	}
}

// 3. method & pointer receiver - Deposit modifies the Account state
// 4. error & 5. fmt.Errorf() with 6. %w - Wrapping sentinel error
func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("failed to deposit %.2ftk into account %d: %w", amount, acc.ID, ErrInvalidAmount)
	}
	acc.Balance += amount
	return nil
}

// 3. method & pointer receiver - Withdraw modifies the Account state
func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("failed to withdraw %.2ftk from account %d: %w", amount, acc.ID, ErrInvalidAmount)
	}
	if amount > acc.Balance {
		return fmt.Errorf("failed to withdraw %.2ftk from account %d: %w", amount, acc.ID, ErrInsufficientBalance)
	}
	acc.Balance -= amount
	return nil
}

func main() {
	// 8. defer & 10. recover() - Catch panic safely when main exits
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n[Panic Recovered]:", r)
		}
	}()

	fmt.Println("============ Mini Project: Bank Account System ============")

	// Create a valid account
	acc := NewAccount(1, "Shamsul Haque", 1000)
	fmt.Printf("Created Account: %+v\n\n", acc)

	// Valid Deposit
	if err := acc.Deposit(500); err != nil {
		fmt.Println("Deposit Error:", err)
	} else {
		fmt.Printf("Deposit 500tk Successful | New Balance: %.2ftk\n", acc.Balance)
	}

	// Valid Withdrawal
	if err := acc.Withdraw(200); err != nil {
		fmt.Println("Withdraw Error:", err)
	} else {
		fmt.Printf("Withdraw 200tk Successful | New Balance: %.2ftk\n", acc.Balance)
	}

	fmt.Println("\n--- Testing Error Handling & errors.Is() ---")

	// Invalid Deposit (Negative amount)
	err := acc.Deposit(-50)
	if err != nil {
		fmt.Println("Error:", err)
		// 7. errors.Is() - Check if wrapped error matches sentinel error
		if errors.Is(err, ErrInvalidAmount) {
			fmt.Println("  -> Detected Sentinel Error: ErrInvalidAmount")
		}
	}

	// Invalid Withdrawal (Insufficient balance)
	err = acc.Withdraw(5000)
	if err != nil {
		fmt.Println("Error:", err)
		// 7. errors.Is() - Check if wrapped error matches sentinel error
		if errors.Is(err, ErrInsufficientBalance) {
			fmt.Println("  -> Detected Sentinel Error: ErrInsufficientBalance")
		}
	}

	fmt.Println("\n--- Testing Panic & Recover ---")
	// Valid creation for Sujon
	acc2 := NewAccount(2, "Sujon", 500)
	fmt.Printf("Successfully created Account 2: %+v\n\n", acc2)

	// Attempting invalid creation to trigger panic (empty owner name)
	fmt.Println("Attempting to create invalid account (empty owner name)...")
	_ = NewAccount(3, "", 500) // This triggers panic!

	// The following line will not execute because panic stops normal flow until defer recover catches it
	fmt.Println("This line will not execute because panic was triggered.")
}
*/

