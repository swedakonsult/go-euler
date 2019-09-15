// Smalest Multiple
package main

func SmallestMultiple(limit int) int {
	result := 0

	multipliers := make([]int, limit-1)
	for i := limit; i > 1; i-- {
		multipliers[i-2] = i
	}

	start := limit * 2 // primary number times the limit would be the smallest number

Outer:
	for i := start; i < 232342342342; {
		result = i
		i = result + limit
		for _, multiplier := range multipliers {
			if result%multiplier != 0 {
				continue Outer
			}
		}
		break
	}

	return result
}
