package aoc23

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/donmahallem/aoc/aoc23/day01"
	"github.com/donmahallem/aoc/aoc23/day02"
	"github.com/donmahallem/aoc/aoc23/day03"
	"github.com/donmahallem/aoc/aoc23/day04"
	"github.com/donmahallem/aoc/aoc23/day05"
	"github.com/donmahallem/aoc/aoc_utils"
)

var implementedParts = [][]aoc_utils.AocPart{{day01.Part1, day01.Part2},
	{day02.Part1, day02.Part2},
	{day03.Part1, day03.Part2},
	{day04.Part1, day04.Part2},
	{day05.Part1, day05.Part2}}

func Aoc23(day int, part int) error {
	if day < 1 && day > len(implementedParts) {
		return errors.New("day is not in supported range")
	} else if part < 1 && part > len(implementedParts[day-1]) {
		return errors.New("requested part is not implemented")
	}
	var startTime = time.Now()
	result := implementedParts[day-1][part-1](os.Stdin)
	var endTime = time.Now()
	fmt.Printf("Result: %d\nTook: %d\n", result, endTime.Sub(startTime).Microseconds())
	return nil
}
