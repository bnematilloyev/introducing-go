package main

import "fmt"

/*
	func zero(x int) {
		x = 0
	}
	func main() {
		// Pointers
		x := 5
		zero(x)
		fmt.Println(x) // x is not changed, still 5
	}
*/

// Pointers - reference a location in memory
// By using a pointer (*int), the zero function is able to modify the original variable
/*
	func zero(xPtr *int) {
		*xPtr = 10
	}

	func main() {
		x := 5
		zero(&x)
		fmt.Println(x)
	}
*/

// The * and & operators
// asterisk (*)
// we use the & operator to find the address of a variable
// &x in main and xPtr in zero refer to the same memory location.

// new
func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	// new takes a type as an argument, allocates enough memory to fit a value of that type,
	// and returns a pointer to it.
	// Go is a garbage-collected programming language.
	// which means memory is cleaned up automatically when nothing refers to it
	// anymore

	/*
		Pointers are rarely used with Go’s built-in types, but as we will see in the next chapter,
		they are extremely useful when paired with structs.
	*/
}
