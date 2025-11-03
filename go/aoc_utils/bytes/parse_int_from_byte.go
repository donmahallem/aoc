package bytes

import "github.com/donmahallem/aoc/go/aoc_utils/math"

func ParseIntFromByte[A math.IntType](b byte) (A, bool) {
	if ByteIsNumber(b) {
		return A(b - '0'), true
	}
	return 0, false
}
