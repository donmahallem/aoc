package day22

import (
	"io"
)

func CreatePatterns(seed *uint32, iterations int) *[]int8 {
	values := make([]int8, 1, iterations)
	tmp := *seed
	values[0] = int8(*seed) % 10
	for i := 1; i < iterations; i++ {
		tmp = Step(tmp)
		values = append(values, int8(tmp%10))
	}
	return &values
}

const encodeBase19Shift3 uint32 = 6859
const encodeBase19Shift2 uint32 = 361
const encodeBase19Shift1 uint32 = 19

func EncodeSequence(b *[]uint32) uint32 {
	tmp := (*b)[0] * encodeBase19Shift3
	tmp += (*b)[1] * encodeBase19Shift2
	tmp += (*b)[2] * encodeBase19Shift1
	tmp += (*b)[3]
	return tmp
}

type CacheMap = map[uint32]uint32

func CreatePatterns2(seed *uint32, iterations int, cache *CacheMap) {
	values := make([]uint32, 1, iterations)
	previousValue := make([]uint32, 0, 4)
	tmp := *seed
	values[0] = ((*seed) % 10)
	for i := 1; i < iterations; i++ {
		tmp = Step(tmp)
		values = append(values, tmp%10)
		if len(previousValue) == 4 {
			previousValue = previousValue[1:]
		}
		previousValue = append(previousValue, 10+values[len(values)-1]-values[len(values)-2])
		if i >= 4 {
			key := EncodeSequence(&previousValue)
			if val, ok := (*cache)[key]; ok {
				(*cache)[key] = val + values[i-1]
			} else {
				(*cache)[key] = values[i-1]
			}
		}
	}
}
func CreatePatternDifferences(seed *uint32, iterations int) *[]int8 {
	buffer := make([]int8, 0, iterations-1)
	tmp := *seed
	var tmp2 uint32
	for range iterations - 1 {
		tmp2 = Step(tmp)
		buffer = append(buffer, int8(tmp2%10-tmp%10))
		tmp = tmp2
	}
	return &buffer
}

func Part2(in io.Reader) uint {
	items := ParseInput(in)
	return AddUpSecrets(items)
}
