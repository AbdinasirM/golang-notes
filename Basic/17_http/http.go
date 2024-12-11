package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// URL to fetch the sample TODOs data
const url = "https://jsonplaceholder.typicode.com/todos"

func main() {
	// Send an HTTP GET request to fetch data from the URL
	response, err := http.Get(url)
	checkError(err) // Check for errors in the HTTP response

	// Print the response type (to illustrate that response is of type *http.Response)
	fmt.Printf("Response type: %T\n", response)

	// Ensure the response body is closed to free up resources
	defer response.Body.Close()

	// Read the response body into a byte slice
	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err) // Handle any errors during reading

	// Convert the byte slice to a string for further processing
	contents := string(bytes)
	// fmt.Println(contents) // Uncomment to print the raw JSON response

	// Convert the JSON string to a slice of Todos structs
	todos := TodosFromJSON(contents)

	// Iterate over the Todos slice and print the titles of each TODO
	fmt.Println("List of TODO Titles:")
	for _, todo := range todos {
		fmt.Println(todo.Title)
	}
}

// checkError is a helper function to handle errors gracefully
func checkError(err error) {
	if err != nil {
		// Panic stops execution and provides an error message
		panic(err)
	}
}

// TodosFromJSON decodes a JSON string into a slice of Todos structs
func TodosFromJSON(content string) []Todos {
	// Initialize an empty slice with a capacity of 100 to hold TODO items
	todos := make([]Todos, 0, 100)

	// Create a JSON decoder for the input string
	decoder := json.NewDecoder(strings.NewReader(content))

	// Read the opening JSON token (e.g., '[' for an array)
	_, err := decoder.Token()
	checkError(err)

	// Iterate over the JSON elements and decode each TODO item
	var todo Todos
	for decoder.More() {
		err := decoder.Decode(&todo) // Decode the current item into the `todo` struct
		checkError(err)
		todos = append(todos, todo) // Append the decoded item to the slice
	}
	return todos
}

// Todos is a struct representing a TODO item
type Todos struct {
	UserId    int    // User ID associated with the TODO
	Id        int    // Unique ID for the TODO
	Title     string // Title of the TODO
	Completed bool   // Status of whether the TODO is completed
}
