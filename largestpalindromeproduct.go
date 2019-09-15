// Largest palindrome product
package main

import "fmt"

func LargestPalindromeProduct() int {
	digitOne, digitTwo := 999, 999
	largestProduct, d1, d2 := 0, 0, 0
	for i := digitOne; i > 1; i-- {
		for j := digitTwo; j > 1; j-- {
			product := i * j
			if product <= largestProduct {
				break
			}
			if isPalindrome(product) {
				largestProduct = product
				d1 = i
				d2 = j
				break
			}
		}
	}

	fmt.Println("Found number with digits", largestProduct, d1, d2)

	return largestProduct
}

func isPalindrome(candidate int) bool {
	stringVersion := fmt.Sprintf("%v", candidate)
	return isStringPalindrome(stringVersion)
}

func isStringPalindrome(s string) bool {
	if len(s) < 3 {
		return false
	}

	if len(s)%2 == 0 { // even
		for i, j := 0, len(s)-1; i < j; i++ {
			if s[i] != s[j] {
				return false
			}
			j--
		}
	} else {
		for i, j := 0, len(s)-1; i < j-1; i++ {
			if s[i] != s[j] {
				return false
			}
			j--
		}
	}

	return true
}
