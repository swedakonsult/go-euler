package main

import "fmt"

func main() {
	number := 10
	result := SumOfPrimes(number)
	fmt.Println("Sum of primes below", number, result)
	number = 2000000
	result = SumOfPrimes(number)
	fmt.Println("Sum of primes below", number, result)
}
