package day11

import (
	"io"
)

/**
 * calculates the combined manhatten distances between all galaxies
 */
func combinedManhattenDistances(galaxies []galaxy) int {

	totalDistance := 0
	totalGalaxies := len(galaxies)
	for idx1 := 0; idx1 < totalGalaxies-1; idx1++ {
		for idx2 := idx1 + 1; idx2 < totalGalaxies; idx2++ {
			g1 := galaxies[idx1]
			g2 := galaxies[idx2]
			totalDistance += int(g1.Position.DistanceManhatten(g2.Position))
		}
	}
	return totalDistance
}

func Part1(in io.Reader) (int, error) {
	start, err := parseInput(in, 1)
	if err != nil {
		return 0, err
	}
	return combinedManhattenDistances(start), nil
}
