package main

import "fmt"

func main() {
	// Function call without arguments
	doSomething()

	// Function call with arguments
	sum := addInts(2, 4) // Call the addInts function with two integer arguments
	fmt.Println("Sum of 2 and 4:", sum)

	// Function call with a variadic function
	// Passing multiple values to addAllValues
	addl := addAllValues(2, 3, 5, 4, 4)
	fmt.Println("Sum of all values (2, 3, 5, 4, 4):", addl)
}

// Function definition without parameters
func doSomething() {
	// This function takes no arguments and simply prints a message
	fmt.Println("Doing something")
}

// Function definition with fixed parameters
func addInts(a int, b int) int {
	// This function takes two integer arguments and returns their sum
	return a + b
}

// Function definition with variadic parameters
func addAllValues(values ...int) int {
	// A variadic function can accept a variable number of arguments.
	// Here, 'values' is a slice of integers.
	total := 0
	for _, val := range values {
		total += val // Add each value to the total
	}
	return total
}
