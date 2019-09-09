package main

import "fmt"

// Multiples of 3 and 5 identifies the sum of all multiples
// of 3 & 5 below 1000.
func MultiplesOf3and5() (int, int) {
	limit := 1000
	// multiples of 3 & 5 below 1000
	sum3, sum5 := 0, 0
	for i := 1; i < limit; i++ {
		result3 := i * 3
		result5 := i * 5
		if result3 >= limit {
			break
		}

		sum3 += result3
		if result5 < limit {
			sum5 += result5
		}
	}

	newSum3, newSum5, newSum15 := 0, 0, 0
	for i := limit - 1; i > 2; i-- {
		if i%3 == 0 {
			newSum3 += i
		}
		if i%5 == 0 {
			newSum5 += i
		}
		if i%15 == 0 {
			newSum15 += i
		}
	}
	fmt.Println(newSum3, newSum5, newSum15, newSum3+newSum5-newSum15)

	return newSum3, newSum5
}
