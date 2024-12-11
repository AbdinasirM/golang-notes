package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare and initialize three integer variables
	i1, i2, i3 := 11, 43, 68

	// Calculate the sum of the integers
	int_sum := i1 + i2 + i3
	// Print the sum of the integers
	fmt.Println("The total sum of integers is:", int_sum)

	// Convert the sum to a floating-point number and round it
	total := math.Round(float64(int_sum))
	// Print the rounded total (though it will be the same in this case since no decimals are added)
	fmt.Println("The rounded total is:", total)

	// Example: Find the average of the integers
	average := float64(int_sum) / 3
	fmt.Printf("The average is: %.2f\n", average)

	// Example: Calculate the square root of the sum
	sqrt_sum := math.Sqrt(float64(int_sum))
	fmt.Printf("The square root of the sum is: %.2f\n", sqrt_sum)

	// Example: Calculate the power of the first number raised to the second number
	power := math.Pow(float64(i1), float64(i2))
	fmt.Printf("%d raised to the power of %d is: %.2f\n", i1, i2, power)

	// Example: Using the modulo operator to find the remainder
	remainder := i3 % i1
	fmt.Printf("The remainder when %d is divided by %d is: %d\n", i3, i1, remainder)
}