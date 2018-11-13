package main
import (
	"fmt"
)

func subtract(base int, nums ...int) int {
	
	for _, num := range nums {
		base -= num
	}

	return base
}

func main() {
	fmt.Println(subtract(4, 1, 2));
}

