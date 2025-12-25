package day23

import (
	"io"
)

const PREFIX_T uint16 = uint16('t')

func StartsWithT(key NodeHash) bool {
	return ((key) >> 8) == PREFIX_T
}

func FindTriplets(data CombinationMap) uint16 {
	results := uint16(0)
	var validEntry1, validEntry2 bool
	for key1 := range data {
		dataKey1 := data[key1]
		validEntry1 = StartsWithT(key1)
		for key2 := range dataKey1 {
			validEntry2 = validEntry1 || StartsWithT(key2)
			dataKey2 := data[key2]
			for key3 := range dataKey2 {
				if validEntry2 || StartsWithT(key3) {
					if _, ok := dataKey1[key3]; ok {
						results++
					}
				}
			}
		}
	}
	return results
}

func Part1(in io.Reader) (uint16, error) {
	items, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	return FindTriplets(items), nil
}
