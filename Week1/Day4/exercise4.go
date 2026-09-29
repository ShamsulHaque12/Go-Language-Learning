package main

/*
import (
	"fmt"
	"sync"
	"time"
)

// Worker ফাংশন: চ্যানেল থেকে job নিয়ে প্রসেস করবে
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started Job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d finished Job %d\n", id, job)
	}
}

func main() {
	fmt.Println("====== Goroutine + Channel + Buffered Channel + WaitGroup ======")

	// ৫ সাইজের একটি Buffered Channel (jobs)
	jobs := make(chan int, 5)

	var wg sync.WaitGroup

	// ৩টা worker goroutine চালু করা হলো
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// ৫টা job চ্যানেলে পাঠানো হলো
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	// সব job পাঠানো শেষ হলে চ্যানেলটি close করতে হবে
	// যাতে worker-রা বুঝতে পারে আর কোনো নতুন job আসবে না
	close(jobs)

	// ৩টা worker-এর কাজ শেষ হওয়া পর্যন্ত অপেক্ষা করবে
	wg.Wait()

	// সব কাজ শেষ হলে প্রিন্ট হবে
	fmt.Println("\nAll jobs completed!")
}

*/
