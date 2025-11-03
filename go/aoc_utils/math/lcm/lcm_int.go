package lcm

import (
	"github.com/donmahallem/aoc/go/aoc_utils/math"
	"github.com/donmahallem/aoc/go/aoc_utils/math/abs"
	"github.com/donmahallem/aoc/go/aoc_utils/math/gcd"
)

func LcmInt[T math.IntType](a, b T) T {
	return abs.AbsInt(a*b) / gcd.GcdInt(a, b)
}
