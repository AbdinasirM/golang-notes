package main

import (
	"fmt"
	"strings"
)


func main(){
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words:= strings.Fields(text)
	counts := map[string]int{}

	for _, word := range words{ 
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)



	var maxWord string
	max_count := 0

	for word, count := range counts{
		if count > max_count{
			maxWord = word
			max_count = count
		}
	}

	fmt.Println("The most appering word is: ", maxWord)

}