package main

import "fmt"

func main() {
	// Understanding Pointers in Go

	// A pointer is a variable that stores the memory address of another variable.
	// If a pointer is declared but not initialized, it defaults to a nil value.
	var p *int
	fmt.Println("The value of p is:", p) // This will print <nil> because p has not been assigned any memory address.

	// Let's demonstrate pointers with an example
	anumb := 45 // Declare a regular integer variable

	// The & operator is used to get the memory address of a variable.
	// Here, numpointer stores the memory address of anumb.
	numpointer := &anumb
	fmt.Println("Memory address of anumb (stored in numpointer):", numpointer)

	// The * operator is used to dereference a pointer,
	// meaning it fetches the value stored at the memory address.
	fmt.Println("Value pointed to by numpointer (value of anumb):", *numpointer)

	// Updating the value of the variable through the pointer
	*numpointer = 50
	fmt.Println("Updated value of anumb through numpointer:", anumb)

	// Example of nil pointer check
	if p == nil {
		fmt.Println("The pointer 'p' is nil, and it does not point to any memory address.")
	}
}
