package day22

import (
	"io"
)

type CacheValue struct {
	LastSeed *uint32
	Value    uint32
}
type CacheMap = map[uint32]CacheValue

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

func CreatePatterns(seed *uint32, iterations int, cache *CacheMap) {
	previousDiffs := make([]uint32, 0, 4)
	tmp := *seed
	previousValue := ((*seed) % 10)
	var currentValue uint32
	for i := 1; i < iterations; i++ {
		tmp = Step(tmp)
		currentValue = tmp % 10
		if len(previousDiffs) == 4 {
			previousDiffs = previousDiffs[1:]
		}
		previousDiffs = append(previousDiffs, 10+currentValue-previousValue)
		if i >= 4 {
			key := EncodeSequence(&previousDiffs)
			if val, ok := (*cache)[key]; ok {
				// prevent duplicates from same seed
				if val.LastSeed != seed {
					val.LastSeed = seed
					val.Value = val.Value + currentValue
					(*cache)[key] = val
				}
			} else {
				(*cache)[key] = CacheValue{LastSeed: seed, Value: currentValue}
			}
		}
		previousValue = currentValue
	}
}

func Part2(in io.Reader) uint32 {
	items := ParseInput(in)
	cache := make(CacheMap)
	for _, item := range *items {
		CreatePatterns(&item, 2000, &cache)
	}
	maxVal := uint32(0)
	for key := range cache {
		if cacheValue := cache[key]; cacheValue.Value > maxVal {
			maxVal = cacheValue.Value
		}
	}
	return maxVal
}
