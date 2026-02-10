package day06

import (
	"io"
	"math"
)

func countOptions(race race) int {
	// guard against invalid input
	if race.Time <= 0 {
		return 0
	}

	// initialDelay := (race.Distance / race.Time)
	// if (race.Distance % race.Time) != 0 {
	// 	initialDelay += 1
	// }
	fTime := float64(race.Time)
	fDistance := float64(race.Distance)
	disc := (fTime * fTime) - 4*fDistance
	if disc < 0 {
		return 0
	}
	tmp := math.Sqrt(disc)
	fLowerBound := (fTime - tmp) / 2
	fUpperBound := (fTime + tmp) / 2
	upperBound := int(fUpperBound)
	if fUpperBound == float64(upperBound) {
		upperBound--
	}
	return upperBound - int(fLowerBound)
}

func Part1(in io.Reader) (int, error) {
	races := parseInput(in)
	score := 1
	for _, race := range races {
		score *= countOptions(race)
	}
	return score, nil
}
