package bytes

import "github.com/donmahallem/aoc/go/aoc_utils/int_util"

func ParseIntFromByte[A int_util.IntType](b byte) (A, bool) {
	if ByteIsNumber(b) {
		return A(b - '0'), true
	}
	return 0, false
}
