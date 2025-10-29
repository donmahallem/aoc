package day21

import (
	"io"
)

// Extrapolates the number of visited cells for large number of steps
func SolveInfinite(inp *parsedInput, totalSteps int) int {

	// currently for the small test it doesnt work. Need to investigate further
	tileSize := inp.Width
	stepsModulo := totalSteps % tileSize

	samples := []int{
		stepsModulo,
		stepsModulo + tileSize,
		stepsModulo + 2*tileSize,
	}

	counts := make([]int, len(samples))
	for i, steps := range samples {
		counts[i] = CountVisitedInfinite(inp, steps)
	}

	n := (totalSteps - stepsModulo) / tileSize

	a := (counts[2] - 2*counts[1] + counts[0]) / 2
	b := counts[1] - a - counts[0]
	c := counts[0]

	return a*n*n + b*n + c
}

func Part2(in io.Reader) int {
	parsed := ParseInput(in)
	return SolveInfinite(&parsed, 26501365)
}
