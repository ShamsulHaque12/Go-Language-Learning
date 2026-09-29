package main

/*
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("======== Standard Library topic: strconv ==========")

	// 1. strconv.Atoi() - string -> int
	numStr := "123"
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Printf("1. Atoi       (string -> int)    : %q -> %d (Type: %T)\n", numStr, numInt, numInt)

	// 2. strconv.Itoa() - int -> string
	valInt := 456
	valStr := strconv.Itoa(valInt)
	fmt.Printf("2. Itoa       (int -> string)    : %d -> %q (Type: %T)\n", valInt, valStr, valStr)

	// 3. strconv.ParseFloat() - string -> float64
	floatStr := "99.99"
	valFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}
	fmt.Printf("3. ParseFloat (string -> float64): %q -> %.2f (Type: %T)\n", floatStr, valFloat, valFloat)

	// 4. strconv.ParseBool() - string -> bool
	boolStr := "true"
	valBool, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
		return
	}
	fmt.Printf("4. ParseBool  (string -> bool)   : %q -> %t (Type: %T)\n", boolStr, valBool, valBool)
}

*/