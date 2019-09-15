package main

import "fmt"

func main() {
	number := 10
	result := SmallestMultiple(number)
	fmt.Println("Smallest multiple", number, result)
	number = 20
	result = SmallestMultiple(number)
	fmt.Println("Smallest multiple", number, result)
}
