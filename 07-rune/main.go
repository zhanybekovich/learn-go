package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		 * rune type -> is a kind of alias for int32 type
			* used to store symbols in UTF-8 format
	*/

	var value = 'ө'
	fmt.Println(value, reflect.TypeOf(value))

	value2 := '\u2318'
	fmt.Println(value2, reflect.TypeOf(value2))
}
