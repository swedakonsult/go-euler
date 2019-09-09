package main

import "fmt"

// Event Fibonaccy Numbers for where the values do not
// exceed 4,000,000
func EvenFibonacciNumbers() int {
	limit := 4000000

	evenSum := 2
	previous := 1
	for f := 2; f < limit; {
		m := nextFibonacci(previous, f)
		previous = f
		f = m
		if m%2 == 0 {
			evenSum += m
			fmt.Println(m)
		}
	}

	return evenSum
}

func nextFibonacci(twoBack int, last int) int {
	return twoBack + last
}
