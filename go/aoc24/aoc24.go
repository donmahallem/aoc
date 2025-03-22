package aoc24

import (
	"errors"
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

var implementedParts = [][]operation{{day01.Part1, day01.Part2}, //day 1
	{day02.Part1, day02.Part2}, // day 2
	{day03.Part1, day03.Part2},
	{day04.Part1, day04.Part2},
	{day05.Part1, day05.Part2},
	{day06.Part1, day06.Part2},
	{day07.Part1, day07.Part2},
	{day08.Part1, day08.Part2},
	{day09.Part1, day09.Part2}}

func Aoc24(day int, part int) error {
	if day < 1 && day > len(implementedParts) {
		return errors.New("day is not in supported range")
	} else if part < 1 && part > len(implementedParts[day-1]) {
		return errors.New("requested part is not implemented")
	}
	var startTime = time.Now()
	implementedParts[day-1][part-1](os.Stdin)
	var endTime = time.Now()
	fmt.Printf("Took: %d\n", endTime.Sub(startTime).Microseconds())
	return nil
}
