package main

/*
import (
	"fmt"
	"sync"
	"time"
)

// ১. Sender ফাংশন: চ্যানেলে ডেটা পাঠাবে
func sender(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Message %d", i)

		fmt.Printf("Sender -> পাঠাতে চাচ্ছে: %s | বর্তমান বাফার: %d/%d\n", msg, len(ch), cap(ch))

		ch <- msg

		fmt.Printf("Sender -> পাঠানো সফল: %s | বাফারে জমা হলো: %d/%d\n\n", msg, len(ch), cap(ch))
		time.Sleep(100 * time.Millisecond)
	}

	close(ch)
}

// ২. Receiver ফাংশন: চ্যানেল থেকে ডেটা গ্রহণ করবে
func receiver(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 👀 লক্ষ্য করুন: Receiver শুরুতে ২ সেকেন্ড অপেক্ষা (Sleep) করবে
	// এর ফলে স্পষ্ট দেখা যাবে যে কোনো Receiver ছাড়াই প্রথম ৩টি মেসেজ বাফারে ঢুকে গেছে!
	time.Sleep(2 * time.Second)

	fmt.Println("Receiver -> এবার ঘুম থেকে উঠে ডেটা নেওয়া শুরু করছে:")
	for msg := range ch {
		fmt.Printf("Receiver <- রিসিভ করেছে: %s | বাফারে অবশিষ্ট: %d/%d\n", msg, len(ch), cap(ch))
		time.Sleep(500 * time.Millisecond) // প্রতিটা ডেটা প্রসেস করতে সময় নিচ্ছে
	}

	fmt.Println("\nReceiver <- চ্যানেল ক্লোজ এবং বাফারের সব ডেটা প্রসেস শেষ!")
}

func main() {
	fmt.Println("=========== Buffered Channel (Capacity = 3) ===========")

	// ৩ ধারণক্ষমতার (Capacity = 3) একটি Buffered Channel
	ch := make(chan string, 3)

	var wg sync.WaitGroup

	wg.Add(1)
	go receiver(ch, &wg) // Receiver গোরুটিন চালু হলো (তবে সে ২ সেকেন্ড ঘুমাবে)
	go sender(ch)        // Sender সাথে সাথেই পাঠানো শুরু করবে

	wg.Wait() // Receiver শেষ না হওয়া পর্যন্ত অপেক্ষা

	fmt.Println("=========== Program Finished ===========")
}

*/
