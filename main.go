package main

import "fmt"

func main() {
	number := 4
	result := LargestProductInSeries(LargestProductSeries, number)
	fmt.Println("Largest product", number, result)
	number = 13
	result = LargestProductInSeries(LargestProductSeries, number)
	fmt.Println("Largest product", number, result)
}
