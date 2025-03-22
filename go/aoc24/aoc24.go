package aoc24

import (
	"fmt"
	"os"
	"time"

	"github.com/donmahallem/aoc/aoc24/day01"
	"github.com/donmahallem/aoc/aoc24/day02"
	"github.com/donmahallem/aoc/aoc24/day03"
	"github.com/donmahallem/aoc/aoc24/day04"
	"github.com/donmahallem/aoc/aoc24/day05"
	"github.com/donmahallem/aoc/aoc24/day06"
	"github.com/donmahallem/aoc/aoc24/day07"
	"github.com/donmahallem/aoc/aoc24/day08"
	"github.com/donmahallem/aoc/aoc24/day09"
)

type operation func(in *os.File)

var a = [][]operation{{day01.Part1, day01.Part2}, //day 1
	{day02.Part1, day02.Part2}, // day 2
	{day03.Part1, day03.Part2},
	{day04.Part1, day04.Part2},
	{day05.Part1, day05.Part2},
	{day06.Part1, day06.Part2},
	{day07.Part1, day07.Part2},
	{day08.Part1, day08.Part2},
	{day09.Part1, day09.Part2}}

func Aoc24(day int, part int) {
	var startTime = time.Now()
	a[day-1][part-1](os.Stdin)
	var endTime = time.Now()
	fmt.Printf("Took: %d\n", endTime.Sub(startTime).Microseconds())
}
