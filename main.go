package main

import "fmt"

func main() {
	// sum3, sum5 := MultiplesOf3and5()
	// sum := EvenFibonacciNumbers()
	// sum := LargestPrimeFactor(600851475143)
	number := 600851475143
	largestFactor := LargestPrimeFactor(number)
	fmt.Println("Number", number, "has largest factor", largestFactor)
}
