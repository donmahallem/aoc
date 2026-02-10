package int_util

import "math"

func Log10Int[T IntType](n T) T {
	if n == 0 {
		return 0
	}
	return T(math.Log10(float64(n)))
}
