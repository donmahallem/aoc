package aoc25

import (
	"github.com/donmahallem/aoc/go/aoc25/day01"
	"github.com/donmahallem/aoc/go/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	regFunc := registry.CreateYearRegistry(25)
	regFunc(1, day01.Part1, day01.Part2)
}
