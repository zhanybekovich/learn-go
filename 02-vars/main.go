package main

import "fmt"

func main() {
	// mannual variable defining
	var value int
	fmt.Println(value)

	// set initial value to the variable
	var value2 int = 10
	fmt.Println(value2)

	// define several vars
	var value3, value4 int = 100, 200
	fmt.Println(value3, value4)

	// or
	var value5, value6 = 300, 400
	fmt.Println(value5, value6)

	// or
	var (
		name     string = "Bob"
		age      int    = 34
		location string = "London"
	)
	fmt.Println(name, age, location)

	// or
	var (
		name2     = "Pete"
		age2      = 45
		location2 = "Moscow"
	)
	fmt.Println(name2, age2, location2)

	// shorthand variable defining
	city := "Bishkek"
	price := 45
	fmt.Println(city, price)

	/*
		naming conventions
		- if identifier will be used in another packages use MixedCaps
		- if identifier will be internal use mixedCaps
	*/
}
