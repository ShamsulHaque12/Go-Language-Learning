package main

/*
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("===== Wait Group with Goroutine =========")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			fmt.Println(i)
		}

	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			fmt.Println(i)
		}

		fmt.Println("Goroutine 2 finished!")
	}()

	wg.Wait()
	fmt.Println("Main function finished!")

}

*/