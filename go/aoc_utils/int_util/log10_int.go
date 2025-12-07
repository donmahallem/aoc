package int_util

import "math"

func Log10Int[T IntType](n T) T {
	return T(math.Log10(float64(n)))
}
