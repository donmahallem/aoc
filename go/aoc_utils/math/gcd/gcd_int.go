package gcd

import (
	"github.com/donmahallem/aoc/go/aoc_utils/math"
)

func GcdInt[T math.IntType](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
