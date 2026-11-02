package aoc21

import (
	"github.com/donmahallem/aoc/go/aoc21/day01"
	"github.com/donmahallem/aoc/go/aoc21/day02"
	"github.com/donmahallem/aoc/go/aoc21/day03"
	"github.com/donmahallem/aoc/go/aoc21/day04"
	"github.com/donmahallem/aoc/go/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	aoc_utils.RegisterDay(registry, 21, 1, day01.Part1, day01.Part2)
	aoc_utils.RegisterDay(registry, 21, 2, day02.Part1, day02.Part2)
	aoc_utils.RegisterDay(registry, 21, 3, day03.Part1, day03.Part2)
	aoc_utils.RegisterDay(registry, 21, 4, day04.Part1, day04.Part2)
}
