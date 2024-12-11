package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Define the content to be written to the file
	content := "Hello from GO!"

	// Create a new file or overwrite if it already exists
	// os.Create creates the file and returns a file descriptor
	file, err := os.Create("./fromString.txt")
	checkError(err) // Handle any errors during file creation

	// Write the string content to the file
	// io.WriteString writes the string content to the file and returns the number of bytes written
	length, err := io.WriteString(file, content)
	checkError(err) // Handle any errors during writing

	// Print the number of characters written to the file
	fmt.Printf("Wrote a file with %v characters\n", length)

	// Use defer to ensure the file is closed after all operations are completed
	defer file.Close()
}

// checkError is a helper function to handle errors
func checkError(err error) {
	if err != nil {
		// Panic stops the program execution and outputs the error message
		panic(err)
	}
}
