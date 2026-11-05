package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		 * floats
			* float32 -> 32 bits
			* float64 -> 8 bytes
	*/

	var value = 10.1
	fmt.Println(value, reflect.TypeOf(value))
}
