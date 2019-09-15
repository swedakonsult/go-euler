// 10,001st prime number
package main

func NthPrimeNumber(n int) int {
	primes := getPrimes(n*n, n)

	return primes[n-1]
}

func getPrimes(largestNumber int, stopAt int) []int {
	primes := calculatePrimes(largestNumber)
	primeNumbers := []int{}

Outer:
	for _, number := range primes {
		for _, prime := range primeNumbers { // get rid of any extra composites
			if number%prime == 0 { // composite number after all, so skip
				continue Outer
			}
		}
		primeNumbers = append(primeNumbers, number)
		if len(primeNumbers) == stopAt {
			return primeNumbers
		}
	}

	return primeNumbers
}
