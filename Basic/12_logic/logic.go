package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a new reader to read user input from the standard input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for their age
	fmt.Print("How old are you? ")

	// Read the user's input until a newline character
	input, _ := reader.ReadString('\n')

	// Remove any surrounding whitespace or newline characters from the input
	input = strings.TrimSpace(input)

	// Convert the cleaned input string into an integer
	age, err := strconv.Atoi(input)
	if err != nil {
		// Handle the case where the input is not a valid number
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Check if the user meets the age requirement
	checkAge(age)
}

// checkAge determines if a user can access the platform based on their age
func checkAge(age int) {
	if age >= 18 {
		// User is 18 or older
		fmt.Println("You can access this platform.")
	} else {
		// User is under 18
		fmt.Println("You cannot access this platform.")
	}
}
