package main

import "fmt"

func main() {
	number := 2
	result := LargestProductInMatrix(LargestProductMatrix, number)
	fmt.Println("Largest product", number, result)
	number = 4
	result = LargestProductInMatrix(LargestProductMatrix, number)
	fmt.Println("Largest product", number, result)
}
