package main

import (
	"fmt"
)

func main() {
	str := "my string with double quotes escape symbols can be \t used inside"
	newStr := `my new \nstr escape symbols do not work`
	fmt.Println(str)
	fmt.Println(newStr)

	multiLineString := `This
	is a
	multiline
	string`
	fmt.Println(multiLineString)

	// indexes
	myValue := "Hello world!"
	fmt.Println(myValue[0])         // returns bytes
	fmt.Println(string(myValue[0])) // change byte to string

	// concat
	str1 := "ab"
	str2 := "bc"
	fmt.Println(str1 + " " + str2)

	// len
	fmt.Println(len(str2))
}
