package day05

import (
	"io"
	"math"

	"github.com/donmahallem/aoc/aoc_utils"
)

func Part2(in io.Reader) int {
	almanac := ParseAlmanac(in)
	lowest := math.MaxInt
	for seedIdx := 0; seedIdx < len(almanac.Seeds); seedIdx += 2 {
		for i := range almanac.Seeds[seedIdx+1] {
			lowest = aoc_utils.Min(lowest, GetPosition(&almanac, almanac.Seeds[seedIdx]+i))
		}
	}
	return lowest
}
