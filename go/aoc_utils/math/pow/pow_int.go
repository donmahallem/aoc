package pow

import "github.com/donmahallem/aoc/aoc_utils/math"

func IntPow[T math.IntType](base T, exp T) T {
	var result T = 1
	for exp > 0 {
		if exp%2 == 1 { // If the exponent is odd
			result *= base
		}
		base *= base // Square the base
		exp /= 2
	}
	return result
}
