package main

import "fmt"

func main() {
	/*
		 * Types in Go
			* basic
				* integers -> signed or unsigned
				* floats
				* complex numbers
				* byte
				* rune
				* string
				* booleans
			* composite
				* arrays
				* structs
				* slices
				* maps
				* channels
				* pointers
				* functions
				* methods
				* interface
	*/

	/*
		 * integers (signed)
			* int -> 32bit = 4 bytes; 64bit = 8 bytes
			* int8 -> 1 byte
			* int16 -> 2 bytes
			* int32 -> 4 bytes
			* int64 -> 8 bytes
		* integers (unsigned)
			* uint -> 32bit = 4 bytes; 64bit = 8 bytes
			* uint8 -> 1 byte
			* uint16 -> 2 bytes
			* uint32 -> 4 bytes
			* uint64 -> 8 bytes
	*/

	value := 0xA1 // hex number
	fmt.Println(value)

	value2 := 0b1011 // binary
	fmt.Println(value2)
}
