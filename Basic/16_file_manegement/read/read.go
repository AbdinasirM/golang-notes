package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Using defer to ensure the readFile function is called at the end of main
	// This demonstrates deferred execution, which is useful for cleanup operations
	defer readFile("../fromString.txt")
}

// readFile reads the content of a file and prints it to the console
func readFile(fileName string) {
	// ReadFile reads the entire file content into memory as a byte slice
	data, err := ioutil.ReadFile(fileName)

	// Check for errors and handle them
	checkError(err)

	// Print the file content as a string
	fmt.Println("Text read from file:", string(data))
}

// checkError handles any errors that occur during file operations
func checkError(err error) {
	if err != nil {
		// panic stops the program and prints the error message
		panic(err)
	}
}
