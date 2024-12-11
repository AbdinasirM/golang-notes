package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date and time using time.Now()
	currentTime := time.Now()
	
	// Print a message along with the current date and time
	fmt.Println("I am learning Go programming now! The current date and time is:", currentTime)
	
	// Example: Formatting the current time to display it in a readable format
	// Format layout uses the reference date Mon Jan 2 15:04:05 MST 2006
	formattedTime := currentTime.Format("Monday, January 2, 2006 at 3:04 PM")
	fmt.Println("Formatted time:", formattedTime)
	
	// Example: Accessing individual components of the date and time
	year, month, day := currentTime.Date()
	hour, minute, second := currentTime.Clock()
	fmt.Printf("Year: %d, Month: %s, Day: %d\n", year, month, day)
	fmt.Printf("Time: %02d:%02d:%02d\n", hour, minute, second)
	
	// Example: Adding 1 day to the current time
	tomorrow := currentTime.Add(24 * time.Hour)
	fmt.Println("This time tomorrow will be:", tomorrow)
	
	// Example: Subtracting 1 week from the current time
	lastWeek := currentTime.Add(-7 * 24 * time.Hour)
	fmt.Println("The same time last week was:", lastWeek)
}
