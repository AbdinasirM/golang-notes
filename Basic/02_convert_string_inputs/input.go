package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a new reader to take input from the user
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter some text
	fmt.Print("Enter something: ")
	input, _ := reader.ReadString('\n') // Read the input until a newline character
	// Print the entered input
	fmt.Println("You entered:", input)

	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")
	numbInput, _ := reader.ReadString('\n') // Read the number as a string

	// Convert the entered string to a floating-point number
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numbInput), 64) // Trim spaces and parse as float

	if err != nil {
		// Handle the error if the input is not a valid number
		fmt.Println("Error:", err)
	} else {
		// Print the parsed floating-point number
		fmt.Println("Value of number:", aFloat)
	}
}