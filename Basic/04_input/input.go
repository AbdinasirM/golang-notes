package main

import (
	"bufio" // Package for buffered I/O
	"fmt"   // Package for formatted I/O
	"os"    // Package for OS-level functionality
)

func main() {
	// Create a new reader to read input from the standard input (console)
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter something
	fmt.Print("Enter something: ")

	// Read a string of input from the user until a newline character is encountered
	input, err := reader.ReadString('\n')
	if err != nil { // Check for any errors during input
		fmt.Println("Error reading input:", err)
		return
	}

	// Print what the user entered, trimming any unnecessary newline
	fmt.Println("You entered:", input)
}
