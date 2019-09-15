package main

import (
	"math"
)

var primeNumbers = make(map[int]int)

func isPrime(number int) bool {
	if _, ok := primeNumbers[number]; ok {
		return true
	}

	return false
}

func LargestPrimeFactor(number int) int {
	// Sleve of Eratosthenes - stop when the square is larger than the wanted number: 775,147
	// remove numbers from the full set

	largestNumber := number
	if number > 1000 {
		largestNumber = int(math.Ceil(math.Pow(float64(number), 0.5)))
	}

	preparePrimes(largestNumber)

	for prime, _ := range primeNumbers {
		if number%prime == 0 {
			return prime
		}
	}
	return 1
}

func preparePrimes(largestNumber int) {
	primes := calculatePrimes(largestNumber)

	for i := len(primes) - 1; i >= 0; i-- {
		number := primes[i]
		primeNumbers[number] = number
	}
}

func calculatePrimes(largestNumber int) []int {

	primeCandidates := make([]int, largestNumber)
	currentIndex := 0
	primeCandidates[currentIndex] = 2
	currentIndex++
	if largestNumber < 3 {
		return primeCandidates[:currentIndex]
	}
	primeCandidates[currentIndex] = 3
	currentIndex++
	if largestNumber < 5 {
		return primeCandidates[:currentIndex]
	}
	primeCandidates[currentIndex] = 5
	currentIndex++
	if largestNumber < 7 {
		return primeCandidates[:currentIndex]
	}
	primeCandidates[currentIndex] = 7
	currentIndex++
	if largestNumber < 9 {
		return primeCandidates[:currentIndex]
	}
	for i := currentIndex; i < largestNumber; i++ {
		// optimize for some known prime values
		if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 {
			primeCandidates[currentIndex] = i
			currentIndex++
		}
	}

	return primeCandidates[:currentIndex]
}
