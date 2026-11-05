package main

import "fmt"

func main() {
	/*
		 * arrays
			* arrays -> data structure with fixed length
			* stores only same type values
			* varName := [n]Type{val1, val2, ...valN}
			* var newArr [n]Type
	*/

	var arr1 [3]int // initiated with default values
	fmt.Println(arr1)

	arr2 := [4]float32{4.12} // [4.12 0 0 0]
	fmt.Println(arr2)

	// change value by index
	arr1[0] = 2
	fmt.Println(arr1)

	// compiler counts the lenght
	arr3 := [...]float32{3.1, 4.5, 6.3}
	fmt.Println(arr3)
	fmt.Println(len(arr3))
}
