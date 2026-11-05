package main

import "fmt"

func main() {
	arr2D := [3][3]int{{2}, {3, 6, 9}}
	fmt.Println(arr2D)

	arr2D[2][2] = 5
	fmt.Println(arr2D)
}
