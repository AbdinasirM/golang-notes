package main

import (
	"fmt"
	"sort"
)

// A slice is a dynamic array in Go that can grow or shrink in size.
// An array is a static array with a fixed size defined at declaration.
func main() {
	// Declare and initialize a slice of strings
	colors := []string{"red", "blue", "green"}
	fmt.Println("Initial slice 'colors':", colors)

	// Add an element to the slice using the append function
	colors = append(colors, "grey")
	fmt.Println("Slice 'colors' after appending an element:", colors)

	// Remove the first element from the slice
	// Use slicing: colors[1:] creates a new slice starting from index 1 to the end
	colors = append(colors[1:len(colors)])
	fmt.Println("Slice 'colors' after removing the first element:", colors)

	// Remove the last element from the slice
	// Use slicing: colors[:len(colors)-1] creates a new slice excluding the last element
	colors = append(colors[:len(colors)-1])
	fmt.Println("Slice 'colors' after removing the last element:", colors)

	// Create a slice of integers with `make`
	// The `make` function takes three arguments: type, length, and capacity
	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 12
	numbers[2] = 32
	numbers[3] = 44
	numbers[4] = 0
	fmt.Println("Slice 'numbers' before sorting:", numbers)

	// Sort the slice of integers in ascending order
	sort.Ints(numbers)
	fmt.Println("Slice 'numbers' after sorting:", numbers)

	// Practical note:
	// The `sort.Ints` function modifies the slice in place, so no reassignment is needed.
}
