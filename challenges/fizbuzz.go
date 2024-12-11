package main

import "fmt"

func main(){	
	number := 9

	fizzbuzz(number)
}


func fizzbuzz(number int) {
	if number%3 == 0 && number%5 == 0{
		fmt.Println("FizzBuzz")
	}else if number % 3 == 0 {
		fmt.Println("Fizz")
	}else if number % 5 == 0 {
		fmt.Println("Buzz")
	}else {
		fmt.Println("Please enter a number")
	}
}