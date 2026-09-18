package main

/*
import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello from goroutine", name)
}

func main() {
	go sayHello("Sujon")
	go sayHello("Haque")

	time.Sleep(time.Second)

	fmt.Println("Main function")
	fmt.Println("Waiting for goroutines to finish")

}
*/

/*
import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(i)
	}
}

func printCharacters() {
	for char := 'A'; char <= 'E'; char++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("%c\n", char)
	}
}

func main() {
	fmt.Println("===== GoRoutine =====")

	go printNumbers()
	go printCharacters()

	time.Sleep(2 * time.Second)

	fmt.Println("Main function finished!")
}
*/
/*
func main() {
	fmt.Println("===== GoRoutine ==========")

	fmt.Println("--- Numbers ---")
	go printNumbers()
	time.Sleep(time.Millisecond * 600)

	fmt.Println("\n--- Characters ---")
	go printCharacters()
	time.Sleep(time.Millisecond * 600)

	fmt.Println("\nMain function: All goroutines finished!")
}
*/
