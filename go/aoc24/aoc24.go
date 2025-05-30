package aoc24

import (
	"github.com/donmahallem/aoc/aoc24/day01"
	"github.com/donmahallem/aoc/aoc24/day02"
	"github.com/donmahallem/aoc/aoc24/day03"
	"github.com/donmahallem/aoc/aoc24/day04"
	"github.com/donmahallem/aoc/aoc24/day05"
	"github.com/donmahallem/aoc/aoc24/day06"
	"github.com/donmahallem/aoc/aoc24/day07"
	"github.com/donmahallem/aoc/aoc24/day08"
	"github.com/donmahallem/aoc/aoc24/day09"
	"github.com/donmahallem/aoc/aoc24/day10"
	"github.com/donmahallem/aoc/aoc24/day11"
	"github.com/donmahallem/aoc/aoc24/day12"
	"github.com/donmahallem/aoc/aoc24/day13"
	"github.com/donmahallem/aoc/aoc24/day14"
	"github.com/donmahallem/aoc/aoc24/day15"
	"github.com/donmahallem/aoc/aoc24/day16"
	"github.com/donmahallem/aoc/aoc24/day17"
	"github.com/donmahallem/aoc/aoc24/day18"
	"github.com/donmahallem/aoc/aoc24/day19"
	"github.com/donmahallem/aoc/aoc24/day20"
	"github.com/donmahallem/aoc/aoc24/day21"
	"github.com/donmahallem/aoc/aoc24/day22"
	"github.com/donmahallem/aoc/aoc24/day23"
	"github.com/donmahallem/aoc/aoc_utils"
)

func RegisterParts(registry *aoc_utils.Registry) {
	regFunc := registry.CreateYearRegistry(24)
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
	regFunc(14, day14.Part1, day14.Part2)
	regFunc(15, day15.Part1, day15.Part2)
	regFunc(16, day16.Part1, day16.Part2)
	regFunc(17, day17.Part1, day17.Part2)
	regFunc(18, day18.Part1, day18.Part2)
	regFunc(19, day19.Part1, day19.Part2)
	regFunc(20, day20.Part1, day20.Part2)
	regFunc(21, day21.Part1, day21.Part2)
	regFunc(22, day22.Part1, day22.Part2)
	regFunc(23, day23.Part1, day23.Part2)
}
