package main

import (
	"fmt"
	"time"
)

func main() {
	// Example of using goroutines for concurrent execution

	// Start a goroutine to print numbers
	go printNumbers("Goroutine 1")

	// Start another goroutine to print letters
	go printLetters("Goroutine 2")

	// Main function waits to allow goroutines to complete
	// Without this, the program might exit before goroutines finish their execution
	time.Sleep(2 * time.Second)

	fmt.Println("Main function exiting...")
}

// printNumbers prints numbers from 1 to 5 with a delay
func printNumbers(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(300 * time.Millisecond) // Simulate some work
	}
}

// printLetters prints letters A to E with a delay
func printLetters(name string) {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%s: %c\n", name, i)
		time.Sleep(400 * time.Millisecond) // Simulate some work
	}
}

/*
	What are Goroutines?

	Goroutines are lightweight threads managed by the Go runtime. They allow functions or methods to run concurrently, making them ideal for writing highly concurrent programs.

	Key Features of Goroutines:
	1. Lightweight: Goroutines use minimal memory and are more efficient compared to traditional operating system threads.
	2. Concurrency: They enable multiple tasks to run simultaneously, improving performance in I/O-bound or CPU-bound operations.
	3. Non-blocking: Goroutines run independently of the main function and do not block its execution.

	How They Work:
	- The `go` keyword starts a new goroutine.
	- The function specified after `go` runs concurrently with the rest of the program.
	- The Go runtime dynamically schedules goroutines across available CPU cores.

	Important Note:
	- The main function must not exit before all goroutines finish. In this example, `time.Sleep` is used to ensure the goroutines have enough time to complete.
	- In real-world applications, synchronization tools like `sync.WaitGroup` or channels are often used to manage goroutine lifecycle and communication.
*/
