package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel for sending and receiving integers
	dataChannel := make(chan int)

	// Start a goroutine to send data into the channel
	go sendData(dataChannel)

	// Start a goroutine to receive data from the channel
	go receiveData(dataChannel)

	// Allow some time for goroutines to complete before the program exits
	time.Sleep(2 * time.Second)

	fmt.Println("Main function exiting...")
}

// sendData sends integers to the channel
func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sending: %d\n", i)
		ch <- i // Send data into the channel
		time.Sleep(300 * time.Millisecond) // Simulate some work
	}
	close(ch) // Close the channel when done sending data
}

// receiveData receives integers from the channel
func receiveData(ch chan int) {
	for val := range ch { // Read values from the channel until it is closed
		fmt.Printf("Received: %d\n", val)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

/*
	What are Channels in Go?

	Channels are a way to communicate between goroutines in Go. They provide a mechanism for goroutines to send and receive data with each other, enabling synchronization and coordination.

	Key Features of Channels:
	1. Type-Specific: Channels can transmit values of a specific type, such as `int`, `string`, or custom structs.
	2. Blocking Nature:
		- Sending to a channel blocks until another goroutine receives the value.
		- Receiving from a channel blocks until there is data available in the channel.
	3. Thread-Safe: Channels are designed to be safe for concurrent access, making them ideal for goroutine communication.

	How They Work:
	- A channel is created using the `make` function, e.g., `make(chan int)`.
	- Data is sent into a channel using the `<-` operator, e.g., `ch <- value`.
	- Data is received from a channel using the `<-` operator, e.g., `val := <-ch`.
	- Channels can be closed using the `close` function to indicate that no more data will be sent.

	Important Notes:
	1. The `range` keyword can be used to read all values from a channel until it is closed.
	2. Closing a channel is essential to signal the end of data transmission. Receivers detect the closure and exit the loop.

	This example demonstrates how a sender goroutine sends integers to a channel, and a receiver goroutine reads and processes those integers.
*/
