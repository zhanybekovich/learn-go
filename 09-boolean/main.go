package main

import "fmt"

func main() {
	// boolean
	var myBool bool
	fmt.Println(myBool) // default false

	myBool2 := true
	fmt.Println(myBool2)

	// logical operators with bool
	var value bool
	fmt.Println(!value)          // not
	fmt.Println(2 > 3 || 4 < 10) // or
	fmt.Println(2 > 3 && 4 < 10) // and
}
