package main

import (
	"fmt"
	"strings"
)

func isBalanced(value string) bool {

	value = strings.ToLower(strings.Replace(value, " ", "", -1))

	if value == "" {
		return false
	}

    for i := 0; i < len(value)/2; i++ {

        if value[i] != value[len(value)-i-1] {

			return false
			
        }
	}
	
    return true
}

func main() {
	fmt.Println("empty string: ", isPalindrome(""))
	fmt.Println("abcdcba: ",isPalindrome("abcdcba"))
	fmt.Println("abcd: ", isPalindrome("abcd"))
	fmt.Println("abcdcba: ", isPalindrome("abcdcba"))
	fmt.Println("A man a plan a canal Panama: ", isPalindrome("A man a plan a canal Panama"))
}
