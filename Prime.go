package main;
import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if(num <= 1) {
		return false
	}

	if(num == 2 || num == 3) {
		return true
	} else {
		
		tmp := int(math.Floor(float64(num) / 2))

		for i := 2; i <= tmp; i++ {
			if num%i == 0 {
				return false
			}
		}
	}

	return true
}

func main() {

	fmt.Println("0 is prime = ", isPrime(0));
	fmt.Println("1 is prime = ", isPrime(1));
	fmt.Println("2 is prime = ", isPrime(2));
	fmt.Println("121 is prime = ", isPrime(121));
	fmt.Println("131 is prime = ", isPrime(131));
	fmt.Println("97 is prime = ", isPrime(97));
	fmt.Println("111 is prime = ", isPrime(111));
	fmt.Println("17 is prime = ", isPrime(17));
	fmt.Println("100000000 is prime = ", isPrime(100000000));
	fmt.Println("384.23 is prime = ", isPrime(384));	
}
