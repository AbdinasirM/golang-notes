package main

import "fmt"

func main() {
    fmt.Println("Demonstrating map initialization and usage in Go")

    // Example: Using 'new' to initialize a map (incorrect usage)
    // The example below demonstrates what happens if 'new' is used incorrectly to create a map.
    // Uncommenting the following lines will cause a runtime crash, as the map is not properly initialized.
    // m := new(map[string]int) // 'new' allocates memory but does not initialize the map.
    // (*m)["age"] = 23 // Causes a panic: assignment to entry in nil map.
    // fmt.Println(*m)

    // Correct way: Using 'make' to initialize a map
    // 'make' creates and initializes the map properly, allowing it to be used safely.
    m := make(map[string]int)
    m["age"] = 23 // Adding a key-value pair to the map
    fmt.Printf("Map after adding a key-value pair: %v\n", m)

    // Accessing a value from the map
    fmt.Printf("Value associated with key 'age': %d\n", m["age"])

    // Deleting a key-value pair from the map
    delete(m, "age")
    fmt.Printf("Map after deleting the key 'age': %v\n", m)

    // Note: Memory is deallocated automatically by Go's garbage collector,
    // so there is no need for explicit memory management.
}
