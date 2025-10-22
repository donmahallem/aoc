package aoc23

import (
	"github.com/donmahallem/aoc/aoc23/day01"
	"github.com/donmahallem/aoc/aoc23/day02"
	"github.com/donmahallem/aoc/aoc23/day03"
	"github.com/donmahallem/aoc/aoc23/day04"
	"github.com/donmahallem/aoc/aoc23/day05"
	"github.com/donmahallem/aoc/aoc23/day06"
	"github.com/donmahallem/aoc/aoc23/day07"
	"github.com/donmahallem/aoc/aoc23/day08"
	"github.com/donmahallem/aoc/aoc23/day09"
	"github.com/donmahallem/aoc/aoc23/day10"
	"github.com/donmahallem/aoc/aoc23/day11"
	"github.com/donmahallem/aoc/aoc23/day12"
	"github.com/donmahallem/aoc/aoc23/day13"
	"github.com/donmahallem/aoc/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	regFunc := registry.CreateYearRegistry(23)
	regFunc(1, day01.Part1, day01.Part2)
	regFunc(2, day02.Part1, day02.Part2)
	regFunc(3, day03.Part1, day03.Part2)
	regFunc(4, day04.Part1, day04.Part2)
	regFunc(5, day05.Part1, day05.Part2)
	regFunc(6, day06.Part1, day06.Part2)
	regFunc(7, day07.Part1, day07.Part2)
	regFunc(8, day08.Part1, day08.Part2)
	regFunc(9, day09.Part1, day09.Part2)
	regFunc(10, day10.Part1, day10.Part2)
	regFunc(11, day11.Part1, day11.Part2)
	regFunc(12, day12.Part1, day12.Part2)
	regFunc(13, day13.Part1, day13.Part2)
}
