package main

import "fmt"

func main() {
	number := 12
	a, b, c := FindPythagoreanTriplet(number)
	fmt.Println("Pythagorean triplet product", number, a*b*c)
	number = 1000
	a, b, c = FindPythagoreanTriplet(number)
	fmt.Println("Pythagorean triplet product", number, a*b*c)
}
