package bytes

import (
	"github.com/donmahallem/aoc/aoc_utils/math"
)

func ByteSequenceToInt[A math.IntType](b []byte) A {
	var num A = 0
	for i := 0; i < len(b); i++ {
		num = num*10 + A(b[i]-'0')
	}
	return num
}
