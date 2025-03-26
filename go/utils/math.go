package utils

import "math"

type IntType interface {
	int | int8 | int16 | int32 | int64
}

func Abs[T IntType](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func IntPow[T IntType](base T, exp T) T {
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

func Log10Int[T IntType](n T) T {
	return T(math.Log10(float64(n))) + 1
}
