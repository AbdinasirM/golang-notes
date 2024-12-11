package main

import "fmt"

// A slice is a dynamic array in Go that can grow or shrink in size.
// An array is a static, fixed-size sequence of elements.
func main() {
	// Declare a slice of strings and initialize it properly
	colors := make([]string, 3) // Allocate a slice with an initial size of 3
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	fmt.Println("Slice 'colors':", colors) // Print the slice

	// Adding elements dynamically to a slice using the append function
	colors = append(colors, "yellow", "purple")
	fmt.Println("Updated slice 'colors' after appending elements:", colors)

	// Declare and initialize an array of integers with a fixed size
	var numbs = [4]int{1, 2, 3, 4} // Fixed-size array
	fmt.Println("Array 'numbs':", numbs)

	// Get the length of the array
	fmt.Println("Length of the array 'numbs':", len(numbs))

	// Comparing slices and arrays:
	// Slices are more flexible as they can grow dynamically,
	// while arrays have a fixed size and cannot be resized.
}
