package aoc22

import (
	"github.com/donmahallem/aoc/go/aoc22/day01"
	"github.com/donmahallem/aoc/go/aoc22/day02"
	"github.com/donmahallem/aoc/go/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	aoc_utils.RegisterDay(registry, 22, 1, day01.Part1, day01.Part2)
	aoc_utils.RegisterDay(registry, 22, 2, day02.Part1, day02.Part2)
}
