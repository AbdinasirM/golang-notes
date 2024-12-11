package main

import "fmt"

func main() {
	// Understanding Maps in Go
	// A map is an unordered collection of key-value pairs.
	// Keys must be unique, and each key maps to exactly one value.

	// Create a map with keys and values as strings using the make function
	states := make(map[string]string) // Initialize an empty map
	fmt.Println("Initial map 'states':", states)

	// Add key-value pairs to the map
	states["MN"] = "Minnesota"
	states["OR"] = "Oregon"
	states["CH"] = "Chicago" // Chicago should be "IL" (Illinois) in real-world use
	fmt.Println("Map 'states' after adding elements:", states)

	// Access a value using its key
	minnesota := states["MN"] // Retrieve the value associated with the key "MN"
	fmt.Println("Value associated with key 'MN':", minnesota)

	// Delete a key-value pair from the map
	delete(states, "OR") // Remove the key "OR" and its associated value
	fmt.Println("Map 'states' after deleting key 'OR':", states)

	// Iterate over each key-value pair in the map using a for loop
	fmt.Println("Iterating over map 'states':")
	for key, value := range states {
		fmt.Printf("%v: %v\n", key, value)
	}

	// Check if a key exists in the map
	// The second return value indicates if the key is present in the map
	value, exists := states["OR"]
	if exists {
		fmt.Println("Key 'OR' exists with value:", value)
	} else {
		fmt.Println("Key 'OR' does not exist in the map.")
	}

	// Example of initializing a map with literal syntax
	otherStates := map[string]string{
		"CA": "California",
		"TX": "Texas",
		"NY": "New York",
	}
	fmt.Println("Another map 'otherStates':", otherStates)
}
