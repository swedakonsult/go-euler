package main

import "fmt"

func main() {
	number := 10
	result := SumSquaredDifference(number)
	fmt.Println("Difference of sum squared", number, result)
	number = 100
	result = SumSquaredDifference(number)
	fmt.Println("Difference of sum squared", number, result)
}
