package main

import "fmt"

func main() {
	number := 6
	result := NthPrimeNumber(number)
	fmt.Println("The Nth prime number", number, result)
	number = 10
	result = NthPrimeNumber(number)
	fmt.Println("The Nth prime number", number, result)
	number = 10001
	result = NthPrimeNumber(number)
	fmt.Println("The Nth prime number", number, result)
}
