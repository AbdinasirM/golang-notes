package main

import "fmt"

func main() {
	// Understanding Structs in Go

	// A struct is a composite data type in Go that groups together variables under one name.
	// These variables, known as fields, can have different types.

	// Note on naming:
	// - Variables or fields with names starting with a capital letter are exported,
	//   meaning they can be accessed from other packages.
	// - Variables or fields with names starting with a lowercase letter are unexported (private)
	//   and can only be accessed within the same package.

	// Create a new Car instance using struct literal syntax
	ford := Car{"Ford", "Mustang", 1965}
	fmt.Printf("Initial struct 'ford': %+v\n", ford) // %+v displays field names and values

	// Access and modify struct fields
	ford.Make = "Ford"          // Change the Make field
	ford.Model = "Explorer"     // Change the Model field
	ford.Released = 2021        // Change the Released field
	fmt.Printf("Updated struct 'ford': %+v\n", ford)

	// Create a struct instance using named fields for clarity
	tesla := Car{
		Make:     "Tesla",
		Model:    "Model S",
		Released: 2022,
	}
	fmt.Printf("Struct 'tesla': %+v\n", tesla)

	// Access individual fields of the struct
	fmt.Println("Car make:", tesla.Make)
	fmt.Println("Car model:", tesla.Model)
	fmt.Println("Car release year:", tesla.Released)
}

// Car is a struct
// A struct groups related data together as fields of different types.
type Car struct {
	Make     string // The make (manufacturer) of the car
	Model    string // The model name of the car
	Released int    // The release year of the car
}
