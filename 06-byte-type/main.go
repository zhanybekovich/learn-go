package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		 * byte type is a kind of alias for uint8 type
			* mostly used to store symbols of ASCII format
	*/

	var value byte = 'd'
	fmt.Println(value, reflect.TypeOf(value))
}
