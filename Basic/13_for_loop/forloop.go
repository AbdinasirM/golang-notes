package main

import "fmt"

func main() {
	// Example 1: Iterating over a slice using a standard for loop
	languages := []string{"Python3", "Javascript", "C++"}
	fmt.Println("Languages slice:", languages)

	// Three-part for loop (similar to traditional loops in other languages)
	fmt.Println("Using a three-part for loop:")
	for i := 0; i < len(languages); i++ {
		fmt.Println("Index", i, "Value:", languages[i])
	}

	// Example 2: Iterating using the range keyword
	// Range provides the index and value of each element in a slice
	fmt.Println("Using range in a for loop:")
	for i := range languages {
		fmt.Println("Index", i, "Value:", languages[i])
	}

	// Example 3: Simplified for-each style loop
	// Using `_` to ignore the index and directly access the value
	fmt.Println("For-each style loop:")
	for _, language := range languages {
		fmt.Println("Language:", language)
	}

	// Example 4: Using a for loop as a replacement for a while loop
	// A `for` loop without initialization and post statements acts like a while loop
	value := 1
	fmt.Println("Using for loop as a replacement for a while loop:")
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	// Example 5: Using break and continue
	// Demonstrates breaking out of a loop and skipping iterations
	sum := 1
	fmt.Println("Loop with break and continue:")
	for sum < 100 {
		sum++
		fmt.Println("Sum:", sum)

		if sum > 45 {
			// Using goto to exit the loop and jump to a label
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of program")
}
