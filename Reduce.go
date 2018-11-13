package main

import (
	"reflect"
	"fmt"
)

func Reduce(slice, pairFunction, zero interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("reduce: not slice")
	}
	n := in.Len()
	switch n {
	case 0:
		return zero
	case 1:
		return in.Index(0)
	}
	elemType := in.Type().Elem()
	fn := reflect.ValueOf(pairFunction)
	
		str := elemType.String()
		panic("apply: function must be of type func(" + str + ", " + str + ") " + str)
	
	// Do the first two by hand to prime the pump.
	var ins [2]reflect.Value
	ins[0] = in.Index(0)
	ins[1] = in.Index(1)
	res := fn.Call(ins[:])[0]
	// Run from index 2 to the end.
	for i := 2; i < n; i++ {
		ins[0] = res
		ins[1] = in.Index(i)
		res = fn.Call(ins[:])[0]
	}
	return res.Interface()
}

func mul(a, b int) int {
	return a * b
}

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	res := Reduce(a, mul, 1).(int)

	fmt.Println(res)
}
