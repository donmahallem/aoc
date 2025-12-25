package aoc25

import (
	"github.com/donmahallem/aoc/go/aoc23/day03"
	"github.com/donmahallem/aoc/go/aoc25/day01"
	"github.com/donmahallem/aoc/go/aoc25/day02"
	"github.com/donmahallem/aoc/go/aoc25/day04"
	"github.com/donmahallem/aoc/go/aoc25/day05"
	"github.com/donmahallem/aoc/go/aoc25/day06"
	"github.com/donmahallem/aoc/go/aoc25/day07"
	"github.com/donmahallem/aoc/go/aoc25/day11"
	"github.com/donmahallem/aoc/go/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	aoc_utils.RegisterDay(registry, 25, 1, day01.Part1, day01.Part2)
	aoc_utils.RegisterDay(registry, 25, 2, day02.Part1, day02.Part2)
	aoc_utils.RegisterDay(registry, 25, 3, day03.Part1, day03.Part2)
	aoc_utils.RegisterDay(registry, 25, 4, day04.Part1, day04.Part2)
	aoc_utils.RegisterDay(registry, 25, 5, day05.Part1, day05.Part2)
	aoc_utils.RegisterDay(registry, 25, 6, day06.Part1, day06.Part2)
	aoc_utils.RegisterDay(registry, 25, 7, day07.Part1, day07.Part2)
	aoc_utils.RegisterDay(registry, 25, 11, day11.Part1, day11.Part2)
}
