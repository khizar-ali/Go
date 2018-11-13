package main

import "strings"
import "fmt"

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}


func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))
}