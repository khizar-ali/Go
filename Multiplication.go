package main
import (
	"fmt"
)

func multiply(nums ...int) int {
	
	total := 1

	for _, num := range nums {
		total *= num
	}

	return total
}

func main() {
	fmt.Println(multiply(4, 1, 2));
}

