package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Factorial for 0 is : %d \n", Factorial(uint64(0)))
	fmt.Printf("Factorial for 1 is : %d \n", Factorial(uint64(1)))
	fmt.Printf("Factorial for 2 is : %d \n", Factorial(uint64(2)))
	fmt.Printf("Factorial for 6 is : %d \n", Factorial(uint64(6)))
	fmt.Printf("Factorial for 32 is : %d \n", Factorial(uint64(32)))
}

func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
