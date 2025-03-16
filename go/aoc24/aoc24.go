package aoc24

import (
	"fmt"
	"time"

	"github.com/donmahallem/aoc/aoc24/day01"
	"github.com/donmahallem/aoc/aoc24/day02"
	"github.com/donmahallem/aoc/aoc24/day03"
)

type operation func()

var a = [][]operation{{day01.Part1, day01.Part2}, //day 1
	{day02.Part1, day02.Part2}, // day 2
	{day03.Part1, day03.Part2}}

func Aoc24(day int, part int) {
	var startTime = time.Now()
	a[day-1][part-1]()
	var endTime = time.Now()
	fmt.Printf("Took: %d\n", endTime.Sub(startTime).Microseconds())
}
