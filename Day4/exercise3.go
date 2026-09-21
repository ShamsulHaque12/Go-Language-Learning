package main

/*
import (
	"fmt"
	"sync"
)

// ==========================================
// ১. Mutex ছাড়া (Race Condition এর উদাহরণ):
// ==========================================
// এখানে দুইটা goroutine একই সাথে counter variable এ write করার চেষ্টা করে।
// যার ফলে Data Race বা Race Condition ঘটে এবং counter এর মান ২০০০ এর কম হতে পারে।
// এটি পরীক্ষা করতে: go run -race exercise3.go

func main() {
	fmt.Println("===== Mutex ছাড়া (Race Condition) ===== ")

	counter := 0
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++ // Mutex ছাড়া Race Condition হবে
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++ // Mutex ছাড়া Race Condition হবে
		}
	}()

	wg.Wait()
	fmt.Println("Final Counter (without Mutex):", counter)
}
*/
/*
// ==========================================
// ২. Mutex ব্যবহার করে সমাধান (Safe & Expected):
// ==========================================
// sync.Mutex এর Lock() এবং Unlock() ব্যবহার করে counter++ কে
// Critical Section হিসেবে সুরক্ষিত রাখা হয়েছে।
// ফলে কোনো Data Race হয় না এবং ফাইনাল মান সবসময় ২০০০ আসে।

func main() {
	fmt.Println("===== Mutex + Race Condition ========== ")

	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	// ২টা goroutine এর জন্য WaitGroup এ কাউন্টার ২ সেট করা হলো
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()   // Critical section lock
			counter++   // Safe increment
			mu.Unlock() // Lock release
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()   // Critical section lock
			counter++   // Safe increment
			mu.Unlock() // Lock release
		}
	}()

	// দুটো goroutine শেষ হওয়া পর্যন্ত অপেক্ষা করবে
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

*/