// Sum square difference
package main

func SumSquaredDifference(limit int) int {
	sumOfSquares, sum := 0, 0
	for i := 1; i <= limit; i++ {
		sumOfSquares += i * i
		sum += i
	}

	return sum*sum - sumOfSquares
}
