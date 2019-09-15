package main

// a + b + c = sum, a^2 + b^2 = c^2, a < b < c
func FindPythagoreanTriplet(sum int) (a int, b int, c int) {
	a = 0
	b = a + 1
	c = b + 1

Outer:
	for ; a < sum; a++ {
		for b = a + 1; b < sum; b++ {
			for c = b + 1; c < sum; c++ {
				if a+b+c == sum { // second check
					if a*a+b*b == c*c { // pythagorean
						break Outer
					}
				}
			}
		}
	}

	return
}
