package main

/*
import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Message %d", i)
		fmt.Println("Sender -> Sending:", msg)
		ch <- msg // চ্যানেলে ডেটা পাঠানো হচ্ছে (Send)
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("Sender -> Closing the channel...")
	close(ch) // সব ডেটা পাঠানো শেষ হলে চ্যানেল ক্লোজ করতে হয়
}

func receiver(ch chan string) {
	// চ্যানেল ক্লোজ না হওয়া পর্যন্ত লুপটি একে একে ডেটা রিসিভ করতে থাকবে
	for msg := range ch {
		fmt.Println("Receiver <- Received:", msg)
	}

	fmt.Println("Receiver <- Channel closed, no more data!")
}

func main() {
	fmt.Println("======== Channel =======")

	// ১. make() দিয়ে একটি string চ্যানেল তৈরি করা হলো
	ch := make(chan string)

	// ২. sender-কে আলাদা একটি Goroutine হিসেবে চালু করা হলো
	// sender ব্যাকগ্রাউন্ডে মেসেজ পাঠাতে থাকবে
	go sender(ch)

	// ৩. receiver-কে main goroutine-এ কল করা হলো
	// চ্যানেল বন্ধ (close) না হওয়া পর্যন্ত receiver অপেক্ষা করবে ও মেসেজ রিসিভ করবে
	receiver(ch)

	fmt.Println("======== Program Finished ========")
}

*/