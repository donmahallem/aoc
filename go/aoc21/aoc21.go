package aoc21

import (
	"github.com/donmahallem/aoc/go/aoc21/day01"
	"github.com/donmahallem/aoc/go/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	aoc_utils.RegisterDay(registry, 21, 1, day01.Part1, day01.Part2)
}
