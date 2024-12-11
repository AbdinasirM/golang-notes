package main

import "fmt"

func main() {
	// Declare a string variable and assign a value to it
	// "name" holds the name of the user
	var name string = "Abdi"
	// Print a greeting message that includes the name
	fmt.Printf("Hello, %s!\n", name)

	// Declare an integer variable and assign a value to it
	// "age" represents the user's age
	var age int = 23
	// Print a message displaying the user's age
	fmt.Printf("You are %d years old.\n", age)

	// Example of floating-point number usage
	// Declare a float variable for height
	var height float64 = 5.9
	// Print the height
	fmt.Printf("Your height is %.1f feet.\n", height)

	// Example of boolean usage
	// Declare a boolean variable to indicate if the user is a student
	var isStudent bool = true
	// Print whether the user is a student or not
	fmt.Printf("Are you a student? %t\n", isStudent)
}