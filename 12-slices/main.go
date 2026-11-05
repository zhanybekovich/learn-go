package main

import "fmt"

func main() {
	/*
		 * slice -> mutable reference type sequence of same data type values
			* inludes:
				* ptr -> pointer to the addr of base array
				* len -> length
				* cap -> capacity
	*/
	slice1 := make([]int, 3)
	// make([]Type, length, capacity)
	fmt.Println(slice1)

	// another way to init slice
	slice2 := []string{"Hi", "Hello"}
	fmt.Println(slice2)

	// or
	var slice3 []int
	fmt.Println(slice3)

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// add element
	slice2 = append(slice2, "Bye")
	fmt.Println(slice2)
}
