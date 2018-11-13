package main

import (
	"reflect"
	"fmt"
)


// IsCollection returns if the argument is a collection.
func IsCollection(in interface{}) bool {
	arrType := reflect.TypeOf(in)

	kind := arrType.Kind()

	return kind == reflect.Array || kind == reflect.Slice
}


func main() {
	
	fmt.Println(IsCollection(map[string]int{"a": 1, "b": 2}))
	fmt.Println(IsCollection([]int{21, 22, 23}))
	
}
