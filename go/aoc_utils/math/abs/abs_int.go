package abs

import "github.com/donmahallem/aoc/go/aoc_utils/math"

func AbsInt[T math.IntType](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
