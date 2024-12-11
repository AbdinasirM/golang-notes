package main

import "fmt"

func main() {
	// Creating an instance of the Dog struct
	poodle := Dog{"Poodle", 10, "Woof!"}

	// Print the entire struct using formatted output
	fmt.Printf("Dog struct: %v\n", poodle)

	// Access and print individual fields of the struct
	fmt.Printf("Breed: %v\nWeight: %v kg\n", poodle.Breed, poodle.Weight)

	// Call the Speak method of the Dog struct
	poodle.Speak()
}

// Dog is a struct that represents information about a dog
type Dog struct {
	Breed  string // The breed of the dog
	Weight int    // The weight of the dog in kilograms
	Sound  string // The sound the dog makes
}

// Speak is a method associated with the Dog struct
// It prints the sound the dog makes
func (d Dog) Speak() {
	fmt.Println("The dog says:", d.Sound)
}
