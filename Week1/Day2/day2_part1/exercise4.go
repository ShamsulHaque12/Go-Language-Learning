package main

/*
import (
	"fmt"
)

// Account struct: ব্যাংকের অ্যাকাউন্টের ডাটা স্ট্রাকচার (OwnerName এবং Balance ধারণ করার জন্য)
type Account struct {
	OwnerName string  // অ্যাকাউন্ট মালিকের নাম
	Balance   float64 // বর্তমান ব্যালেন্স
}

// deposit: অ্যাকাউন্টে টাকা জমা করার মেথড (টাকা সফলভাবে জমা হলে true, অন্যথায় false রিটার্ন করে)
func (a *Account) deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}

	a.Balance += amount
	return true
}

// withdraw: অ্যাকাউন্ট থেকে টাকা তোলার মেথড (টাকা সফলভাবে তোলা গেলে true, অন্যথায় false রিটার্ন করে)
func (a *Account) withdraw(amount float64) bool {
	if amount <= 0 {
		return false
	}

	if a.Balance < amount {
		return false
	}

	a.Balance -= amount
	return true
}

// getBalance: অ্যাকাউন্টের বর্তমান ব্যালেন্স রিটার্ন করার মেথড
func (a *Account) getBalance() float64 {
	return a.Balance
}

// displayAccount: অ্যাকাউন্টের বিস্তারিত তথ্য প্রিন্ট করার মেথড
func (a *Account) displayAccount() {
	fmt.Println("\n--- Account Details ---")
	fmt.Println("Owner Name: ", a.OwnerName)
	fmt.Println("Balance:    ", a.Balance)
	fmt.Println("-----------------------")
}

// menu: ব্যবহারকারীর সুবিধার জন্য মেনু নির্দেশিকা দেখানোর ফাংশন
func menu() {
	fmt.Println("\n1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Get Balance")
	fmt.Println("4. Display Account")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

func main() {
	fmt.Println("======== Exercise 4 ==========")

	var account Account // একটি Account অবজেক্ট/ভেরিয়েবল তৈরি করা হলো
	var choice int      // ইউজার কোনটা বেছে নিবে তার জন্য ভেরিয়েবল

	// অ্যাকাউন্ট মালিকের নাম ইনপুট নেওয়া
	fmt.Print("Enter Account Owner Name: ")
	fmt.Scan(&account.OwnerName)

	// অসীম লুপ (Infinite Loop), ব্যবহারকারী ৫ (Exit) চেপে বের না হওয়া পর্যন্ত লুপ চলবে
	for {
		menu()            // অপশন প্রদর্শন করবে
		fmt.Scan(&choice) // ইউজারের চয়েস ইনপুট নিবে

		switch choice {
		case 1:
			// টাকা জমা দেয়ার লজিক
			fmt.Print("Enter Amount to Deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if account.deposit(amount) {
				fmt.Println("Deposit successful")
			} else {
				fmt.Println("Invalid deposit amount")
			}
		case 2:
			// টাকা তোলার লজিক
			fmt.Print("Enter Amount to Withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if account.withdraw(amount) {
				fmt.Println("Withdrawal successful")
			} else {
				fmt.Println("Withdrawal failed")
			}
		case 3:
			// ব্যালেন্স দেখার লজিক
			fmt.Printf("Balance: %.2f\n", account.getBalance())
		case 4:
			// অ্যাকাউন্টের বিবরণ দেখার লজিক
			account.displayAccount()
		case 5:
			// প্রোগ্রাম থেকে বের হয়ে যাওয়া
			fmt.Println("Exiting...")
			return
		default:
			// ভুল ইনপুট দিলে মেসেজ প্রদর্শন
			fmt.Println("Invalid Choice! Please try again.")
		}
	}
}

*/