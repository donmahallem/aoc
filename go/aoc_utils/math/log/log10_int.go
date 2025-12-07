package log

import (
	"math"

	util_math "github.com/donmahallem/aoc/go/aoc_utils/math"
)

func Log10Int[T util_math.IntType](n T) T {
	return T(math.Log10(float64(n)))
}
