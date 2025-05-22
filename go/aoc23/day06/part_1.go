package day06

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func parseInput(in io.Reader) []Race {
	row1 := make([]int, 0, 4)
	row2 := make([]int, 0, 4)

	s := bufio.NewScanner(in)
	s.Scan()
	for idx, item := range strings.Fields(s.Text()) {
		if idx < 1 {
			continue
		}
		val, _ := strconv.Atoi(item)
		row1 = append(row1, val)
	}
	s.Scan()
	for idx, item := range strings.Fields(s.Text()) {
		if idx < 1 {
			continue
		}
		val, _ := strconv.Atoi(item)
		row2 = append(row2, val)
	}
	races := make([]Race, len(row1))
	for idx := range row1 {
		races[idx] = Race{Time: row1[idx], Distance: row2[idx]}
	}
	return races
}

func CountOptions(race Race) int {

	initialDelay := (race.Distance / race.Time)
	if (race.Distance % race.Time) != 0 {
		initialDelay += 1
	}
	fTime := float64(race.Time)
	fDistance := float64(race.Distance)
	tmp := math.Sqrt((fTime * fTime) - 4*fDistance)
	fLowerBound := (fTime - tmp) / 2
	fUpperBound := (fTime + tmp) / 2
	upperBound := int(fUpperBound)
	if fUpperBound == float64(upperBound) {
		upperBound--
	}
	return upperBound - int(fLowerBound)
}

func Part1(in io.Reader) int {
	races := parseInput(in)
	score := 1
	for _, race := range races {
		score *= CountOptions(race)
	}
	return score
}
